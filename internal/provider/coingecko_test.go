package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestFetchPrice(t *testing.T) {
	err := godotenv.Load("../../.env")
	require.NoError(t, err)

	symbols := "bitcoin,injective-protocol"
	cgClient := NewCoinGeckoProvider(os.Getenv("COINGECKO_API_KEY"))
	resp, err := cgClient.FetchPrice(symbols)
	fmt.Println("Price of", symbols, "is", resp)

	require.NoError(t, err)
	require.Greater(t, resp[0].CurrentPrice, 0.0)
}
