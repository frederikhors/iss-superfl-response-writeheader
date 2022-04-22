package main

import (
	"iss-superfl-response-writeheader/graph"
	"iss-superfl-response-writeheader/graph/generated"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

func main() {
	router := bunrouter.New(bunrouter.Use(reqlog.NewMiddleware()))

	router.GET("/", indexHandler)

	gqlgenHandler := newGqlgenHandler()

	router.WithGroup("/graphql", func(group *bunrouter.Group) {
		group.GET("", bunrouter.HTTPHandler(gqlgenHandler))
		group.POST("", bunrouter.HTTPHandler(gqlgenHandler))
	})

	router.WithGroup("/playground", func(group *bunrouter.Group) {
		group.GET("", bunrouter.HTTPHandler(playground.Handler("temp", "/graphql")))
	})

	log.Println("listening on http://localhost:9999")
	log.Println(http.ListenAndServe(":9999", router))
}

func indexHandler(w http.ResponseWriter, req bunrouter.Request) error {
	return bunrouter.JSON(w, bunrouter.H{"ping": "pong"})
}

func newGqlgenHandler() *handler.Server {
	srv := handler.New(generated.NewExecutableSchema(
		generated.Config{Resolvers: &graph.Resolver{}}),
	)

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// return true
				return false
			},
		}},
	)

	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.SetQueryCache(lru.New(1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{Cache: lru.New(100)})

	return srv
}
