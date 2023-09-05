package openapi

import (
	"context"
	p "path"
	"strings"

	"github.com/go-openapi/spec"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableOpenAPIV2Path(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openapi_v2_path",
		Description: "Path object specified in OpenAPI V2 specification file.",
		List: &plugin.ListConfig{
			ParentHydrate: listOpenAPIFiles,
			Hydrate:       listOpenAPIV2Paths,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "api_path",
				Description: "A relative path to an individual endpoint.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "method",
				Description: "Specify the HTTP method (e.g., GET, POST, PUT, DELETE).",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "A verbose explanation of the operation behavior.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Operation.Description"),
			},
			{
				Name:        "deprecated",
				Description: "Indicates whether the operation is deprecated.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Operation.Deprecated"),
			},
			{
				Name:        "summary",
				Description: "A short summary of what the operation does.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Operation.Summary"),
			},
			{
				Name:        "operation_id",
				Description: "Unique string used to identify the operation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Operation.ID"),
			},
			{
				Name:        "parameters",
				Description: "A JSON object containing a list of parameters applicable for this operation.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Operation.Parameters"),
			},
			{
				Name:        "consumes",
				Description: "A JSON object specifying the MIME types the operation can consume.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Operation.Consumes"),
			},
			{
				Name:        "produces",
				Description: "A JSON object specifying the MIME types the operation can produce.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Operation.Produces"),
			},
			{
				Name:        "responses",
				Description: "A JSON object specifying the possible responses returned by the operation.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Operation.Responses"),
			},
			{
				Name:        "schemes",
				Description: "A JSON object listing the supported schemes (e.g., HTTP, HTTPS).",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Operation.Schemes"),
			},
			{
				Name:        "security",
				Description: "A JSON object declaring the security mechanisms applicable to the operation.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Operation.Security"),
			},
			{
				Name:        "external_docs",
				Description: "A JSON object containing links to additional external documentation for the operation.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Operation.ExternalDocs"),
			},
			{
				Name:        "tags",
				Description: "A JSON object containing a list of tags used for API documentation control.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Operation.Tags"),
			},
			{
				Name:        "path",
				Description: "Path to the file containing the API operation definition or related data.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

type openAPIV2Path struct {
	Path      string
	ApiPath   string
	Method    string
	Operation *spec.Operation
}

//// LIST FUNCTION

func listOpenAPIV2Paths(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	// Get the parsed contents
	doc, err := getV2Doc(ctx, d, path)
	if err != nil {
		plugin.Logger(ctx).Error("openapi_v2_path.listOpenAPIV2Paths", "parse_error", err)
		return nil, err
	}

	// check if the doc is swagger 2.0
	if doc.Swagger != "2.0" {
		return nil, nil
	}

	// For each path, scan its arguments
	for apiPath, item := range doc.Paths.Paths {
		for _, op := range OperationTypes {
			operation := getV2OperationInfoByType(op, item)

			// Skip if no method defined
			if operation == nil {
				continue
			}

			d.StreamListItem(ctx, openAPIV2Path{
				Path:      path,
				ApiPath:   p.Join(apiPath, op),
				Method:    strings.ToUpper(op),
				Operation: operation,
			})

			// Context may get cancelled due to manual cancellation or if the limit has been reached
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getV2OperationInfoByType(operationType string, pathItem spec.PathItem) *spec.Operation {
	switch operationType {
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
	}

	return nil
}
