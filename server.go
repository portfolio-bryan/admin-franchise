package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bperezgo/admin_franchise/config"
	"github.com/bperezgo/admin_franchise/graph"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		log.Fatal("error loading .env file", err)
	}

	c := config.GetConfig()
	port := c.ServerPort

	resolver := graph.NewResolver()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
