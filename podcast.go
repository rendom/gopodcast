package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/rendom/gopodcast/resolver"
)

var schema *graphql.Schema

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func main() {

	s, err := getSchema("./schema.graphql")
	if err != nil {
		log.Panic(err)
	}

	schema := graphql.MustParseSchema(s, &resolver.Resolver{})

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "graphiql.html")
	}))

	http.Handle("/query", &relay.Handler{Schema: schema})

	addr := "localhost:8080"
	fmt.Println(addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}

}
