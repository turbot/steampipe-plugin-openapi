package openapi

import (
	"context"

	"github.com/go-openapi/spec"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableOpenAPIV2Response(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openapi_v2_response",
		Description: "Path object specified in OpenAPI V2 specification file.",
		List: &plugin.ListConfig{
			ParentHydrate: listOpenAPIFiles,
			Hydrate:       listOpenAPIV2Responses,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "description",
				Description: "A brief description of the response.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "schema",
				Description: "The schema defining the type used for the response.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "headers",
				Description: "A map of headers that should be included in the response.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "examples",
				Description: "Examples of the response message.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "extensions",
				Description: "Custom extensions for the Response.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "ref",
				Description: "A reference to an external definition that replaces this definition.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "path",
				Description: "Path to the file.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "key",
				Description: "The key used to refer or search the response.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

type openAPIV2Response struct {
	Path string
	Key  string
	spec.Response
}

//// LIST FUNCTION

func listOpenAPIV2Responses(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	// Get the parsed contents
	doc, err := getV2Doc(ctx, d, path)
	if err != nil {
		plugin.Logger(ctx).Error("openapi_v2_response.listOpenAPIV2Responses", "parse_error", err)
		return nil, err
	}

	// check if the doc is swagger 2.0
	if doc.Swagger != "2.0" {
		return nil, nil
	}

	// For each path, scan its arguments
	for key, item := range doc.Responses {
		d.StreamListItem(ctx, openAPIV2Response{path, key, item})

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
