package openapi

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type openAPIConfig struct {
	Paths []string `hcl:"paths" steampipe:"watch"`
}

func ConfigInstance() interface{} {
	return &openAPIConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) openAPIConfig {
	if connection == nil || connection.Config == nil {
		return openAPIConfig{}
	}
	config, _ := connection.Config.(openAPIConfig)
	return config
}
