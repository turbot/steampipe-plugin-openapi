package openapi

import (
	"context"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableOpenAPIComponentResponse(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openapi_component_response",
		Description: "Describes the response from an API Operation",
		List: &plugin.ListConfig{
			ParentHydrate: listFiles,
			Hydrate:       listOpenAPIComponentResponses,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: []*plugin.Column{
			{Name: "key", Description: "The key of the response object definition.", Type: proto.ColumnType_STRING},
			{Name: "description", Description: "A description of the response.", Type: proto.ColumnType_STRING},
			{Name: "content", Description: "A map containing descriptions of potential response payloads.", Type: proto.ColumnType_JSON},
			{Name: "headers", Description: "Maps a header name to its definition.", Type: proto.ColumnType_JSON, Transform: transform.FromField("Raw.Headers")},
			{Name: "links", Description: "A map of operations links that can be followed from the response.", Type: proto.ColumnType_JSON, Transform: transform.FromField("Raw.Links")},
			{Name: "path", Description: "Path to the file.", Type: proto.ColumnType_STRING},
		},
	}
}

type openAPIComponentResponse struct {
	Path        string
	Content     []map[string]interface{}
	Key         string
	Description string
	Raw         openapi3.Response
}

//// LIST FUNCTION

func listOpenAPIComponentResponses(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	doc, err := getDoc(ctx, d, path)
	if err != nil {
		return nil, err
	}

	// Return nil, if no schema defined
	if doc.Components == nil || doc.Components.Responses == nil {
		return nil, nil
	}

	for k, v := range doc.Components.Responses {
		responseObject := openAPIComponentResponse{
			Path:        path,
			Key:         k,
			Description: *v.Value.Description,
		}

		for header, content := range v.Value.Content {
			responseObject.Content = append(responseObject.Content, map[string]interface{}{
				"contentType": header,
				"examples":    content.Examples,
				"schema":      content.Schema,
				"schemaType":  content.Schema.Value.Type,
			})
		}
		responseObject.Raw = *v.Value
		d.StreamListItem(ctx, responseObject)

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
