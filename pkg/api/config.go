package api

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config ...
type Config struct {
	Version     string            `ignored:"true"`
	Env         string            `default:"local"`
	CMSMapping  map[string]string `envconfig:"MIND_GRAPH_CMS_URL_MAPPING" required:"true"`
	GraphCMSURL string
}

func NewConfig() Config {
	var c Config
	err := envconfig.Process("mind", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	return c
}
