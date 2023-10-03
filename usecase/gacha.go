package usecase

import (
	"ca-tech/domain/model"
	"ca-tech/domain/service"
	"context"
)

type GachaUsecase interface {
	Draw(ctx context.Context, times int64, token string) ([]*model.Character, error)
}

type gachaUsecase struct {
	svc service.GachaService
}

func NewGachaUsecase(gs service.GachaService) GachaUsecase {
	return &gachaUsecase{
		svc: gs,
	}
}

func (gu *gachaUsecase) Draw(ctx context.Context, times int64, token string) ([]*model.Character, error) {
	// user
	// gacha service
	err := gu.svc.InitCharacterList(ctx)
	if err != nil {
		return nil, err
	}
	return gu.svc.Draw(ctx, times, token)
}
