package storage_test

import (
	"context"
	"testing"

	"github.com/TropicalDog17/price-updater/internal/model"
	"github.com/TropicalDog17/price-updater/internal/storage"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	ctx := context.Background()

	// load env

	s, err := storage.NewLocalRedisClient()
	require.NoError(t, err)
	t.Run("Update Price", func(t *testing.T) {
		s.UpdatePrice(ctx, &model.PriceInfo{
			Symbol: "mockTest",
			Price:  60000,
		})
		price, err := s.GetPrice(ctx, "mockTest")
		require.NoError(t, err)
		require.Equal(t, 60000.0, price)
		s.UpdatePrice(ctx, &model.PriceInfo{
			Symbol: "mockTest",
			Price:  75000,
		})
		price, err = s.GetPrice(ctx, "mockTest")
		require.NoError(t, err)
		require.Equal(t, 75000.0, price)
	})

}
