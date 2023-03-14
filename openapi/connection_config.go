package openapi

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type openAPIConfig struct {
	Paths []string `cty:"paths" steampipe:"watch"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"paths": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{Type: schema.TypeString},
	},
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
