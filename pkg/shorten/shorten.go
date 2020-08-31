package shorten

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wadehrarshpreet/short/pkg/util"
)

// shorten URL routes

const (
	apiVersionV1 = "v1"
)

var (
	characterLength, _ = strconv.Atoi(util.Getenv("SHORT_CHAR_LENGTH", "7"))
	// ShortURLAPIRoute route of Short URL API
	ShortURLAPIRoute = fmt.Sprintf("/%s/short", apiVersionV1)
)

// URLShorteningRequest is request of url shortening api
type URLShorteningRequest struct {
	// URL long url
	URL string `json:"url"`
	// Custom name for short URL only valid for logged in user
	Custom string `json:"custom"`
}

type urlShortenModel struct {
	ID        string    `json:"shortUrl" bson:"_id"`
	URL       string    `json:"url" bson:"url"`
	URLHash   string    `json:"urlHash" bson:"urlHash"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	// CreatedBy store either username or IP address
	CreatedBy  string `json:"createdBy" bson:"createdBy"`
	VisitCount int    `json:"visitCount" bson:"visitCount"`
	CustomURL  bool   `json:"customURL" bson:"customURL"`
}

type urlShorteningResponse struct {
	// URL is long URL (input)
	URL string `json:"url"`
	// ShortURL actual shortURL
	ShortURL string `json:"shortUrl"`
}

// Init is initializing api routes related to url shortener
func Init(e *echo.Group) error {

	// Shortening URL Route

	e.POST(ShortURLAPIRoute, urlShortener)

	// e.Post("/sm")
	return nil
}

// InitRedirectionRoute intializes short to long url redirection routes
func InitRedirectionRoute(e *echo.Echo) error {

	e.GET("/:shortCode", urlRedirection)

	return nil
}

func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func (u *urlShortenModel) generateNewID() {
	u.ID = util.GenerateRandomAlphaNumericString(characterLength)
}
