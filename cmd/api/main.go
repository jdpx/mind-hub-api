package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jdpx/mind-hub-api/pkg/api"
)

const (
	graphCMSURLKey = "GRAPH_CMS_URL"
	environment    = "local"
)

var buildVersion = "0.0.1"

func main() {
	fmt.Println("Start Local Graphql API", buildVersion)

	c := api.Config{
		Version:     buildVersion,
		Env:         environment,
		GraphCMSURL: os.Getenv(graphCMSURLKey),
	}

	router := api.NewRouter(&c)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
