package openapi

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type openAPIConfig struct {
	Paths []string `hcl:"paths,optional" steampipe:"watch"`
}

func ConfigInstance() interface{} {
	return &openAPIConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) openAPIConfig {
	if connection == nil || connection.GetConfig() == nil {
		return openAPIConfig{}
	}
	config, _ := connection.GetConfig().(openAPIConfig)
	return config
}
