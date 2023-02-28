package openapi

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	"github.com/getkin/kin-openapi/openapi3"
)

//// TABLE DEFINITION

func tableOpenAPIInfo(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openapi_info",
		Description: "Info specified by OpenAPI/Swagger standard version 3",
		List: &plugin.ListConfig{
			ParentHydrate: listFiles,
			Hydrate:       listInfo,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: []*plugin.Column{
			{Name: "title", Description: "The title of the API.", Type: proto.ColumnType_STRING},
			{Name: "description", Description: "A description of the API.", Type: proto.ColumnType_STRING},
			{Name: "version", Description: "The version of the OpenAPI document.", Type: proto.ColumnType_STRING},
			{Name: "terms_of_service", Description: "A URL to the Terms of Service for the API.", Type: proto.ColumnType_STRING},
			{Name: "contact", Description: "The contact information for the exposed API.", Type: proto.ColumnType_JSON},
			{Name: "license", Description: "The license information for the exposed API.", Type: proto.ColumnType_JSON},
			{Name: "path", Description: "Path to the file.", Type: proto.ColumnType_STRING},
		},
	}
}

type openAPIInfo struct {
	Path string
	openapi3.Info
}

//// LIST FUNCTION

func listInfo(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	doc, err := openapi3.NewLoader().LoadFromFile(path)
	if err != nil {
		plugin.Logger(ctx).Error("listInfo", "file_error", err, "path", path)
		return nil, fmt.Errorf("failed to load file %s: %v", path, err)
	}
	d.StreamListItem(ctx, openAPIInfo{path, *doc.Info})

	return nil, nil
}
