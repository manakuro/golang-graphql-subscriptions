package main

import (
	"errors"
	"log"

	"github.com/labstack/echo/v4"

	"golang-graphql-subscriptions/graph"
	"golang-graphql-subscriptions/infrastructure/datastore"
	"golang-graphql-subscriptions/infrastructure/graphql"
	"golang-graphql-subscriptions/infrastructure/router"
)

func main() {
	client, err := datastore.NewRedisClient("localhost:6379")
	if !errors.Is(err, nil) {
		log.Fatalln(err)
	}
	defer client.Close()

	r := graph.NewResolver(client)
	r.SubscribeRedis()
	srv := graphql.NewGraphQLServer(r)

	e := router.NewRouter(echo.New(), srv)
	e.Logger.Fatal(e.Start(":8080"))
}
