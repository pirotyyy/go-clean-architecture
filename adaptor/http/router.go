package http

import (
	"ca-tech/domain/service"
	"ca-tech/infra"
	"ca-tech/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	healthCheckGroup := e.Group("/health_check")
	{
		relativePath := ""
		healthCheckGroup.GET(relativePath, healthChcek)
	}

	sqlConn := infra.SqlConnector()
	userRepostiroy := infra.NewUserRepository(sqlConn.Conn)
	userService := service.NewUserService(userRepostiroy)
	userUsecase := usecase.NewUserUsecase(userService)

	userGroup := e.Group("/user")
	{
		handler := NewUserHandler(userUsecase)
		userGroup.POST("/create", handler.CreateUser())
		userGroup.GET("/get", handler.GetUser())
		userGroup.PUT("/update", handler.UpdateUser())
	}

	gachaRepository := infra.NewGachaRepository(sqlConn.Conn)
	gachaService := service.NewGachaService(gachaRepository)
	gachaUsecase := usecase.NewGachaUsecase(gachaService)

	gachaGroup := e.Group("/gacha")
	{
		handler := NewGachaHandler(gachaUsecase)
		gachaGroup.POST("/draw", handler.Draw())
	}

	return e
}
