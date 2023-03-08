package openapi

import (
	"context"
	p "path"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/getkin/kin-openapi/openapi3"
)

//// TABLE DEFINITION

func tableOpenAPIPath(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openapi_path",
		Description: "Paths specified by OpenAPI/Swagger standard version 3",
		List: &plugin.ListConfig{
			ParentHydrate: listFiles,
			Hydrate:       listOpenAPIPaths,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: []*plugin.Column{
			{Name: "api_path", Description: "A relative path to an individual endpoint.", Type: proto.ColumnType_STRING},
			{Name: "method", Description: "Specify the HTTP method.", Type: proto.ColumnType_STRING},
			{Name: "description", Description: "A verbose explanation of the operation behavior.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Operation.Description")},
			{Name: "deprecated", Description: "True, if the operation to be deprecated.", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Operation.Deprecated")},
			{Name: "summary", Description: "A short summary of what the operation does.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Operation.Summary")},
			{Name: "operation_id", Description: "Unique string used to identify the operation.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Operation.OperationID")},

			// JSON fields
			{Name: "parameters", Description: "A list of parameters that are applicable for this operation.", Type: proto.ColumnType_JSON, Transform: transform.FromField("Operation.Parameters")},
			{Name: "request_body", Description: "The request body applicable for this operation.", Type: proto.ColumnType_JSON, Transform: transform.FromField("Operation.RequestBody")},
			{Name: "responses", Description: "The list of possible responses as they are returned from executing this operation.", Type: proto.ColumnType_JSON, Transform: transform.FromField("Operation.Responses")},
			{Name: "callbacks", Description: "A map of possible out-of band callbacks related to the parent operation.", Type: proto.ColumnType_JSON, Transform: transform.FromField("Operation.Callbacks")},
			{Name: "security", Description: "A declaration of which security mechanisms can be used for this operation. The list of values includes alternative security requirement objects that can be used.", Type: proto.ColumnType_JSON, Transform: transform.FromField("Operation.Security")},
			{Name: "servers", Description: "An alternative server array to service this operation. If an alternative server object is specified at the Path Item Object or Root level, it will be overridden by this value.", Type: proto.ColumnType_JSON, Transform: transform.FromField("Operation.Servers")},
			{Name: "external_docs", Description: "Additional external documentation for this operation.", Type: proto.ColumnType_JSON, Transform: transform.FromField("Operation.ExternalDocs")},
			{Name: "tags", Description: "A list of tags for API documentation control.", Type: proto.ColumnType_JSON, Transform: transform.FromField("Operation.Tags")},
			{Name: "path", Description: "Path to the file.", Type: proto.ColumnType_STRING},
		},
	}
}

type openAPIPath struct {
	Path      string
	ApiPath   string
	Method    string
	Operation *openapi3.Operation
}

//// LIST FUNCTION

func listOpenAPIPaths(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
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

			d.StreamListItem(ctx, openAPIPath{
				Path:      path,
				ApiPath:   p.Join(apiPath, op),
				Method:    strings.ToUpper(op),
				Operation: operation,
			})
		}
	}

	return nil, nil
}

func getOperationInfoByType(operationType string, pathItem *openapi3.PathItem) *openapi3.Operation {
	switch operationType {
	case "connect":
		return pathItem.Connect
	case "delete":
		return pathItem.Delete
	case "get":
		return pathItem.Get
	case "head":
		return pathItem.Head
	case "options":
		return pathItem.Options
	case "patch":
		return pathItem.Patch
	case "post":
		return pathItem.Post
	case "put":
		return pathItem.Put
	case "trace":
		return pathItem.Trace
	}

	return nil
}
