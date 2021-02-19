package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/jdpx/mind-hub-api/pkg/request"
	"github.com/sirupsen/logrus"
)

// NewRouter ...
func NewRouter(config *Config) *gin.Engine {
	log := logging.New()

	log = log.WithFields(
		logrus.Fields{
			logging.APIVersionKey:     config.Version,
			logging.APIEnvironmentKey: config.Env,
			logging.APICMSMapping:     config.CMSMapping,
		},
	)

	log.Info("Starting Gin Router")
	r := gin.Default()

	r.Use(request.CORSMiddleware())
	r.Use(logging.GinRequestLoggerMiddleware())
	r.Use(request.ContextMiddleware())
	r.Use(request.VersionMiddleware(config.Version))

	r.POST("v1/query", graphqlHandler(config))

	return r
}
