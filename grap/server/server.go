package main

import (
	"net/http"
	
	"log"
	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	grap "github.com/siulfe/gql"
	DDBB "github.com/siulfe/gql/Database"
)

const defaultPort = "8080"

func main() {
	resp,err := DDBB.LeerArchivo()
	log.Println("Respuesta: ",resp)
	if err != nil{
		panic(err)
	}

	port := resp["port"]
	if port == "" {
		port = defaultPort
	}
	
	err = DDBB.ConnectDatabase(resp)

	if err != nil{
		panic(err)
	}

	router := chi.NewRouter()

	router.Use(grap.Middleware())

	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", handler.GraphQL(grap.NewExecutableSchema(grap.Config{Resolvers: &grap.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

	go grap.DesLoged()

}
