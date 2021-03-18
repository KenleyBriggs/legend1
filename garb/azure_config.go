package azuresqltarget

import (
	libadapter "github.com/triggermesh/knative-targets/pkg/adapter"
)

// EnvAccessorCtor for configuration parameters
func EnvAccessorCtor() libadapter.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	libadapter.EnvConfig

	ServerURL      string `envconfig:"AZURE_SQL_SERVER_URL" required:"true"`
	ServerPort     int    `envconfig:"AZURE_SQL_SERVER_PORT" required:"true"`
	ServerPassword string `envconfig:"AZURE_SQL_SERVER_PASSWORD" required:"true"`
	ServerUser     string `envconfig:"AZURE_SQL_SERVER_USER" required:"true"`
	ServerDatabase string `envconfig:"AZURE_SQL_SERVER_DATABASE" required:"true"`
}
