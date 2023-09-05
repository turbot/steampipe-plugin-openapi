package openapi

import (
	"context"

	"github.com/go-openapi/spec"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableOpenAPIV2SecurityDefinition(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openapi_v2_security_definition",
		Description: "Path object specified in OpenAPI V2 specification file.",
		List: &plugin.ListConfig{
			ParentHydrate: listOpenAPIFiles,
			Hydrate:       listOpenAPIV2SecurityDefinitions,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "description",
				Description: "Detailed description of the security scheme.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "type",
				Description: "Type of the security scheme (e.g., 'apiKey', 'http', 'oauth2').",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "Name of the header, query or cookie parameter to be used (relevant for apiKey type).",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "in",
				Description: "Location of the API key (e.g., 'header', 'query').",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "flow",
				Description: "OAuth2 flow type (e.g., 'implicit', 'password', 'application', 'accessCode').",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "authorization_url",
				Description: "URL for OAuth2 authorization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "token_url",
				Description: "URL for obtaining the OAuth2 token.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "scopes",
				Description: "Scopes available for OAuth2.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "extensions",
				Description: "Custom extensions for the Security Scheme.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "path",
				Description: "Path to the file.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "key",
				Description: "The key used to refer or search the security definition.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

type openAPIV2SecurityDefinition struct {
	Path string
	Key  string
	spec.SecurityScheme
}

//// LIST FUNCTION

func listOpenAPIV2SecurityDefinitions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	// Get the parsed contents
	doc, err := getV2Doc(ctx, d, path)
	if err != nil {
		plugin.Logger(ctx).Error("openapi_v2_security_definition.listOpenAPIV2SecurityDefinitions", "parse_error", err)
		return nil, err
	}

	// check if the doc is swagger 2.0
	if doc.Swagger != "2.0" {
		return nil, nil
	}

	// For each path, scan its arguments
	for key, item := range doc.SecurityDefinitions {
		d.StreamListItem(ctx, openAPIV2SecurityDefinition{path, key, *item})

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
