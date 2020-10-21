package api

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jdpx/mind-hub-api/pkg/logging"
	"github.com/sirupsen/logrus"
)

const (
	correlationIDHeader = "X-Correlation-Id"
)

// CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "OPTIONS", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "x-correlation-id"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Methods", "Access-Control-Allow-Headers", "x-correlation-id"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

// RequestLoggerMiddleware ...
func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cID := c.Request.Header.Get(correlationIDHeader)

		if cID == "" {
			id, _ := uuid.NewUUID()
			cID = id.String()
		}

		log := logging.New().WithFields(logrus.Fields{
			logging.CorrelationIDKey: cID,
		})

		log.Info(fmt.Sprintf("Request %s starting", c.Request.URL.Path))

		t := time.Now()

		c.Next()

		latency := time.Since(t).Milliseconds()
		status := c.Writer.Status()

		log.WithFields(logrus.Fields{
			logging.RequestDurationKey: latency,
			logging.HTTPStatusKey:      status,
		}).Info(fmt.Sprintf("Request %s completed ", c.Request.URL.Path))
	}
}
