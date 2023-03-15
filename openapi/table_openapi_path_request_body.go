package openapi

import (
	"context"
	p "path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableOpenAPIPathRequestBody(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openapi_path_request_body",
		Description: "Path request body object.",
		List: &plugin.ListConfig{
			ParentHydrate: listOpenAPIFiles,
			Hydrate:       listOpenAPIPathRequestBodies,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: []*plugin.Column{
			{Name: "api_path", Description: "The key of the request body object definition.", Type: proto.ColumnType_STRING},
			{Name: "api_method", Description: "Specifies the HTTP method.", Type: proto.ColumnType_STRING},
			{Name: "description", Description: "A description of the request body.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Raw.Description")},
			{Name: "required", Description: "If true, the request body is required.", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Raw.Required")},
			{Name: "request_body_ref", Description: "The reference to the components request body object.", Type: proto.ColumnType_STRING},
			{Name: "content", Description: "A map containing descriptions of potential request body payloads.", Type: proto.ColumnType_JSON},
			{Name: "path", Description: "Path to the file.", Type: proto.ColumnType_STRING},
		},
	}
}

type openAPIPathRequestBody struct {
	Path           string
	ApiPath        string
	ApiMethod      string
	RequestBodyRef string
	Content        []map[string]interface{}
	Raw            openapi3.RequestBody
}

//// LIST FUNCTION

func listOpenAPIPathRequestBodies(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	// Get the parsed contents
	doc, err := getDoc(ctx, d, path)
	if err != nil {
		plugin.Logger(ctx).Error("openapi_path_request_body.listOpenAPIPathRequestBodies", "parse_error", err)
		return nil, err
	}

	// For each path, scan its request body object arguments
	for apiPath, item := range doc.Paths {
		for _, op := range OperationTypes {
			operation := getOperationInfoByType(op, item)

			// Skip if no method defined
			if operation == nil {
				continue
			}

			// Skip if no request body defined
			if operation.RequestBody == nil {
				continue
			}

			requestBodyObject := openAPIPathRequestBody{
				Path:           path,
				ApiPath:        p.Join(apiPath, op),
				ApiMethod:      strings.ToUpper(op),
				RequestBodyRef: operation.RequestBody.Ref,
			}

			for header, content := range operation.RequestBody.Value.Content {
				var schema interface{}
				if content.Schema.Ref != "" {
					schema = content.Schema.Ref
				} else {
					schema = content.Schema
				}
				requestBodyObject.Content = append(requestBodyObject.Content, map[string]interface{}{
					"contentType": header,
					"examples":    content.Examples,
					"schema":      schema,
					"schemaType":  content.Schema.Value.Type,
					"encoding":    content.Encoding,
				})
				requestBodyObject.Raw = *operation.RequestBody.Value
			}
			d.StreamListItem(ctx, requestBodyObject)

			// Context may get cancelled due to manual cancellation or if the limit has been reached
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}
