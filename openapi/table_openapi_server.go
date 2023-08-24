package openapi

import (
	"context"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/getkin/kin-openapi/openapi3"
)

//// TABLE DEFINITION

func tableOpenAPIServer(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openapi_server",
		Description: "Server object specified in OpenAPI specification file.",
		List: &plugin.ListConfig{
			ParentHydrate: listOpenAPIFiles,
			Hydrate:       listOpenAPIServers,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: openAPICommonColumns([]*plugin.Column{
			{Name: "url", Description: "A URL to the target host.", Type: proto.ColumnType_STRING, Transform: transform.FromField("URL")},
			{Name: "description", Description: "An optional string describing the host designated by the URL.", Type: proto.ColumnType_STRING},
			{Name: "variables", Description: "A map between a variable name and its value, used for substitution in the server's URL template.", Type: proto.ColumnType_JSON},
			{Name: "path", Description: "Path to the file.", Type: proto.ColumnType_STRING},
		}),
	}
}

type openAPIServer struct {
	Path      string
	StartLine int
	EndLine   int
	openapi3.Server
}

//// LIST FUNCTION

func listOpenAPIServers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	file, err := os.Open(path)
	if err != nil {
		plugin.Logger(ctx).Error("openapi_server.listOpenAPIServers", "file_open_error", err)
		return nil, err
	}

	// Get the parsed contents
	doc, err := getDoc(ctx, d, path)
	if err != nil {
		plugin.Logger(ctx).Error("openapi_server.listOpenAPIServers", "parse_error", err)
		return nil, err
	}

	// For each server, scan its arguments
	for _, server := range doc.Servers {
		startLine, endLine := findBlockLines(file, "servers", server.URL)
		d.StreamListItem(ctx, openAPIServer{path, startLine, endLine, *server})

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
