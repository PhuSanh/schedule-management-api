package router

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"os"
	"os/signal"
	"schedule-management-api/config"
	"schedule-management-api/handler"
	"schedule-management-api/setting"
	"syscall"
	"time"
)

var cfg = config.GetConfig()

func InitRouter() {

	// init services
	_ = setting.InitMysql()


	// init router
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, "Pong")
	})

	apiGroup := e.Group("/api")

	authHandler := new(handler.AuthHandler)
	apiGroup.POST("/login", authHandler.Login)
	apiGroup.POST("/register", authHandler.Register)

	v1Group := apiGroup.Group("/v1")
	v1Group.Use(authHandler.Authenticate)

	userHandler := new(handler.UserHandler)
	v1Group.GET("/users", userHandler.GetList)

	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%v", cfg.Port)); err != nil {
			log.Println("⇛ shutting down the server")
			log.Printf("%v\n", err.Error())
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
