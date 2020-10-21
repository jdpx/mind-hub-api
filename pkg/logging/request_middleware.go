package logging

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// RequestLoggerMiddleware ...
func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cID := c.Request.Header.Get(correlationIDHeader)

		if cID == "" {
			id, _ := uuid.NewUUID()
			cID = id.String()
		}

		log := New().WithFields(logrus.Fields{
			CorrelationIDKey: cID,
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
