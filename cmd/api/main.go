package main

import (
	"net/http"

	"github.com/jdpx/mind-hub-api/pkg/api"
	"github.com/jdpx/mind-hub-api/pkg/logging"
)

const apiPort = ":8080"

var buildVersion = "0.0.1"

func main() {
	logging.New().Info("Start Local Graphql API", buildVersion)

	c := api.NewConfig()
	c.Version = buildVersion

	router := api.NewRouter(c)

	err := http.ListenAndServe(apiPort, router)
	if err != nil {
		panic(err)
	}
}
