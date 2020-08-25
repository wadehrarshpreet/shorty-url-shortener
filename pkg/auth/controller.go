package auth

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/wadehrarshpreet/short/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// @Summary Login API
// @Description Login to Application
// @ID login
// @Accept  json
// @Produce  json
// @tags auth
// @Param user body userRequest true "Login API"
// @Success 200 {object} auth.userResponse
// @Failure 400 {object} util.ErrorResponse "10001 - Invalid Request Params, 10009 - Invalid Username or Password"
// @Failure 500 {object} util.ErrorResponse "10002 - Something Went Wrong! Please try again later."
// @Router /login [post]
func login(ctx echo.Context) error {
	u := new(userRequest)
	if err := ctx.Bind(u); err != nil {
		return util.GenerateErrorResponse(ctx, http.StatusBadRequest, "INVALID_REQUEST_PARAM")
	}

	// validate params
	if len(u.Username) == 0 || len(u.Password) == 0 {
		return util.GenerateErrorResponse(ctx, http.StatusBadRequest, "INVALID_REQUEST_PARAM")
	}

	// make db query
	userCollection := util.DB.Collection("users")

	query := bson.M{"_id": u.Username}

	returnUser := new(User)
	err := userCollection.FindOne(context.Background(), query).Decode(&returnUser)

	if err != nil {
		// ctx.Logger().Errorf("Error in fetching user data...%s", err)
		return util.GenerateErrorResponse(ctx, http.StatusBadRequest, "AUTH_FAILED")
	}

	// if user found compare password
	matched, err := comparePasswords(returnUser.EncryptPwd, []byte(u.Password))
	if err != nil {
		ctx.Logger().Errorf("Error in comparing password...%s", err)
		return util.GenerateErrorResponse(ctx, http.StatusInternalServerError, "SOMETHING_WRONG")
	}

	if !matched {
		return util.GenerateErrorResponse(ctx, http.StatusBadRequest, "AUTH_FAILED")
	}
	// if password match send JWT token & create session

	// generate new Session
	returnUser.LastLogin = lastLogin{time.Now(), uuid.NewV4().String(), ctx.RealIP()}
	// update session details on db async
	go func() {
		updateResult, err := userCollection.UpdateOne(context.Background(), query, bson.M{"$set": bson.M{"lastLogin": returnUser.LastLogin}})
		if err != nil {
			ctx.Logger().Errorf("Error in updating last Login Info...%s", err)
		}
		ctx.Logger().Infof("Successfully update login Session %v", updateResult)
	}()

	token, err := returnUser.getJWTToken()
	if err != nil {
		ctx.Logger().Errorf("Error in generating JWT token...%s", err)
		return util.GenerateErrorResponse(ctx, http.StatusInternalServerError, "SOMETHING_WRONG")
	}

	return ctx.JSON(http.StatusOK, returnUser.getUserResponse(token))
}

// @Summary Register API
// @Description Register User to Application
// @ID register
// @Accept  json
// @Produce  json
// @tags auth
// @Param registeration body User true "Register API"
// @Success 200 {object} auth.userResponse
// @Failure 400 {object} util.ErrorResponse "10001 - Invalid Request Params, 10007 - Email Already exist, 10006 - Username already taken, 10005 - Password must be 8 character long with at least one number & one character, 10004 - Invalid Email Address, 10003 - Username must be between 4 and 32 characters"
// @Failure 500 {object} util.ErrorResponse "10002 - Something Went Wrong! Please try again later"
// @Router /signup [post]
func signup(ctx echo.Context) error {
	u := new(User)
	if err := ctx.Bind(u); err != nil {
		return util.GenerateErrorResponse(ctx, http.StatusBadRequest, "INVALID_REQUEST_PARAM")
	}

	// validate Input

	// validate Username
	if len(u.Username) < 4 || len(u.Username) > 32 {
		return util.GenerateErrorResponse(ctx, http.StatusBadRequest, "AUTH_INVALID_USERNAME")
	}

	// validate email
	if !util.ValidateEmail(u.Email) {
		return util.GenerateErrorResponse(ctx, http.StatusBadRequest, "AUTH_INVALID_EMAIL")
	}

	// validate password
	if !util.ValidatePassword(u.Password) {
		return util.GenerateErrorResponse(ctx, http.StatusBadRequest, "AUTH_INVALID_PASSWORD")
	}

	// Encrypt Password
	encryptPassword, err := generatePasswordHash([]byte(u.Password))

	if err != nil {
		ctx.Logger().Errorf("Error in generating password hash...%s", err)
		return util.GenerateErrorResponse(ctx, http.StatusInternalServerError, "SOMETHING_WRONG")
	}

	u.Plan = "FREE"
	u.EncryptPwd = encryptPassword
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	// generate Login SESSION
	u.LastLogin = lastLogin{time.Now(), uuid.NewV4().String(), ctx.RealIP()}

	// Write to DB (unique indexes on username & email)
	userCollection := util.DB.Collection("users")

	_, writeErr := userCollection.InsertOne(context.Background(), u)

	if writeErr != nil {
		mongoErr := writeErr.(mongo.WriteException)
		ctx.Logger().Errorf("Error in Inserting User ...%s", err)
		if len(mongoErr.WriteErrors) > 0 {
			for _, mErr := range mongoErr.WriteErrors {
				// username already exist
				if mErr.Code == 11000 {
					if strings.Contains(mErr.Message, "email") {
						return util.GenerateErrorResponse(ctx, http.StatusBadRequest, "AUTH_EMAIL_ALREADY_EXIST")
					}
					return util.GenerateErrorResponse(ctx, http.StatusBadRequest, "AUTH_USER_ALREADY_EXIST")
				}
			}
		}
		return util.GenerateErrorResponse(ctx, http.StatusInternalServerError, "SOMETHING_WRONG")
	}

	// generate JWT Token

	token, err := u.getJWTToken()
	if err != nil {
		ctx.Logger().Errorf("Error in generating JWT token...%s", err)
		return util.GenerateErrorResponse(ctx, http.StatusInternalServerError, "SOMETHING_WRONG")
	}

	return ctx.JSON(http.StatusOK, u.getUserResponse(token))
}
