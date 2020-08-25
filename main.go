package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/wadehrarshpreet/short/docs"
	"github.com/wadehrarshpreet/short/pkg/auth"
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

// @host localhost:1234
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	// Create channel for shutdown signals.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, syscall.SIGTERM)

	// Initialize env variable
	err := godotenv.Load("configs/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := util.Getenv("PORT", "1234")
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

	apiGroup := e.Group("/api", middleware.JWT([]byte(util.Getenv("JWT_SECRET", "shorty-secret123"))))

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
