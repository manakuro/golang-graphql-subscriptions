package datastore

import (
	"errors"

	"github.com/go-redis/redis"
)

func NewRedisClient(url string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: url,
	})

	_, err := client.Ping().Result()
	if !errors.Is(err, nil) {
		return nil, err
	}

	return client, nil
}
