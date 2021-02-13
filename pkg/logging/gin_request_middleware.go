package logging

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jdpx/mind-hub-api/pkg/auth"
	"github.com/jdpx/mind-hub-api/pkg/request"
	"github.com/sirupsen/logrus"
)

// GinRequestLoggerMiddleware ...
func GinRequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cID := request.GetCorrelationIDHeader(c.Request.Header)

		if cID == "" {
			id, _ := uuid.NewUUID()
			cID = id.String()
		}

		aH, _ := request.GetGinContextAuthToken(c)
		orgID, _ := auth.GetTokenOrganisationID(aH)

		log := New().WithFields(logrus.Fields{
			CorrelationIDKey:  cID,
			OrganisationIDKey: orgID,
		})

		log.Info(fmt.Sprintf("Request %s starting", c.Request.URL.Path))

		t := time.Now()

		c.Next()

		latency := time.Since(t).Milliseconds()
		status := c.Writer.Status()

		log.WithFields(logrus.Fields{
			RequestDurationKey: latency,
			HTTPStatusKey:      status,
		}).Info(fmt.Sprintf("Request %s completed ", c.Request.URL.Path))
	}
}
