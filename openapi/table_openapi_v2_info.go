package openapi

import (
	"context"

	"github.com/go-openapi/spec"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableOpenAPIV2Info(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openapi_v2_info",
		Description: "Info object specified in OpenAPI specification file.",
		List: &plugin.ListConfig{
			ParentHydrate: listOpenAPIFiles,
			Hydrate:       listOpenAPIV2Info,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "title",
				Description: "The title of the API.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "A description of the API.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "version",
				Description: "The version of the OpenAPI document.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "terms_of_service",
				Description: "A URL to the Terms of Service for the API.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "contact",
				Description: "The contact information for the exposed API.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "license",
				Description: "The license information for the exposed API.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "specification_version",
				Description: "The version of the OpenAPI specification.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "path",
				Description: "Path to the file.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "global_schemes",
				Description: "The global schemes.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

type openAPIV2Info struct {
	Path                 string
	SpecificationVersion string
	GlobalSchemes        []string
	spec.Info
}

//// LIST FUNCTION

func listOpenAPIV2Info(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	// Get the parsed contents
	doc, err := getV2Doc(ctx, d, path)
	if err != nil {
		plugin.Logger(ctx).Error("openapi_v2_info.listOpenAPIV2Info", "parse_error", err)
		return nil, err
	}

	// check if the doc is swagger 2.0
	if doc.Swagger != "2.0" {
		return nil, nil
	}

	d.StreamListItem(ctx, openAPIV2Info{path, doc.Swagger, doc.Schemes, *doc.Info})

	// Context may get cancelled due to manual cancellation or if the limit has been reached
	if d.RowsRemaining(ctx) == 0 {
		return nil, nil
	}

	return nil, nil
}
