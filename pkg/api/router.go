package api

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	"github.com/jdpx/mind-hub-api/pkg/graphql/graph"
	"github.com/jdpx/mind-hub-api/pkg/graphql/graph/generated"
	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/jdpx/mind-hub-api/pkg/request"
	"github.com/machinebox/graphql"
)

// Config ...
type Config struct {
	GraphCMSURL string
}

// NewRouter ...
func NewRouter(config *Config) *gin.Engine {
	log := logging.New()

	log.Info("Gin cold start")
	r := gin.Default()

	r.Use(request.CORSMiddleware())
	r.Use(logging.RequestLoggerMiddleware())
	r.Use(request.ContextMiddleware())

	// Setting up Gin
	r.POST("v1/query", graphqlHandler(config))
	r.GET("v1/playground", playgroundHandler())
	r.GET("v1/ping", func(c *gin.Context) {
		log.Info("Handler!!")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}

// Defining the Graphql handler
func graphqlHandler(config *Config) gin.HandlerFunc {
	graphCMSClient := graphql.NewClient(config.GraphCMSURL)

	client := graphcms.NewClient(graphCMSClient)
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
