package openapi

import (
	"context"

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
			Hydrate:       listOpenAPIInfo,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: []*plugin.Column{
			{Name: "title", Description: "The title of the API.", Type: proto.ColumnType_STRING},
			{Name: "description", Description: "A description of the API.", Type: proto.ColumnType_STRING},
			{Name: "version", Description: "The version of the OpenAPI document.", Type: proto.ColumnType_STRING},
			{Name: "terms_of_service", Description: "A URL to the Terms of Service for the API.", Type: proto.ColumnType_STRING},
			{Name: "contact", Description: "The contact information for the exposed API.", Type: proto.ColumnType_JSON},
			{Name: "license", Description: "The license information for the exposed API.", Type: proto.ColumnType_JSON},
			{Name: "specification_version", Description: "The version of the OpenAPI specification.", Type: proto.ColumnType_STRING},
			{Name: "path", Description: "Path to the file.", Type: proto.ColumnType_STRING},
		},
	}
}

type openAPIInfo struct {
	Path                 string
	SpecificationVersion string
	openapi3.Info
}

//// LIST FUNCTION

func listOpenAPIInfo(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	doc, err := getDoc(ctx, d, path)
	if err != nil {
		return nil, err
	}
	d.StreamListItem(ctx, openAPIInfo{path, doc.OpenAPI, *doc.Info})

	return nil, nil
}
