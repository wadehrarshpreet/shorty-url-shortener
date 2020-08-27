package shorten

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/wadehrarshpreet/short/pkg/util"
	"github.com/wadehrarshpreet/short/pkg/web"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @Summary URL Shortner
// @Description Shorten your URL
// @ID url-shortener
// @Accept  json
// @Produce  json
// @tags url-shortener
// @Param URL-Shortening body urlShorteningRequest true "URL Shortening"
// @Success 200 {object} urlShorteningResponse
// @Security ApiKeyAuth
// @Failure 400 {object} util.ErrorResponse "10001 - Invalid Request Params, 10007 - Email Already exist, 10006 - Username already taken, 10005 - Password must be 8 character long with at least one number & one character, 10004 - Invalid Email Address, 10003 - Username must be between 4 and 32 characters"
// @Failure 500 {object} util.ErrorResponse "10002 - Something Went Wrong! Please try again later"
// @Router /api/v1/short [post]
func urlShortener(c echo.Context) error {
	// Parse Request
	r := new(urlShorteningRequest)
	if err := c.Bind(r); err != nil {
		return util.GenerateErrorResponse(c, http.StatusBadRequest, "INVALID_REQUEST_PARAM")
	}

	user, userFound := c.Get("user").(*jwt.Token)
	var username = c.RealIP()

	if !userFound {
		// if no userFound then Custom parameters will be ignored
		r.Custom = ""
	} else if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		username = claims["name"].(string)
	}

	// validate URL
	if !isURL(r.URL) {
		return util.GenerateErrorResponse(c, http.StatusBadRequest, "INVALID_REQUEST_PARAM")
	}

	isCustomURL := len(r.Custom) != 0

	// user choose reserve keywords or invalidate custom URL
	if isCustomURL && (web.ReservedWebsiteRoutesMap[strings.ToLower(r.Custom)] || !util.ValidateCustomShortURL(r.Custom)) {
		return util.GenerateErrorResponse(c, http.StatusBadRequest, "URL_ALREADY_TAKEN")
	}

	// Limit max 100(get from env) short url based on IP address
	// Login User can do 200 max(FREE plan)

	// prevent recreate
	// md5 hash of long URL
	md5Hash := fmt.Sprintf("%x", md5.Sum([]byte(r.URL)))

	mongoObj := urlShortenModel{
		ID:         "",
		URL:        r.URL,
		URLHash:    md5Hash,
		CreatedAt:  time.Now(),
		CreatedBy:  username,
		VisitCount: 0,
		CustomURL:  isCustomURL,
	}

	responseData := make(chan interface{})
	// give max 5 second for generate short code & write
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go createShortURL(&mongoObj, r.Custom, responseData)
	// shortCode, err :=
	select {
	case shortURL := <-responseData:
		{
			if err, isError := shortURL.(error); isError {
				c.Logger().Errorf("Error in generating short URL...%s", err)
				if fmt.Sprintf("%s", err) == "URL_ALREADY_TAKEN" {
					return util.GenerateErrorResponse(c, http.StatusBadRequest, "URL_ALREADY_TAKEN")
				}
				return util.GenerateErrorResponse(c, http.StatusInternalServerError, "SOMETHING_WRONG")
			}

			return c.JSON(http.StatusOK, urlShorteningResponse{
				URL:      r.URL,
				ShortURL: shortURL.(string),
			})
		}
	case <-ctx.Done():
		{
			// timeout
			c.Logger().Errorf("20s Timeout exceed...")
			return util.GenerateErrorResponse(c, http.StatusInternalServerError, "SOMETHING_WRONG")
		}
	}
}

func createShortURL(mongoObj *urlShortenModel, customURL string, responseData chan<- interface{}) {
	// Generate Random String
	var shortCode string
	isCustom := false
	if len(customURL) == 0 {
		shortCode = util.GenerateRandomAlphaNumericString(characterLength)
	} else {
		shortCode = customURL
		isCustom = true
	}

	// Write to DB
	mongoObj.ID = shortCode
	urlCollection := util.DB.Collection("url")

	updateOpt := options.FindOneAndUpdate()
	updateOpt.SetUpsert(true)
	updateOpt.SetReturnDocument(options.After)

	var writeErr error
	returnData := new(urlShortenModel)

	// don't return custom short URL for non logged In user
	// write only if new URL else return old

	if isCustom {
		// always run insert command
		writeErr = urlCollection.FindOneAndUpdate(context.Background(), bson.M{"$and": bson.A{
			bson.M{"_id": mongoObj.ID},
			bson.M{"urlHash": mongoObj.URLHash},
		}}, bson.M{"$setOnInsert": mongoObj}, updateOpt).Decode(&returnData)

	} else {
		writeErr = urlCollection.FindOneAndUpdate(context.Background(), bson.M{"$and": bson.A{
			bson.M{"urlHash": mongoObj.URLHash},
			bson.M{"URL": mongoObj.URL},
			bson.M{"customURL": bson.M{"$ne": !isCustom}},
		}}, bson.M{"$setOnInsert": mongoObj}, updateOpt).Decode(&returnData)
	}

	if writeErr != nil {
		mongoErr := writeErr.(mongo.CommandError)
		if mongoErr.Code == 11000 {
			if isCustom {
				responseData <- errors.New("URL_ALREADY_TAKEN")
			} else {
				createShortURL(mongoObj, customURL, responseData)
			}
			return
		}
		responseData <- writeErr
		return
	}

	responseData <- returnData.ID
}
