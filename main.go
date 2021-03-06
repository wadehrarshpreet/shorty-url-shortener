package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/wadehrarshpreet/short/docs"
	"github.com/wadehrarshpreet/short/pkg/auth"
	"github.com/wadehrarshpreet/short/pkg/shorten"
	"github.com/wadehrarshpreet/short/pkg/util"
	"github.com/wadehrarshpreet/short/pkg/web"
)

// @title Shorty URL Shortner
// @version 0.0.1
// @description This is a API Docs of Shorty URL Shortner Service.

// @contact.name Arshpreet Wadehra
// @contact.url https://github.com/wadehrarshpreet/shorty
// @contact.email me@wadehrarshpreet.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	// Create channel for shutdown signals.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, syscall.SIGTERM)

	// pick PORT from env before initialize .env
	port := util.Getenv("PORT", "1234")

	// Initialize env variable
	err := godotenv.Load("configs/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	// e.Logger.SetLevel(log.DEBUG)

	// Init Swagger Routes
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Connect Database
	dbErr := util.ConnectDatabase()
	if dbErr != nil {
		e.Logger.Fatalf("Error in DB Connection...%s", dbErr)
	}

	// Init Request Id middleware
	e.Use(middleware.RequestID())

	// Init Request Logger
	e.Use(middleware.Logger())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// if JWT pass then check for valid Auth if not bypass and let API decide
	apiGroup := e.Group("/api", middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(util.Getenv("JWT_SECRET", "shorty-secret123")),
		Skipper: func(c echo.Context) bool {
			headerToken := c.Request().Header.Get("Authorization")
			if strings.Contains(headerToken, "Bearer") {
				// skip check if hitting shorten url api & custom is not provided
				if strings.Contains(c.Request().URL.String(), shorten.ShortURLAPIRoute) {
					rq := c.Request()
					bodyBytes, _ := ioutil.ReadAll(rq.Body)
					rq.Body.Close() //  must close
					rq.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
					// check if custom param missing
					r := new(shorten.URLShorteningRequest)

					if err := json.Unmarshal(bodyBytes, &r); err != nil {
						return false
					}
					if len(r.Custom) == 0 {
						return true
					}

				}
				return false
			}
			return true
		},
		ErrorHandlerWithContext: func(err error, c echo.Context) error {
			return util.GenerateErrorResponse(c, http.StatusUnauthorized, "UNAUTHORIZED")
		},
	}))

	apiRouteError := shorten.Init(apiGroup)
	if apiRouteError != nil {
		e.Logger.Fatalf("Error in initializing APIs...%s", apiRouteError)
	}

	apiGroup.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"status":     "ok",
			"appVersion": util.Getenv("APP_VERSION", "0.0.1"),
		})
	})

	// Init Static Assets
	e.Static("/assets", "./web/dist")

	// Initialize Auth Service
	authServiceErr := auth.Init(e)
	if authServiceErr != nil {
		e.Logger.Fatalf("Error in initializing Auth Service... %s", authServiceErr)
	}

	// Initialize Webview
	web.InitWebsite(e)

	// Short URL Redirection Route
	shorten.InitRedirectionRoute(e)

	go func() {
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
	}()

	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	e.Logger.Info("Disconnecting Database...")
	if err = util.DbConn.Disconnect(ctx); err != nil {
		panic(err)
	}
	e.Logger.Info("Shutting Down Server...")
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
