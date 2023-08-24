package openapi

import (
	"context"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	"github.com/getkin/kin-openapi/openapi3"
)

//// TABLE DEFINITION

func tableOpenAPIInfo(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openapi_info",
		Description: "Info object specified in OpenAPI specification file.",
		List: &plugin.ListConfig{
			ParentHydrate: listOpenAPIFiles,
			Hydrate:       listOpenAPIInfo,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: openAPICommonColumns([]*plugin.Column{
			{Name: "title", Description: "The title of the API.", Type: proto.ColumnType_STRING},
			{Name: "description", Description: "A description of the API.", Type: proto.ColumnType_STRING},
			{Name: "version", Description: "The version of the OpenAPI document.", Type: proto.ColumnType_STRING},
			{Name: "terms_of_service", Description: "A URL to the Terms of Service for the API.", Type: proto.ColumnType_STRING},
			{Name: "contact", Description: "The contact information for the exposed API.", Type: proto.ColumnType_JSON},
			{Name: "license", Description: "The license information for the exposed API.", Type: proto.ColumnType_JSON},
			{Name: "specification_version", Description: "The version of the OpenAPI specification.", Type: proto.ColumnType_STRING},
			{Name: "path", Description: "Path to the file.", Type: proto.ColumnType_STRING},
		}),
	}
}

type openAPIInfo struct {
	Path                 string
	SpecificationVersion string
	StartLine            int
	EndLine              int
	openapi3.Info
}

//// LIST FUNCTION

func listOpenAPIInfo(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	// get the start and end lines of the info block
	file, err := os.Open(path)
	if err != nil {
		plugin.Logger(ctx).Error("openapi_info.listOpenAPIInfo", "file_open_error", err)
		return nil, err
	}
	startLine, endLine := findBlockLines(file, "info", "")

	// Get the parsed contents
	doc, err := getDoc(ctx, d, path)
	if err != nil {
		plugin.Logger(ctx).Error("openapi_info.listOpenAPIInfo", "parse_error", err)
		return nil, err
	}
	d.StreamListItem(ctx, openAPIInfo{path, doc.OpenAPI, startLine, endLine, *doc.Info})

	// Context may get cancelled due to manual cancellation or if the limit has been reached
	if d.RowsRemaining(ctx) == 0 {
		return nil, nil
	}

	return nil, nil
}
