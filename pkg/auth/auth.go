package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/wadehrarshpreet/short/pkg/util"
	"go.mongodb.org/mongo-driver/mongo"
)

// User struct contains user info
type User struct {
	Username   string    `json:"username" bson:"_id"`
	Email      string    `json:"email" bson:"email"`
	Password   string    `json:"password" bson:"-"`
	Plan       string    `json:"-" bson:"plan"`
	EncryptPwd string    `json:"-" bson:"password"`
	CreatedAt  time.Time `json:"-" bson:"createdAt"`
	UpdatedAt  time.Time `json:"-" bson:"UpdatedAt"`
}

// Init initializes auth routes and services
func Init(e *echo.Echo) error {

	// initate routes
	e.POST("/login", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, echo.Map{
			"success": true,
		})
	})

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

		// Write to DB (unique indexes on username & email)
		userCollection := util.DB.Collection("users")

		res, err := userCollection.InsertOne(context.Background(), u)

		if err != nil {
			mongoErr := err.(mongo.WriteException)
			e.Logger.Errorf("Error in Inserting User ...%s", err)
			if len(mongoErr.WriteErrors) > 0 {
				for _, mErr := range mongoErr.WriteErrors {
					fmt.Printf("prin %v", mErr.Message)
					// username already exist
					if mErr.Code == 11000 {
						return util.GenerateErrorResponse(ctx, http.StatusInternalServerError, "USER_ALREADY_EXIST")
					}
				}
			}
			return util.GenerateErrorResponse(ctx, http.StatusInternalServerError, "SOMETHING_WRONG")
		}

		// generate JWT Token
		fmt.Printf("data %v", u)
		fmt.Printf("data %v", res)

		token, err := u.getJWTToken()
		if err != nil {
			return util.GenerateErrorResponse(ctx, http.StatusInternalServerError, "SOMETHING_WRONG")
		}

		return ctx.JSON(http.StatusOK, echo.Map{
			"token": token,
		})
	})

	return nil
}

func (u *User) getJWTToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = u.Username
	claims["plan"] = u.Plan
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(util.Getenv("JWT_SECRET", "shorty-secret123")))

	if err != nil {
		return "", err
	}

	return t, nil
}
