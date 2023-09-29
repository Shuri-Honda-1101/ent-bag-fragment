package main

import (
	"context"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Shuri-Honda-1101/ent-bug-fragment/ent"
	"github.com/Shuri-Honda-1101/ent-bug-fragment/ent/migrate"
	"github.com/Shuri-Honda-1101/ent-bug-fragment/resolver"

	_ "github.com/lib/pq"
)

func main() {
	// Create ent.Client and run the schema migration.
	client, err := ent.Open(dialect.Postgres, "host=localhost port=5434 user=user dbname=db password=password sslmode=disable")
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	// Configure the server and start listening on :8081.
	schema := resolver.NewSchema(client)
	srv := handler.NewDefaultServer(schema)
	http.Handle("/",
		playground.Handler("GraphQL", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
