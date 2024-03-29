package openapi

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-openapi"

// Plugin creates this (openapi) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel().Transform(transform.NullIfZeroValue),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		TableMap: map[string]*plugin.Table{
			"openapi_component_header":          tableOpenAPIComponentHeader(ctx),
			"openapi_component_parameter":       tableOpenAPIComponentParameter(ctx),
			"openapi_component_request_body":    tableOpenAPIComponentRequestBody(ctx),
			"openapi_component_response":        tableOpenAPIComponentResponse(ctx),
			"openapi_component_schema":          tableOpenAPIComponentSchema(ctx),
			"openapi_component_security_scheme": tableOpenAPIComponentSecurityScheme(ctx),
			"openapi_info":                      tableOpenAPIInfo(ctx),
			"openapi_path":                      tableOpenAPIPath(ctx),
			"openapi_path_request_body":         tableOpenAPIPathRequestBody(ctx),
			"openapi_path_response":             tableOpenAPIPathResponse(ctx),
			"openapi_server":                    tableOpenAPIServer(ctx),
		},
	}

	return p
}
