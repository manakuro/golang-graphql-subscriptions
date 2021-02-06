//go:generate go run github.com/99designs/gqlgen

package graph

import (
	"encoding/json"
	"errors"
	"golang-graphql-subscriptions/graph/model"
	"log"
	"sync"

	"github.com/go-redis/redis"
)

type Resolver struct {
	RedisClient     *redis.Client
	messageChannels map[string]chan *model.Message
	mutex           sync.Mutex
}

func NewResolver(client *redis.Client) *Resolver {
	return &Resolver{
		RedisClient:     client,
		messageChannels: map[string]chan *model.Message{},
		mutex:           sync.Mutex{},
	}
}

func (r *Resolver) SubscribeRedis() {
	log.Println("Start Redis Stream...")

	go func() {
		pubsub := r.RedisClient.Subscribe("room")
		defer pubsub.Close()

		for {
			psms, err := pubsub.Receive()
			if !errors.Is(err, nil) {
				panic(err)
			}

			switch data := psms.(type) {
			case *redis.Message:
				msg := &model.Message{}
				if err := json.Unmarshal([]byte(data.Payload), &msg); !errors.Is(err, nil) {
					log.Println(err)
					continue
				}

				r.mutex.Lock()
				for _, ch := range r.messageChannels {
					ch <- msg
				}
				r.mutex.Unlock()
			}
		}
	}()
}
