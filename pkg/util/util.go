package util

import (
	"os"
	"regexp"

	"github.com/labstack/echo/v4"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var passwordRegexp = regexp.MustCompile("(?i)^(?:[0-9]+[a-z]|[a-z]+[0-9])[a-z0-9]*$")

// Getenv reads environment variables for given key and second argument as fallback value
func Getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GenerateErrorResponse used to return json response for error
func GenerateErrorResponse(ctx echo.Context, code int, pathOrMessage string) error {
	var appErrCode = 999
	var errMessage = pathOrMessage
	message, ok := errorMap[pathOrMessage]
	if !ok {
		errMessage = pathOrMessage
	} else {
		appErrCode = message.errCode
		errMessage = message.message
	}

	return ctx.JSON(code, echo.Map{
		"error":     true,
		"message":   errMessage,
		"errorCode": appErrCode,
	})
}

// ValidateEmail to validate an email address
func ValidateEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// ValidatePassword to validate password
func ValidatePassword(password string) bool {
	return len(password) > 8 && passwordRegexp.MatchString(password)
}
