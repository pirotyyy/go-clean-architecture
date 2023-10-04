package repository

import (
	"ca-tech/domain/model"
	"context"
)

type GachaRepository interface {
	Draw(ctx context.Context, times int64) ([]*model.GachaResult, error)
}
