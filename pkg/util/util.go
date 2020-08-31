package util

import (
	"crypto/rand"
	"log"
	"os"
	"regexp"

	"github.com/labstack/echo/v4"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var passwordRegexp = regexp.MustCompile("(?i)^(?:[0-9]+[a-z]|[a-z]+[0-9])[a-z0-9]*$")
var alphaNumericURegex = regexp.MustCompile("[a-zA-Z0-9_]{4,}")

// Getenv reads environment variables for given key and second argument as fallback value
func Getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// ErrorResponse error response format
type ErrorResponse struct {
	Error     bool   `json:"error"`
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
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

	return ctx.JSON(code, ErrorResponse{
		Error:     true,
		Message:   errMessage,
		ErrorCode: appErrCode,
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

// ValidateCustomShortURL to validate custom short url
func ValidateCustomShortURL(url string) bool {
	return alphaNumericURegex.MatchString(url)
}

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890" // 52 possibilities
	letterIdxBits = 6                                                                // 6 bits to represent 64 possibilities / indexes
	letterIdxMask = 1<<letterIdxBits - 1                                             // All 1-bits, as many as letterIdxBits
)

// GenerateRandomAlphaNumericString generate alpha numeric string random
func GenerateRandomAlphaNumericString(length int) string {

	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			randomBytes = SecureRandomBytes(bufferSize)
		}
		if idx := int(randomBytes[j%length] & letterIdxMask); idx < len(letterBytes) {
			result[i] = letterBytes[idx]
			i++
		}
	}

	return string(result)
}

// SecureRandomBytes returns the requested number of bytes using crypto/rand
func SecureRandomBytes(length int) []byte {
	var randomBytes = make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("Unable to generate random bytes")
	}
	return randomBytes
}
