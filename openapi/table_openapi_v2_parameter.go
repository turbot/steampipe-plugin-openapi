package openapi

import (
	"context"

	"github.com/go-openapi/spec"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableOpenAPIV2Parameter(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openapi_v2_parameter",
		Description: "Path object specified in OpenAPI V2 specification file.",
		List: &plugin.ListConfig{
			ParentHydrate: listOpenAPIFiles,
			Hydrate:       listOpenAPIV2Parameters,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "maximum",
				Description: "The maximum numeric value allowed for the property.",
				Type:        proto.ColumnType_DOUBLE,
			},
			{
				Name:        "exclusive_maximum",
				Description: "Indicates if the maximum value is exclusive.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "minimum",
				Description: "The minimum numeric value allowed for the property.",
				Type:        proto.ColumnType_DOUBLE,
			},
			{
				Name:        "exclusive_minimum",
				Description: "Indicates if the minimum value is exclusive.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "max_length",
				Description: "Maximum length allowed for the property.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "min_length",
				Description: "Minimum length allowed for the property.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "pattern",
				Description: "Regular expression pattern that the property must match.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "max_items",
				Description: "Maximum number of items allowed for array type property.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "min_items",
				Description: "Minimum number of items required for array type property.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "unique_items",
				Description: "Indicates if the array type property must have unique items.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "multiple_of",
				Description: "Indicates that the property value must be a multiple of the given number.",
				Type:        proto.ColumnType_DOUBLE,
			},
			{
				Name:        "enum",
				Description: "List of valid values for the property.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "ref",
				Description: "A reference to an external definition that replaces this definition.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "type",
				Description: "Data type of the property.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "nullable",
				Description: "Indicates if the property can be null.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "format",
				Description: "Format of the property if applicable.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "items",
				Description: "Schema or array of schemas for array type properties.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "collection_format",
				Description: "Format of the array if type array is used.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "default",
				Description: "Default value for the property.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "example",
				Description: "Example value for the property.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "extensions",
				Description: "Custom extensions for the CommonValidations.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "description",
				Description: "Detailed description of the property.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "Name of the property.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "in",
				Description: "Location of the property (e.g., 'header', 'query').",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "required",
				Description: "Indicates if the property is required.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "schema",
				Description: "Schema defining the type used for the property.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "allow_empty_value",
				Description: "Indicates if an empty value is allowed for the property.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "path",
				Description: "Path to the file.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "key",
				Description: "The key used to refer or search the parameter.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

type openAPIV2Parameter struct {
	Path string
	Key  string
	spec.Parameter
}

//// LIST FUNCTION

func listOpenAPIV2Parameters(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
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

	// For each path, scan its arguments
	for key, item := range doc.Parameters {
		d.StreamListItem(ctx, openAPIV2Parameter{path, key, item})

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
