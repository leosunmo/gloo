package runner

import (
	"github.com/kelseyhightower/envconfig"
)

type Settings struct {
	GlooAddress string `envconfig:"GLOO_ADDRESS" default:"control-plane:8080"`
	DebugPort   int    `envconfig:"DEBUG_PORT" default:"9091"`
	ServerPort  int    `envconfig:"SERVER_PORT" default:"8083"`
}

func NewSettings() Settings {
	var s Settings

	err := envconfig.Process("", &s)
	if err != nil {
		panic(err)
	}

	return s
}
