package gacha

import (
	"ca-tech/domain/model/gacha"
	"context"
)

type GachaRepository interface {
	Draw(ctx context.Context, times int64) ([]*gacha.GachaResult, error)
}
