package openapi

import (
	"context"
	p "path"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableOpenAPIPathResponse(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openapi_path_response",
		Description: "Describes the response from an API Operation",
		List: &plugin.ListConfig{
			ParentHydrate: listFiles,
			Hydrate:       listOpenAPIPathResponses,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: []*plugin.Column{
			{Name: "api_path", Description: "The key of the response object definition.", Type: proto.ColumnType_STRING},
			{Name: "response_status", Description: "The key of the response object definition.", Type: proto.ColumnType_STRING},
			{Name: "response_ref", Description: "The reference to the components response object.", Type: proto.ColumnType_STRING},
			{Name: "content", Description: "A map containing descriptions of potential response payloads.", Type: proto.ColumnType_JSON},
			{Name: "headers", Description: "Maps a header name to its definition.", Type: proto.ColumnType_JSON, Transform: transform.FromField("Raw.Headers")},
			{Name: "links", Description: "A map of operations links that can be followed from the response.", Type: proto.ColumnType_JSON, Transform: transform.FromField("Raw.Links")},
			{Name: "description", Description: "A description of the response.", Type: proto.ColumnType_STRING},
			{Name: "path", Description: "Path to the file.", Type: proto.ColumnType_STRING},
		},
	}
}

type openAPIPathResponse struct {
	Path           string
	ApiPath        string
	ResponseStatus string
	ResponseRef    string
	Content        []map[string]interface{}
	Description    string
	Raw            openapi3.Response
}

//// LIST FUNCTION

func listOpenAPIPathResponses(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	doc, err := getDoc(ctx, d, path)
	if err != nil {
		return nil, err
	}

	operationTypes := []string{"connect", "delete", "get", "head", "options", "patch", "post", "put", "trace"}
	for apiPath, item := range doc.Paths {
		for _, op := range operationTypes {
			operation := getOperationInfoByType(op, item)

			// Skip if no method defined
			if operation == nil {
				continue
			}

			for responseStatus, response := range operation.Responses {
				responseObject := openAPIPathResponse{
					Path:           path,
					ApiPath:        p.Join(apiPath, op),
					ResponseStatus: responseStatus,
					Description:    *response.Value.Description,
					ResponseRef:    response.Ref,
				}

				for header, content := range response.Value.Content {
					var schema interface{}
					if content.Schema.Ref != "" {
						schema = content.Schema.Ref
					} else {
						schema = content.Schema
					}
					responseObject.Content = append(responseObject.Content, map[string]interface{}{
						"contentType": header,
						"examples":    content.Examples,
						"schema":      schema,
						"schemaType":  content.Schema.Value.Type,
					})
				}
				responseObject.Raw = *response.Value

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
