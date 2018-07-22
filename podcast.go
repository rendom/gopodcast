package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rendom/gopodcast/resolver"
	"github.com/rendom/gopodcast/service"
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

	config := GetConfig()

	db, err := sqlx.Connect("sqlite3", config.DBFile)
	if err != nil {
		log.Fatalln(err)
	}

	sqlMigrate(db)

	if config.Debug {
		fmt.Printf("Config: %+v\n", config)
	}

	authService := &service.AuthService{
		PubKey:  config.JWTPublicKeyFile,
		PrivKey: config.JWTPrivateKeyFile,
		Expire:  config.JWTExpire,
	}

	schema := graphql.MustParseSchema(s, &resolver.Resolver{
		UserService: &service.User{
			DB: db,
		},
		AuthService: authService,
		PodcastService: &service.Podcast{
			DB: db,
		},
		EpisodeService: &service.Episode{
			DB: db,
		},
		SubscriptionService: &service.Subscription{
			DB: db,
		},
	})

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "graphiql.html")
	}))
	http.Handle("/query", authService.Middleware(
		&relay.Handler{Schema: schema},
	))

	addr := "localhost:8080"
	fmt.Println(addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}

}
