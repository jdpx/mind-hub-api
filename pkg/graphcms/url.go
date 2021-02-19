package graphcms

import "fmt"

const (
	graphCMSBaseURL = "https://api-eu-central-1.graphcms.com/v2/%s"
)

func NewCMSUrl(url string) string {
	return fmt.Sprintf(graphCMSBaseURL, url)
}
