package auth

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/wadehrarshpreet/short/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type lastLogin struct {
	Time          time.Time `bson:"time"`
	SessionID     string    `bson:"sessionId"`
	RemoteAddress string    `bson:"remoteIP"`
}

// User struct contains user info
type User struct {
	Username   string    `json:"username" bson:"_id"`
	Email      string    `json:"email" bson:"email"`
	Password   string    `json:"password" bson:"-"`
	Plan       string    `json:"plan" bson:"plan"`
	EncryptPwd string    `json:"-" bson:"password"`
	CreatedAt  time.Time `json:"-" bson:"createdAt"`
	UpdatedAt  time.Time `json:"-" bson:"UpdatedAt"`
	LastLogin  lastLogin `json:"-" bson:"lastLogin"`
}

// Init initializes auth routes and services
func Init(e *echo.Echo) error {

	// initate Auth Routes

	// Login Route
	e.POST("/login", func(ctx echo.Context) error {
		u := new(User)
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
			// e.Logger.Errorf("Error in fetching user data...%s", err)
			return util.GenerateErrorResponse(ctx, http.StatusBadRequest, "AUTH_FAILED")
		}

		// if user found compare password
		matched, err := comparePasswords(returnUser.EncryptPwd, []byte(u.Password))
		if err != nil {
			e.Logger.Errorf("Error in comparing password...%s", err)
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
				e.Logger.Errorf("Error in updating last Login Info...%s", err)
			}
			e.Logger.Infof("Successfully update login Session %v", updateResult)
		}()

		token, err := returnUser.getJWTToken()
		if err != nil {
			e.Logger.Errorf("Error in generating JWT token...%s", err)
			return util.GenerateErrorResponse(ctx, http.StatusInternalServerError, "SOMETHING_WRONG")
		}

		return ctx.JSON(http.StatusOK, returnUser.getUserResponse(token))
	})

	// SIGNUP Route
	e.POST("/signup", func(ctx echo.Context) error {
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
			e.Logger.Errorf("Error in generating password hash...%s", err)
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
			e.Logger.Errorf("Error in Inserting User ...%s", err)
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
			e.Logger.Errorf("Error in generating JWT token...%s", err)
			return util.GenerateErrorResponse(ctx, http.StatusInternalServerError, "SOMETHING_WRONG")
		}

		return ctx.JSON(http.StatusOK, u.getUserResponse(token))
	})

	return nil
}

// getJWTToken generate JWT token based on User
func (u *User) getJWTToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = u.Username
	claims["plan"] = u.Plan
	claims["sessionId"] = u.LastLogin.SessionID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(util.Getenv("JWT_SECRET", "shorty-secret123")))

	if err != nil {
		return "", err
	}

	return t, nil
}

func (u *User) getUserResponse(token string) echo.Map {
	return echo.Map{
		"username": u.Username,
		"email":    u.Email,
		"token":    token,
		"plan":     u.Plan,
	}
}
