package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/TropicalDog17/price-updater/internal/provider"
	"github.com/TropicalDog17/price-updater/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	s, err := storage.NewRedisClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	symbols := []string{"bitcoin", "injective-protocol", "cosmos", "ethereum", "dogecoin", "tether", "dogwifcoin"}
	provider := provider.NewCoinGeckoProvider(os.Getenv("COINGECKO_API_KEY"))
	go func() {
		for {
			prices, err := s.FetchPriceFromProvider(ctx, provider, symbols)
			if err != nil {
				fmt.Println("Error fetching prices", err)
			}

			for _, price := range prices {
				err = s.UpdatePrice(ctx, &price)
				if err != nil {
					fmt.Println("Error updating price", err)
				}
				err = s.Update24hChange(ctx, &price)
				if err != nil {
					fmt.Println("Error updating 24h change", err)
				}
				err = s.Update1hChange(ctx, &price)
				if err != nil {
					fmt.Println("Error updating 1h change", err)
				}
				err = s.Update7dChange(ctx, &price)
				if err != nil {
					fmt.Println("Error updating 7d change", err)
				}
				err = s.UpdateHigh24h(ctx, &price)
				if err != nil {
					fmt.Println("Error updating 24h high", err)
				}
				err = s.UpdateLow24h(ctx, &price)
				if err != nil {
					fmt.Println("Error updating 24h low", err)
				}
			}

			time.Sleep(60 * time.Second)
		}
	}()

	router := gin.Default()

	router.GET("/price", func(c *gin.Context) {
		symbol := c.Query("symbol")
		price, err := s.GetPrice(ctx, symbol)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"price": price})
	})
	router.GET("/change", func(c *gin.Context) {
		symbol := c.Query("symbol")
		price, err := s.Get24hChange(ctx, symbol)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"change": price})
	})
	router.Run(":8081")
}
