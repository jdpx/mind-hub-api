package api

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	"github.com/jdpx/mind-hub-api/pkg/graphql"
	"github.com/jdpx/mind-hub-api/pkg/graphql/generated"
	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/jdpx/mind-hub-api/pkg/request"
	"github.com/jdpx/mind-hub-api/pkg/service"
	"github.com/jdpx/mind-hub-api/pkg/store"
	graphqlClient "github.com/machinebox/graphql"
)

// Config ...
type Config struct {
	Env         string
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
	graphqlClient := graphqlClient.NewClient(
		config.GraphCMSURL,
		graphqlClient.WithHTTPClient(graphcms.DefaultHTTPClient()),
	)

	cms := graphcms.NewClient(graphqlClient)
	cmsResolver := graphcms.NewResolver(cms)

	var dynamoClient *dynamodb.Client
	if config.Env == "local" {
		dynamoClient = store.NewLocalClient()
	} else {
		dynamoClient = store.NewClient()
	}

	s := store.NewStore(
		store.WithDynamoDB(dynamoClient),
	)

	noteStore := store.NewNoteStore(s)
	progressStore := store.NewProgressStore(s, store.GenerateID)
	timemapStore := store.NewTimemapStore(s, store.GenerateID)

	courseProgressService := service.NewCourseProgressService(cmsResolver, progressStore)
	courseService := service.NewCourseService(cmsResolver)
	sessionService := service.NewSessionService(cmsResolver)
	courseNoteService := service.NewCourseNoteService(noteStore)
	stepProgressService := service.NewStepProgressService(progressStore)
	stepService := service.NewStepService(cmsResolver)
	stepNoteService := service.NewStepNoteService(noteStore)
	timemapService := service.NewTimemapService(timemapStore)

	serv := service.New(
		service.WithCourse(courseService),
		service.WithCourseProgress(courseProgressService),
		service.WithCourseNote(courseNoteService),
		service.WithSession(sessionService),
		service.WithStep(stepService),
		service.WithStepNote(stepNoteService),
		service.WithStepProgress(stepProgressService),
		service.WithTimemap(timemapService),
	)

	resolver := graphql.NewResolver(
		graphql.WithService(serv),
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
