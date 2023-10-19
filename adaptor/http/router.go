package http

import (
	gachaHTTP "ca-tech/adaptor/http/gacha"
	healthHTTP "ca-tech/adaptor/http/health"
	userHTTP "ca-tech/adaptor/http/user"
	userCharacterHTTP "ca-tech/adaptor/http/usercharacter"
	characterModel "ca-tech/domain/model/character"
	cacheCharacterService "ca-tech/domain/service/cachecharacter"
	gachaService "ca-tech/domain/service/gacha"
	userService "ca-tech/domain/service/user"
	userCharacterService "ca-tech/domain/service/usercharacter"
	cacheCharacterCache "ca-tech/infra/cache/character"
	characterDB "ca-tech/infra/db/character"
	userDB "ca-tech/infra/db/user"
	userCharacterDB "ca-tech/infra/db/usercharacter"
	gachaUsecase "ca-tech/usecase/gacha"
	userUsecase "ca-tech/usecase/user"
	userCharacterUsecase "ca-tech/usecase/usercharacter"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
	"log"
)

func InitRouter(dbConn *sql.DB, cacheConn *redis.Client) *echo.Echo {
	e := echo.New()

	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	err := InitCacheCharacter(context.Background(), dbConn, cacheConn)
	if err != nil {
		log.Println(err)
	}

	ur := userDB.NewUserRepository(dbConn)
	cr := characterDB.NewCharacterRepository(dbConn)
	ucr := userCharacterDB.NewUserCharacterRepository(dbConn)
	ccr := cacheCharacterCache.NewCacheCharacterRepository(cacheConn)

	us := userService.NewUserService(ur)
	gs := gachaService.NewGachaService(ur, cr, ucr)
	ucs := userCharacterService.NewUserCharacterRepository(ur, ucr)
	ccs := cacheCharacterService.NewCacheCharacterService(ccr)

	uu := userUsecase.NewUserUsecase(us)
	gu := gachaUsecase.NewGachaUsecase(us, cr, gs, ucs, ccs)
	ucu := userCharacterUsecase.NewUserCharacterUsecase(ucs)

	healthCheckGroup := e.Group("/health_check")
	{
		handler := healthHTTP.NewHealthCheckHandler()
		relativePath := ""
		healthCheckGroup.GET(relativePath, handler.HealthCheck())
	}

	userGroup := e.Group("/user")
	{
		handler := userHTTP.NewUserHandler(uu)
		userGroup.POST("/create", handler.CreateUser())
		userGroup.GET("/get", handler.GetUser())
		userGroup.PUT("/update", handler.UpdateUser())
	}

	gachaGroup := e.Group("/gacha")
	{
		handler := gachaHTTP.NewGachaHandler(gu)
		gachaGroup.POST("/draw", handler.Draw())
	}

	userCharacterGroup := e.Group("/usercharacter")
	{
		handler := userCharacterHTTP.NewUserCharacterHandler(ucu)
		userCharacterGroup.GET("/list", handler.GetUserCharactersByToken())
	}

	return e
}

func GetCharactersByRarity(ctx context.Context, db *sql.DB, rarity characterModel.Rarity) ([]*characterModel.Character, error) {
	const selectCommand = "SELECT id, name, rarity FROM game_character WHERE rarity = ?"

	rows, err := db.QueryContext(ctx, selectCommand, rarity)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var characters []*characterModel.Character
	for rows.Next() {
		var c characterModel.Character
		if err := rows.Scan(&c.CharacterID, &c.Name, &c.Rarity); err != nil {
			return nil, err
		}
		characters = append(characters, &c)
	}

	return characters, nil
}

func InitCacheCharacter(ctx context.Context, db *sql.DB, cacheConn *redis.Client) error {
	charactersN, err := GetCharactersByRarity(ctx, db, characterModel.N)
	if err != nil {
		return err
	}

	charactersR, err := GetCharactersByRarity(ctx, db, characterModel.R)
	if err != nil {
		return err
	}
	charactersSR, err := GetCharactersByRarity(ctx, db, characterModel.SR)
	if err != nil {
		return err
	}
	charactersSSR, err := GetCharactersByRarity(ctx, db, characterModel.SSR)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(charactersN)
	if err != nil {
		return err
	}

	err = cacheConn.Set(ctx, "N", jsonData, 0).Err()
	if err != nil {
		return err
	}

	jsonData, err = json.Marshal(charactersR)
	if err != nil {
		return err
	}

	err = cacheConn.Set(ctx, "R", jsonData, 0).Err()
	if err != nil {
		return err
	}

	jsonData, err = json.Marshal(charactersSR)
	if err != nil {
		return err
	}

	err = cacheConn.Set(ctx, "SR", jsonData, 0).Err()
	if err != nil {
		return err
	}

	jsonData, err = json.Marshal(charactersSSR)
	if err != nil {
		return err
	}

	err = cacheConn.Set(ctx, "SSR", jsonData, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
