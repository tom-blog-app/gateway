package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tom-blog-app/gataway/api/controllers"
	"math/rand"
	"net/http"
	"time"
)

type EchoApp struct {
	*echo.Echo
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//e := echo.New()
	app := &EchoApp{echo.New()}

	app.Use(middleware.Recover())
	app.Use(middleware.CORS())

	//appController := controllers.NewAppControllers()
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	app.GET("/help", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!2")
	})
	app.GET("/log", func(c echo.Context) error {
		LogViaGrpc("test", "test2")
		return c.String(http.StatusOK, "Hello, World!")
	})
	//app.POST("/log/:name/:data", app.LogViaGrpc)

	app.setupControllers()
	app.Logger.Fatal(app.Start(":80"))
}

func (e *EchoApp) setupControllers() {
	//userController := &controllers.PostController{Echo: e.Echo}
	controllers.SetupPostController(e.Echo)
	controllers.SetupHealthController(e.Echo)
}
