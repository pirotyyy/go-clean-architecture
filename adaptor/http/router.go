package http

import (
	"ca-tech/domain/service"
	"ca-tech/infra/db"
	"ca-tech/infra/db/character"
	"ca-tech/infra/db/user"
	userCharacter "ca-tech/infra/db/user_character"
	"ca-tech/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	dbConn := db.DBConnector()
	userRepostiroy := user.NewUserRepository(dbConn.SqlConn)
	characterRepository := character.NewCharacterRepository(dbConn.SqlConn)
	userCharacterRepository := userCharacter.NewUserCharacterRepository(dbConn.SqlConn)

	userService := service.NewUserService(userRepostiroy)
	gachaService := service.NewGachaService(userRepostiroy, characterRepository, userCharacterRepository)
	userCharacterService := service.NewUserCharacterRepository(userRepostiroy, userCharacterRepository)

	userUsecase := usecase.NewUserUsecase(userService)
	gachaUsecase := usecase.NewGachaUsecase(gachaService)
	userCharacterUsecase := usecase.NewUserCharacterUsecase(userCharacterService)

	healthCheckGroup := e.Group("/health_check")
	{
		relativePath := ""
		healthCheckGroup.GET(relativePath, healthChcek)
	}

	userGroup := e.Group("/user")
	{
		handler := NewUserHandler(userUsecase)
		userGroup.POST("/create", handler.CreateUser())
		userGroup.GET("/get", handler.GetUser())
		userGroup.PUT("/update", handler.UpdateUser())
	}

	gachaGroup := e.Group("/gacha")
	{
		handler := NewGachaHandler(gachaUsecase)
		gachaGroup.POST("/draw", handler.Draw())
	}

	userCharacterGroup := e.Group("/user_character")
	{
		handler := NewUserCharacterHandler(userCharacterUsecase)
		userCharacterGroup.GET("/list", handler.GetUserCharactersByToken())
	}

	return e
}
