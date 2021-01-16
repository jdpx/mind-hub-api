package request

import "net/http"

const (
	correlationIDHeader = "X-Correlation-Id"
	authorizationHeader = "Authorization"
)

func GetHeader(key string, header http.Header) string {
	return header.Get(key)
}

func GetCorrelationIDHeader(header http.Header) string {
	return GetHeader(correlationIDHeader, header)
}

func GetAuthorizationHeader(header http.Header) string {
	return GetHeader(authorizationHeader, header)
}
