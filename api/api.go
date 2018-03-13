package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/vektah/gqlgen/graphql"
	"github.com/vektah/gqlgen/handler"
	"github.com/vektah/gqlgen/neelance/errors"
	"github.com/vektah/gqlgen/neelance/query"
	"github.com/vektah/gqlgen/neelance/validation"

	"github.com/tanema/vGoDoRPQl/api/graph_api"
)

type params struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

func main() {
	defer graph_api.DB.Close()
	http.Handle("/", handler.Playground("Todo", "/graphql"))
	http.Handle("/graphql", GraphQL(graph_api.MakeExecutableSchema(graph_api.New())))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GraphQL(exec graphql.ExecutableSchema) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setCors(w)
		w.Header().Set("Content-Type", "application/json")
		var reqParams params
		if r.Method == "GET" {
			reqParams.Query = r.URL.Query().Get("query")
			reqParams.OperationName = r.URL.Query().Get("operationName")

			if variables := r.URL.Query().Get("variables"); variables != "" {
				if err := json.Unmarshal([]byte(variables), &reqParams.Variables); err != nil {
					sendErrorf(w, http.StatusBadRequest, "variables could not be decoded")
					return
				}
			}
		} else if r.Method == "OPTIONS" { // preflight check
			w.WriteHeader(http.StatusOK)
			return
		} else {
			if err := json.NewDecoder(r.Body).Decode(&reqParams); err != nil {
				sendErrorf(w, http.StatusBadRequest, "json body could not be decoded: "+err.Error())
				return
			}
		}

		doc, qErr := query.Parse(reqParams.Query)
		if qErr != nil {
			sendError(w, http.StatusUnprocessableEntity, qErr)
			return
		}

		errs := validation.Validate(exec.Schema(), doc)
		if len(errs) != 0 {
			sendError(w, http.StatusUnprocessableEntity, errs...)
			return
		}

		op, err := doc.GetOperation(reqParams.OperationName)
		if err != nil {
			sendErrorf(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		switch op.Type {
		case query.Query:
			b, err := json.Marshal(exec.Query(r.Context(), doc, reqParams.Variables, op, errorRecover))
			if err != nil {
				panic(err)
			}
			w.Write(b)
		case query.Mutation:
			b, err := json.Marshal(exec.Mutation(r.Context(), doc, reqParams.Variables, op, errorRecover))
			if err != nil {
				panic(err)
			}
			w.Write(b)
		default:
			sendErrorf(w, http.StatusBadRequest, "unsupported operation type")
		}
	})
}

func errorRecover(err interface{}) error {
	log.Println("Error encountered")
	return nil
}

// util
func getFrontendUrl() string {
	if os.Getenv("APP_ENV") == "production" {
		return "http://localhost:3000" // change this to production domain
	} else {
		return "http://localhost:3000"
	}
}

func setCors(w http.ResponseWriter) {
	frontendUrl := getFrontendUrl()
	w.Header().Set("Access-Control-Allow-Origin", frontendUrl)
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func sendError(w http.ResponseWriter, code int, errs ...*errors.QueryError) {
	w.WriteHeader(code)
	b, err := json.Marshal(&graphql.Response{Errors: errs})
	if err != nil {
		panic(err)
	}
	w.Write(b)
}

func sendErrorf(w http.ResponseWriter, code int, format string, args ...interface{}) {
	sendError(w, code, &errors.QueryError{Message: fmt.Sprintf(format, args...)})
}
