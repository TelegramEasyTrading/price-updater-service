package storage

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

type StorageInterface interface {

	// Price related methods
	// GetPrice(ctx context.Context, symbol string) (float64, error)
	// GetBatchPrices(ctx context.Context, symbols []string) (map[string]float64, error)

	GetDB() *redis.Client
	// Close() error
}

// Storage contains an SQL db. Storage implements the StorageInterface.
type Storage struct {
	DB *redis.Client
}

func (s *Storage) GetDB() *redis.Client {
	return s.DB
}

func NewRedisClient() (*Storage, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ENDPOINT") + ":" + os.Getenv("REDIS_PORT"),
		DB:       0,
		Password: os.Getenv("REDIS_PASSWORD"),
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &Storage{
		DB: client,
	}, nil
}

func NewLocalRedisClient() (*Storage, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &Storage{
		DB: client,
	}, nil
}
