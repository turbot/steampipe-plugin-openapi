package openapi

import (
	"context"
	"os"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/getkin/kin-openapi/openapi3"
)

//// TABLE DEFINITION

func tableOpenAPIComponentSecurityScheme(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openapi_component_security_scheme",
		Description: "Components security scheme object.",
		List: &plugin.ListConfig{
			ParentHydrate: listOpenAPIFiles,
			Hydrate:       listOpenAPIComponentSecuritySchemes,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: openAPICommonColumns([]*plugin.Column{
			{Name: "key", Description: "The key used to refer or search the security scheme.", Type: proto.ColumnType_STRING},
			{Name: "name", Description: "The name of the header, query or cookie parameter to be used.", Type: proto.ColumnType_STRING},
			{Name: "type", Description: "The type of the security scheme. Valid values are apiKey, http, mutualTLS, oauth2, openIdConnect.", Type: proto.ColumnType_STRING},
			{Name: "location", Description: "The location of the API key. Possible values are query, header or cookie.", Type: proto.ColumnType_STRING, Transform: transform.FromField("In")},
			{Name: "description", Description: "A description for security scheme.", Type: proto.ColumnType_STRING},
			{Name: "scheme", Description: "The name of the HTTP Authorization scheme to be used in the Authorization header as defined in [RFC7235]. The values used SHOULD be registered in the IANA Authentication Scheme registry.", Type: proto.ColumnType_STRING},
			{Name: "bearer_format", Description: "A hint to the client to identify how the bearer token is formatted.", Type: proto.ColumnType_STRING},
			{Name: "open_id_connect_url", Description: "OpenId Connect URL to discover OAuth2 configuration values.", Type: proto.ColumnType_STRING},
			{Name: "flows", Description: "An object containing configuration information for the flow types supported.", Type: proto.ColumnType_JSON},
			{Name: "path", Description: "Path to the file.", Type: proto.ColumnType_STRING},
		}),
	}
}

type openAPIComponentSecurityScheme struct {
	Path      string
	Key       string
	StartLine int
	EndLine   int
	openapi3.SecurityScheme
}

//// LIST FUNCTION

func listOpenAPIComponentSecuritySchemes(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	file, err := os.Open(path)
	if err != nil {
		plugin.Logger(ctx).Error("openapi_component_security_scheme.listOpenAPIComponentSecuritySchemes", "file_open_error", err)
		return nil, err
	}

	// Get the parsed contents
	doc, err := getDoc(ctx, d, path)
	if err != nil {
		plugin.Logger(ctx).Error("openapi_component_security_scheme.listOpenAPIComponentSecuritySchemes", "parse_error", err)
		return nil, err
	}

	// Return nil, if no security schemes object defined
	if doc.Components == nil || doc.Components.SecuritySchemes == nil {
		return nil, nil
	}

	// For each security scheme, scan its arguments
	for k, v := range doc.Components.SecuritySchemes {

		// fetch start and end line for each securitySchemes
		var startLine, endLine int
		if strings.HasSuffix(path, "json") {
			startLine, endLine = findBlockLinesFromJSON(file, "components", k)
		} else {
			startLine, endLine = findBlockLinesFromYML(file, "components", k)
		}

		d.StreamListItem(ctx, openAPIComponentSecurityScheme{path, k, startLine, endLine, *v.Value})

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
