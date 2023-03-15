package openapi

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	"github.com/getkin/kin-openapi/openapi3"
)

//// TABLE DEFINITION

func tableOpenAPIComponentRequestBody(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openapi_component_request_body",
		Description: "Components request body object.",
		List: &plugin.ListConfig{
			ParentHydrate: listOpenAPIFiles,
			Hydrate:       listOpenAPIComponentRequestBodies,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: []*plugin.Column{
			{Name: "key", Description: "The key used to refer or search the request body.", Type: proto.ColumnType_STRING},
			{Name: "description", Description: "A brief description of the request body.", Type: proto.ColumnType_STRING},
			{Name: "required", Description: "True, if the request body is required.", Type: proto.ColumnType_BOOL},
			{Name: "content", Description: "The content of the request body.", Type: proto.ColumnType_JSON},
			{Name: "path", Description: "Path to the file.", Type: proto.ColumnType_STRING},
		},
	}
}

type openAPIComponentRequestBody struct {
	Path    string
	Key     string
	Content []map[string]interface{}
	Raw     openapi3.RequestBody
}

//// LIST FUNCTION

func listOpenAPIComponentRequestBodies(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	doc, err := getDoc(ctx, d, path)
	if err != nil {
		plugin.Logger(ctx).Error("openapi_component_request_body.listOpenAPIComponentRequestBodies", "parse_error", err)
		return nil, err
	}

	// Return nil, if no request bodies object defined
	if doc.Components == nil || doc.Components.RequestBodies == nil {
		return nil, nil
	}

	// For each request body, scan its arguments
	for k, v := range doc.Components.RequestBodies {
		requestBodyObject := openAPIComponentRequestBody{
			Path: path,
			Key:  k,
		}

		for header, content := range v.Value.Content {
			requestBodyObject.Content = append(requestBodyObject.Content, map[string]interface{}{
				"contentType": header,
				"examples":    content.Examples,
				"schema":      content.Schema,
				"schemaType":  content.Schema.Value.Type,
				"encoding":    content.Encoding,
			})
		}
		requestBodyObject.Raw = *v.Value
		d.StreamListItem(ctx, requestBodyObject)

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
