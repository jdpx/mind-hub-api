package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/jdpx/mind-hub-api/pkg/api"
)

var ginLambda *ginadapter.GinLambda

const graphCMSURLKey = "GRAPH_CMS_URL"

// Handler ...
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		fmt.Println("Start Graphql Lambda API")

		c := api.Config{
			Env:         "prod",
			GraphCMSURL: os.Getenv(graphCMSURLKey),
		}

		router := api.NewRouter(&c)

		ginLambda = ginadapter.New(router)
	}

	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
