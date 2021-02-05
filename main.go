package main

import (
	"errors"
	"golang-graphql-subscriptions/infrastructure/graphql"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"golang-graphql-subscriptions/infrastructure/datastore"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	})

	client, err := datastore.NewRedisClient("localhost:6379")
	if !errors.Is(err, nil) {
		log.Fatalln(err)
	}
	defer client.Close()

	{
		// For CORS
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))

		srv := graphql.NewGraphQLServer(client)

		// For Query and Mutations
		e.POST("/query", func(c echo.Context) error {
			srv.ServeHTTP(c.Response(), c.Request())
			return nil
		})

		// For Subscriptions
		e.GET("/subscriptions", func(c echo.Context) error {
			srv.ServeHTTP(c.Response(), c.Request())
			return nil
		})

		e.GET("/playground", func(c echo.Context) error {
			playground.Handler("GraphQL", "/query").ServeHTTP(c.Response(), c.Request())
			return nil
		})
	}

	e.Logger.Fatal(e.Start(":8080"))
}
