package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/wadehrarshpreet/short/pkg/auth"
	"github.com/wadehrarshpreet/short/pkg/util"
	"github.com/wadehrarshpreet/short/pkg/web"
)

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

	// Connect Database
	dbErr := util.ConnectDatabase()
	if dbErr != nil {
		e.Logger.Fatalf("Error in DB Connection...%s", dbErr)
	}

	// Init Request Id middleware
	e.Use(middleware.RequestID())

	// Init Request Logger
	e.Use(middleware.Logger())

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
