package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-graphql-subscriptions/graph/generated"
	"golang-graphql-subscriptions/graph/model"
	"log"

	"github.com/thanhpk/randstr"

	"github.com/go-redis/redis"
)

func (r *mutationResolver) CreateMessage(ctx context.Context, message string) (*model.Message, error) {
	m := model.Message{
		Message: message,
	}

	r.RedisClient.XAdd(&redis.XAddArgs{
		Stream: "room",
		ID:     "*",
		Values: map[string]interface{}{
			"message": m.Message,
		},
	})

	return &m, nil
}

func (r *subscriptionResolver) MessageCreated(ctx context.Context) (<-chan *model.Message, error) {
	token := randstr.Hex(16)
	mc := make(chan *model.Message, 1)
	r.mutex.Lock()
	r.messageChannels[token] = mc
	r.mutex.Unlock()

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.messageChannels, token)
		r.mutex.Unlock()
		log.Println("Deleted")
	}()

	log.Println("Subscription: message created")

	return mc, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
