package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct{
	AccountUrl string `envConfig:"ACCOUNT_SERVICE_URL"`
	CatalogUrl string `envConfig:"CATALOG_SERVICE_URL"`
	OrderUrl string `envConfig:"ORDER_SERVICE_URL"`
}

func main() {
	var cfg AppConfig
	err := envconfig.Process("app",&cfg)
	if err!=nil {
		log.Fatal(err)
	}

	s,err:= NewGraphqlServer(cfg.AccountUrl,cfg.CatalogUrl,cfg.OrderUrl)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/graphql",handler.New(s.ToExecutableSchema()))
	http.Handle("/playground",playground.Handler("GraphQL playground", "/graphql"))
}