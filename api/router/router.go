package router

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mopeneko/hacku2020-web/api/controller"
)

var router *echo.Echo

func init() {
	ctx := context.Background()

	router = echo.New()

	router.Use(middleware.Recover())

	controller.Bootstrap(ctx, router, nil)
}

func Run() {
	router.Logger.Fatal(router.Start(":1323"))
}
