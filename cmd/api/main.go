package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jdpx/mind-hub-api/pkg/api"
)

const graphCMSURLKey = "GRAPH_CMS_URL"

func main() {
	fmt.Println("Start Graphql API")

	c := api.Config{
		GraphCMSURL: os.Getenv(graphCMSURLKey),
	}

	router := api.NewRouter(&c)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
