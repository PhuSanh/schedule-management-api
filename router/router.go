package router

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	_ = setting.InitRedis()


	// init router
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, "Pong")
	})

	apiGroup := e.Group("/api")

	v1Group := apiGroup.Group("/v1")

	authHandler := new(handler.AuthHandler)
	v1Group.POST("/login", authHandler.Login)
	v1Group.POST("/register", authHandler.Register)

	//v1Group.Use(authHandler.Authenticate)

	userHandler := new(handler.UserHandler)
	userGroup := v1Group.Group("/user")
	userGroup.GET("/list", userHandler.GetList)
	userGroup.POST("/update/:id", userHandler.Update)
	userGroup.POST("/delete/:id", userHandler.Delete)

	userCategoryHandler := new(handler.UserCategoryHandler)
	userCategoryGroup := v1Group.Group("/user-category")
	userCategoryGroup.GET("/list", userCategoryHandler.GetList)
	userCategoryGroup.POST("/create", userCategoryHandler.Create)
	userCategoryGroup.POST("/update/:id", userCategoryHandler.Update)
	userCategoryGroup.POST("/delete/:id", userCategoryHandler.Delete)

	groupHandler := new(handler.GroupHandler)
	groupGroup := v1Group.Group("/group")
	groupGroup.GET("/list", groupHandler.GetList)
	groupGroup.POST("/create", groupHandler.Create)
	groupGroup.POST("/update/:id", groupHandler.Update)
	groupGroup.POST("/delete/:id", groupHandler.Delete)

	scheduleHandler := new(handler.ScheduleHandler)
	scheduleGroup := v1Group.Group("/schedule")
	scheduleGroup.GET("/list", scheduleHandler.GetList)
	scheduleGroup.POST("/create", scheduleHandler.Create)
	scheduleGroup.POST("/update/:id", scheduleHandler.Update)
	scheduleGroup.POST("/delete/:id", scheduleHandler.Delete)

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
