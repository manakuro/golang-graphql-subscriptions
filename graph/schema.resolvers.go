package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/thanhpk/randstr"

	"golang-graphql-subscriptions/graph/generated"
	"golang-graphql-subscriptions/graph/model"
)

func (r *mutationResolver) CreateMessage(ctx context.Context, message string) (*model.Message, error) {
	m := model.Message{
		Message: message,
	}
	data, err := json.Marshal(m)
	if !errors.Is(err, nil) {
		return nil, err
	}

	r.RedisClient.Publish("room", data)

	return &m, nil
}

func (r *queryResolver) Message(ctx context.Context) ([]*model.Message, error) {
	panic(fmt.Errorf("not implemented"))
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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
