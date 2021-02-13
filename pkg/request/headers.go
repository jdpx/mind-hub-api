package request

import "net/http"

const (
	CorrelationIDHeader = "X-Correlation-Id"
	authorizationHeader = "Authorization"
)

func GetHeader(key string, header http.Header) string {
	return header.Get(key)
}

func GetCorrelationIDHeader(header http.Header) string {
	return GetHeader(CorrelationIDHeader, header)
}

func GetAuthorizationHeader(header http.Header) string {
	return GetHeader(authorizationHeader, header)
}
