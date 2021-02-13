package logging

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/jdpx/mind-hub-api/pkg/request"
	"github.com/sirupsen/logrus"
)

// HTTPTransportLogger ...
type HTTPTransportLogger struct {
	next       http.RoundTripper
	dst        io.Writer
	clientName string
}

func NewHTTPTransportLogger(name string) *HTTPTransportLogger {
	return &HTTPTransportLogger{
		next:       http.DefaultTransport,
		dst:        os.Stderr,
		clientName: name,
	}
}

func (l *HTTPTransportLogger) RoundTrip(req *http.Request) (res *http.Response, err error) {
	cID := request.GetCorrelationIDHeader(req.Header)
	log := New().WithFields(logrus.Fields{
		CorrelationIDKey:  cID,
		HTTPURLKey:        req.URL.Path,
		HTTPMethodKey:     req.Method,
		HTTPClientNameKey: l.clientName,
	})

	log.Info(fmt.Sprintf("%s Request starting", l.clientName))

	defer func(begin time.Time) {
		log.WithFields(logrus.Fields{
			RequestDurationKey: time.Since(begin).Milliseconds(),
			HTTPStatusKey:      res.StatusCode,
		}).Info(fmt.Sprintf("%s Request completed ", l.clientName))
	}(time.Now())

	return l.next.RoundTrip(req)
}
