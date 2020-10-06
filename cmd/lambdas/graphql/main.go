package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	"github.com/jdpx/mind-hub-api/pkg/graphql/graph"
	"github.com/jdpx/mind-hub-api/pkg/graphql/graph/generated"
)

var ginLambda *ginadapter.GinLambda

type Config struct {
	graphcmsURL string
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {

	c := Config{
		graphcmsURL: os.Getenv("GRAPH_CMS_URL"),
	}

	fmt.Println(c.graphcmsURL)

	client := graphcms.NewClient(c.graphcmsURL)
	resolver := graph.NewResolver(
		graph.WithClient(client),
	)

	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Handler ...
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	if ginLambda == nil {
		// stdout and stderr are sent to AWS CloudWatch Logs
		log.Printf("Gin cold start")
		r := gin.Default()

		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "PUT", "OPTIONS", "PATCH", "DELETE"},
			AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Methods", "Access-Control-Allow-Headers"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))

		// Setting up Gin
		r.POST("v1/query", graphqlHandler())
		r.GET("v1/playground", playgroundHandler())
		r.GET("v1/ping", func(c *gin.Context) {
			log.Println("Handler!!")
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		ginLambda = ginadapter.New(r)
	}

	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	// Setting up Gin
	lambda.Start(Handler)
}
