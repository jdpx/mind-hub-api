package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/jdpx/mind-hub-api/pkg/api"
	"github.com/jdpx/mind-hub-api/pkg/logging"
)

var ginLambda *ginadapter.GinLambda

var buildVersion = "0.0.1"

// Handler ...
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		logging.New().Info("Start Graphql Lambda API", buildVersion)

		c := api.NewConfig()
		c.Version = buildVersion

		router := api.NewRouter(c)

		ginLambda = ginadapter.New(router)
	}

	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
