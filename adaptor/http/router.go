package http

import (
	"ca-tech/adaptor/http/gacha"
	"ca-tech/adaptor/http/health"
	userHTTP "ca-tech/adaptor/http/user"
	userCharacterHTTP "ca-tech/adaptor/http/user_character"
	"ca-tech/domain/model"
	"ca-tech/domain/service"
	"ca-tech/infra/cache"
	"ca-tech/infra/db"
	"ca-tech/infra/db/character"
	"ca-tech/infra/db/user"
	userCharacter "ca-tech/infra/db/user_character"
	"ca-tech/usecase"
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	dbConn := db.DBConnector()
	cacheConn := cache.CacheConnector()

	characters, err := GetCharacters(context.Background(), dbConn.SqlConn)
	if err != nil {
		log.Println(err)
	}

	jsonData, err := json.Marshal(characters)
	if err != nil {
		log.Fatal(err)
	}

	err = cacheConn.RedisConn.Set(context.Background(), "characters", jsonData, 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	userRepostiroy := user.NewUserRepository(dbConn.SqlConn)
	characterRepository := character.NewCharacterRepository(dbConn.SqlConn)
	userCharacterRepository := userCharacter.NewUserCharacterRepository(dbConn.SqlConn)

	userService := service.NewUserService(userRepostiroy)
	gachaService := service.NewGachaService(userRepostiroy, characterRepository, userCharacterRepository)
	userCharacterService := service.NewUserCharacterRepository(userRepostiroy, userCharacterRepository)

	userUsecase := usecase.NewUserUsecase(userService)
	gachaUsecase := usecase.NewGachaUsecase(userService, characterRepository, gachaService, userCharacterService, cacheConn.RedisConn)
	userCharacterUsecase := usecase.NewUserCharacterUsecase(userCharacterService)

	healthCheckGroup := e.Group("/health_check")
	{
		handler := health.NewHealthCheckHandler()
		relativePath := ""
		healthCheckGroup.GET(relativePath, handler.HealthCheck())
	}

	userGroup := e.Group("/user")
	{
		handler := userHTTP.NewUserHandler(userUsecase)
		userGroup.POST("/create", handler.CreateUser())
		userGroup.GET("/get", handler.GetUser())
		userGroup.PUT("/update", handler.UpdateUser())
	}

	gachaGroup := e.Group("/gacha")
	{
		handler := gacha.NewGachaHandler(gachaUsecase)
		gachaGroup.POST("/draw", handler.Draw())
	}

	userCharacterGroup := e.Group("/user_character")
	{
		handler := userCharacterHTTP.NewUserCharacterHandler(userCharacterUsecase)
		userCharacterGroup.GET("/list", handler.GetUserCharactersByToken())
	}

	return e
}

func GetCharacters(ctx context.Context, db *sql.DB) ([]*model.Character, error) {
	const selectCommand = "SELECT id, name, rarity FROM game_character"

	rows, err := db.QueryContext(ctx, selectCommand)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var characters []*model.Character
	for rows.Next() {
		var c model.Character
		if err := rows.Scan(&c.CharacterID, &c.Name, &c.Rarity); err != nil {
			return nil, err
		}
		characters = append(characters, &c)
	}

	return characters, nil
}
