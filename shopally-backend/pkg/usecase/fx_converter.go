package usecase

import (
	"context"

	"github.com/shopally-ai/pkg/domain"
)

// PriceConverter converts prices using the FX client.
type PriceConverter struct {
	FX domain.IFXClient
}

// USDToETB converts a USD amount to ETB and returns both the converted amount and the rate used.
func (c PriceConverter) USDToETB(ctx context.Context, usd float64) (etb float64, rate float64, err error) {
	rate, err = c.FX.GetRate(ctx, "USD", "ETB")
	if err != nil {
		return 0, 0, err
	}
	return usd * rate, rate, nil
}
