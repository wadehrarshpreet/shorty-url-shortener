package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/wadehrarshpreet/short/pkg/util"
)

type lastLogin struct {
	Time          time.Time `bson:"time"`
	SessionID     string    `bson:"sessionId"`
	RemoteAddress string    `bson:"remoteIP"`
}

// for swagger usage
type userRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// User struct contains user info
type User struct {
	Username   string    `json:"username" bson:"_id"`
	Password   string    `json:"password" bson:"-"`
	Email      string    `json:"email" bson:"email"`
	Plan       string    `json:"plan,omitempty" bson:"plan"`
	EncryptPwd string    `json:"-" bson:"password"`
	CreatedAt  time.Time `json:"-" bson:"createdAt"`
	UpdatedAt  time.Time `json:"-" bson:"UpdatedAt"`
	LastLogin  lastLogin `json:"-" bson:"lastLogin"`
}

type userResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Plan     string `json:"plan"`
	Token    string `json:"token"`
}

// Init initializes auth routes and services
// initate Auth Routes
func Init(e *echo.Echo) error {

	// Login Route
	e.POST("/login", login)

	// SIGNUP Route
	e.POST("/signup", signup)

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

func (u *User) getUserResponse(token string) userResponse {
	return userResponse{
		Username: u.Username,
		Email:    u.Email,
		Token:    token,
		Plan:     u.Plan,
	}
}
