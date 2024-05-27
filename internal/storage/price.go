package storage

import (
	"context"
	"strings"

	"github.com/TropicalDog17/price-updater/internal/model"
	"github.com/TropicalDog17/price-updater/internal/provider"
)

func (s *Storage) UpdatePrice(ctx context.Context, price *model.PriceInfo) error {
	return s.DB.HSet(ctx, "prices", price.Symbol, price.Price).Err()
}

// Update24hChange updates the 24h change of the symbol
func (s *Storage) Update24hChange(ctx context.Context, price *model.PriceInfo) error {
	return s.DB.HSet(ctx, "change24h", price.Symbol, price.Change24h).Err()
}

// Update1hChange updates the 1h change of the symbol
func (s *Storage) Update1hChange(ctx context.Context, price *model.PriceInfo) error {
	return s.DB.HSet(ctx, "change1h", price.Symbol, price.Change1h).Err()
}

// Update7dChange updates the 7d change of the symbol
func (s *Storage) Update7dChange(ctx context.Context, price *model.PriceInfo) error {
	return s.DB.HSet(ctx, "change7d", price.Symbol, price.Change7d).Err()
}

// UpdateHigh24h updates the 24h high of the symbol
func (s *Storage) UpdateHigh24h(ctx context.Context, price *model.PriceInfo) error {
	return s.DB.HSet(ctx, "high24h", price.Symbol, price.High24h).Err()
}

// UpdateLow24h updates the 24h low of the symbol
func (s *Storage) UpdateLow24h(ctx context.Context, price *model.PriceInfo) error {
	return s.DB.HSet(ctx, "low24h", price.Symbol, price.Low24h).Err()
}

// GetPrice returns the price of the symbol
func (s *Storage) GetPrice(ctx context.Context, symbol string) (float64, error) {
	return s.DB.HGet(ctx, "prices", symbol).Float64()
}

// Get24hChange returns the 24h change of the symbol, in range 0-1
func (s *Storage) Get24hChange(ctx context.Context, symbol string) (float64, error) {
	return s.DB.HGet(ctx, "24hchanges", symbol).Float64()
}

func (s *Storage) Get1hChange(ctx context.Context, symbol string) (float64, error) {
	return s.DB.HGet(ctx, "1hchanges", symbol).Float64()
}

func (s *Storage) Get7dChange(ctx context.Context, symbol string) (float64, error) {
	return s.DB.HGet(ctx, "7dchanges", symbol).Float64()
}

func (s *Storage) GetHigh24h(ctx context.Context, symbol string) (float64, error) {
	return s.DB.HGet(ctx, "high24h", symbol).Float64()
}

func (s *Storage) GetLow24h(ctx context.Context, symbol string) (float64, error) {
	return s.DB.HGet(ctx, "low24h", symbol).Float64()
}

func (s *Storage) FetchPriceFromProvider(ctx context.Context, provider provider.DataProvider, symbols []string) (map[string]model.PriceInfo, error) {
	result := make(map[string]model.PriceInfo)
	symbolsString := strings.Join(symbols, ",")
	resp, err := provider.FetchPrice(symbolsString)
	if err != nil {
		return nil, err
	}
	for _, data := range resp {
		result[data.Symbol] = model.PriceInfo{
			Symbol:    data.Symbol,
			Price:     float32(data.CurrentPrice),
			Change24h: float32(data.Change24h),
			Change1h:  float32(data.Change1h),
			Change7d:  float32(data.Change7d),
			High24h:   float32(data.High24h),
			Low24h:    float32(data.Low24h),
		}
	}
	return result, nil
}
