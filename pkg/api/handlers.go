package api

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
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

// Defining the Graphql handler
func graphqlHandler(config Config) gin.HandlerFunc {
	cmsClient := newCMSClient(config)
	storeClient := newStoreClient(config)

	cmsResolver := graphcms.NewResolver(cmsClient)
	noteStore := store.NewNoteStore(storeClient)
	progressStore := store.NewProgressStore(storeClient)
	timemapStore := store.NewTimemapStore(storeClient)

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

func newCMSClient(config Config) *graphcms.Client {
	log := logging.New()
	cmsHTTPClient := request.DefaultHTTPClient(
		request.WithTransport(logging.NewHTTPTransportLogger("GraphCMS")),
	)

	var cmsClients []graphcms.Option

	for key, cmsURL := range config.CMSMapping {
		log.Info(fmt.Sprintf("Registered Organisation %s GraphCMS Client", key))

		graphqlClient := graphqlClient.NewClient(
			graphcms.NewCMSUrl(cmsURL),
			graphqlClient.WithHTTPClient(cmsHTTPClient),
		)

		cmsClients = append(cmsClients, graphcms.WithOrganisationClient(key, graphqlClient))
	}

	return graphcms.NewClient(cmsClients...)
}

func newStoreClient(config Config) *store.Store {
	var dynamoClient *dynamodb.Client
	if config.IsLocal() {
		dynamoClient = store.NewLocalClient()
	} else {
		dynamoClient = store.NewClient()
	}

	s := store.NewStore(
		store.WithDynamoDB(dynamoClient),
	)

	return s
}
