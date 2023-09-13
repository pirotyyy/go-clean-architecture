package service

import (
	"ca-tech/domain/model"
	"ca-tech/domain/repository"
	"context"
)

type GachaService interface {
	Draw(ctx context.Context, times int64, token string) (characters []*model.Character, err error)
}

type gachaService struct {
	repo repository.GachaRepository
}

func NewGachaService(gr repository.GachaRepository) GachaService {
	return &gachaService{
		repo: gr,
	}
}

func (gs *gachaService) Draw(ctx context.Context, times int64, token string) (characters []*model.Character, err error) {
	return gs.repo.Draw(ctx, times, token)
}
