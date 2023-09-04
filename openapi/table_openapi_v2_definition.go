package openapi

import (
	"context"
	"strings"

	"github.com/go-openapi/spec"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableOpenAPIV2Definition(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openapi_v2_definition",
		Description: "Path object specified in OpenAPI V2 specification file.",
		List: &plugin.ListConfig{
			ParentHydrate: listOpenAPIFiles,
			Hydrate:       listOpenAPIV2Definitions,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "Unique identifier for the schema.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "ref",
				Description: "Reference to another schema.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "schema",
				Description: "URL pointing to the schema.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "Detailed description of the schema.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "type",
				Description: "Data type(s) accepted by the schema.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "nullable",
				Description: "Indicates if the value can be null.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "format",
				Description: "Specifies the data format.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "title",
				Description: "Title for the schema.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "default",
				Description: "Default value for the property.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "maximum",
				Description: "Maximum numerical value allowed.",
				Type:        proto.ColumnType_DOUBLE,
			},
			{
				Name:        "exclusive_maximum",
				Description: "Indicates whether the maximum is exclusive.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "minimum",
				Description: "Minimum numerical value allowed.",
				Type:        proto.ColumnType_DOUBLE,
			},
			{
				Name:        "exclusive_minimum",
				Description: "Indicates whether the minimum is exclusive.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "max_length",
				Description: "Maximum length for a string.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "min_length",
				Description: "Minimum length for a string.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "pattern",
				Description: "Regular expression pattern for string validation.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "max_items",
				Description: "Maximum number of items for an array.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "min_items",
				Description: "Minimum number of items for an array.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "unique_items",
				Description: "Indicates if all items in an array must be unique.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "multiple_of",
				Description: "Indicates a number that the value should be a multiple of.",
				Type:        proto.ColumnType_DOUBLE,
			},
			{
				Name:        "enum",
				Description: "List of valid enumeration values.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "max_properties",
				Description: "Maximum number of properties an object can have.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "min_properties",
				Description: "Minimum number of properties an object must have.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "required",
				Description: "List of required properties for an object.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "items",
				Description: "Schema or list of schemas for array items.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "all_of",
				Description: "List of schemas the data must adhere to.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "one_of",
				Description: "List of schemas, one of which the data must adhere to.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "any_of",
				Description: "List of schemas, any of which the data can adhere to.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "not",
				Description: "Schema that the data must not adhere to.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "properties",
				Description: "Properties of the schema.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "additional_properties",
				Description: "Schema for any additional properties.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "pattern_properties",
				Description: "Properties defined by a pattern.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "dependencies",
				Description: "Schema dependencies.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "additional_items",
				Description: "Schema for any additional items in an array.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "definitions",
				Description: "Definitions of sub-schemas.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "discriminator",
				Description: "Property name used to discriminate between different schema types.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "read_only",
				Description: "Indicates if the property is read-only.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "xml",
				Description: "XML representation details of the property.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("XML"),
			},
			{
				Name:        "external_docs",
				Description: "URL for external documentation related to the property.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "example",
				Description: "Example value for the property.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "extensions",
				Description: "Custom extensions to the Swagger Schema.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "extra_props",
				Description: "Additional arbitrary properties.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ExtraProps"),
			},
			{
				Name:        "path",
				Description: "Path to the file.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "key",
				Description: "The key used to refer or search the definition.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

type openAPIV2Definition struct {
	Path string
	Key  string
	spec.Schema
}

//// LIST FUNCTION

func listOpenAPIV2Definitions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	// Get the parsed contents
	doc, err := getV2Doc(ctx, d, path)
	if err != nil {
		plugin.Logger(ctx).Error("openapi_v2_definition.listOpenAPIV2Definitions", "parse_error", err)
		return nil, err
	}

	// check if the doc is swagger 2.0
	swagger := doc.Swagger
	if !strings.Contains(swagger, "2.0") {
		return nil, nil
	}

	// For each path, scan its arguments
	for key, item := range doc.Definitions {
		d.StreamListItem(ctx, openAPIV2Definition{path, key, item})

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
