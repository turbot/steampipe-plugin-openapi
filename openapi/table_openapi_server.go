package openapi

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/getkin/kin-openapi/openapi3"
)

//// TABLE DEFINITION

func tableOpenAPIServer(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openapi_server",
		Description: "Server specified by OpenAPI/Swagger standard version 3",
		List: &plugin.ListConfig{
			ParentHydrate: listFiles,
			Hydrate:       listOpenAPIServers,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: []*plugin.Column{
			{Name: "url", Description: "A URL to the target host.", Type: proto.ColumnType_STRING, Transform: transform.FromField("URL")},
			{Name: "description", Description: "An optional string describing the host designated by the URL.", Type: proto.ColumnType_STRING},
			{Name: "variables", Description: "A map between a variable name and its value, used for substitution in the server’s URL template.", Type: proto.ColumnType_JSON},
			{Name: "path", Description: "Path to the file.", Type: proto.ColumnType_STRING},
		},
	}
}

type openAPIServer struct {
	Path string
	openapi3.Server
}

//// LIST FUNCTION

func listOpenAPIServers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	doc, err := openapi3.NewLoader().LoadFromFile(path)
	if err != nil {
		plugin.Logger(ctx).Error("openapi_server.listOpenAPIServers", "file_error", err, "path", path)
		return nil, fmt.Errorf("failed to load file %s: %v", path, err)
	}

	for _, server := range doc.Servers {
		d.StreamListItem(ctx, openAPIServer{path, *server})
	}

	return nil, nil
}
