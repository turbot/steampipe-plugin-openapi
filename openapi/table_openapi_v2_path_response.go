package openapi

import (
	"context"
	p "path"
	"strings"

	"github.com/go-openapi/spec"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableOpenAPIV2PathResponse(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openapi_v2_path_response",
		Description: "Path response object.",
		List: &plugin.ListConfig{
			ParentHydrate: listOpenAPIFiles,
			Hydrate:       listOpenAPIV2PathResponses,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "api_path",
				Description: "The key of the response object definition.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "api_method",
				Description: "Specifies the HTTP method.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "response_status",
				Description: "The key of the response object definition.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "response_ref",
				Description: "The reference to the paths response object.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "headers",
				Description: "Maps a header name to its definition.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "examples",
				Description: "Examples associated to the response.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "schema",
				Description: "The response schema.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "extensions",
				Description: "Maps a header name to its definition.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "description",
				Description: "A description of the response.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "path",
				Description: "Path to the file.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

type openAPIV2PathResponse struct {
	Path           string
	ApiPath        string
	ApiMethod      string
	ResponseStatus int
	ResponseRef    spec.Ref
	Description    string
	Headers        map[string]spec.Header
	Examples       map[string]interface{}
	Schema         *spec.Schema
	Extensions     spec.Extensions
}

//// LIST FUNCTION

func listOpenAPIV2PathResponses(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	// Get the parsed contents
	doc, err := getV2Doc(ctx, d, path)
	if err != nil {
		plugin.Logger(ctx).Error("openapi_v2_parameter.listOpenAPIV2Parameters", "parse_error", err)
		return nil, err
	}

	// check if the doc is swagger 2.0
	if doc.Swagger != "2.0" {
		return nil, nil
	}

	// For each path, scan its response object arguments
	for apiPath, item := range doc.Paths.Paths {
		for _, op := range OperationTypes {
			operation := getV2OperationInfoByType(op, item)

			// Skip if no method defined
			if operation == nil || operation.Responses == nil {
				continue
			}

			for responseStatus, response := range operation.Responses.StatusCodeResponses {
				responseObject := openAPIV2PathResponse{
					Path:           path,
					ApiPath:        p.Join(apiPath, op),
					ApiMethod:      strings.ToUpper(op),
					ResponseStatus: responseStatus,
					Description:    response.Description,
					ResponseRef:    response.Ref,
					Headers:        response.Headers,
					Examples:       response.Examples,
					Schema:         response.Schema,
					Extensions:     response.Extensions,
				}
				d.StreamListItem(ctx, responseObject)

				// Context may get cancelled due to manual cancellation or if the limit has been reached
				if d.RowsRemaining(ctx) == 0 {
					return nil, nil
				}
			}
		}
	}

	return nil, nil
}
