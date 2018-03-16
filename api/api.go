package main

import (
	"log"
	"net/http"
	"os"

	"github.com/vektah/gqlgen/handler"

	"github.com/tanema/vGoDoRPQl/api/data"
	"github.com/tanema/vGoDoRPQl/api/gql"
	"github.com/tanema/vGoDoRPQl/api/resolvers"
)

type params struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

var frontendURL = "http://localhost:3000"

func init() {
	if os.Getenv("APP_ENV") == "production" {
		frontendURL = "http://localhost:3000" // change this to production domain
	}
}

func main() {
	defer data.DB.Close()

	schema := gql.MakeExecutableSchema(resolvers.New())
	graphQlHandler := corsMiddleware(preflightMiddleware(handler.GraphQL(schema)))

	http.Handle("/", handler.Playground("Todo", "/graphql"))
	http.Handle("/graphql", graphQlHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", frontendURL)
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		next.ServeHTTP(w, r)
	})
}

func preflightMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
