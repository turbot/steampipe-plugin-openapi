package openapi

import (
	"context"
	"os"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/getkin/kin-openapi/openapi3"
)

//// TABLE DEFINITION

func tableOpenAPIComponentSchema(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openapi_component_schema",
		Description: "Components schema object.",
		List: &plugin.ListConfig{
			ParentHydrate: listOpenAPIFiles,
			Hydrate:       listOpenAPIComponentSchemas,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: openAPICommonColumns([]*plugin.Column{
			{Name: "name", Description: "The name of the property.", Type: proto.ColumnType_STRING},
			{Name: "type", Description: "The type of the schema.", Type: proto.ColumnType_STRING},
			{Name: "format", Description: "The format of a specific schema type.", Type: proto.ColumnType_STRING},
			{Name: "deprecated", Description: "True, if the schema is deprecated.", Type: proto.ColumnType_BOOL},
			{Name: "title", Description: "The title of the schema.", Type: proto.ColumnType_STRING},
			{Name: "description", Description: "A description of the schema.", Type: proto.ColumnType_STRING},
			{Name: "default_value", Description: "The default value set for the schema.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Default")},

			// boolean fields
			{Name: "unique_items", Description: "True, if the items in an array are required to be unique.", Type: proto.ColumnType_BOOL},
			{Name: "exclusive_min", Description: "Specify the minimum value allowed for a numeric property, where the minimum value is exclusive.", Type: proto.ColumnType_BOOL},
			{Name: "exclusive_max", Description: "Specify the maximum value allowed for a numeric property, where the maximum value is exclusive", Type: proto.ColumnType_BOOL},
			{Name: "nullable", Description: "If true, null value can be set to a string property.", Type: proto.ColumnType_BOOL},
			{Name: "read_only", Description: "If true, the property value cannot be modified.", Type: proto.ColumnType_BOOL},
			{Name: "write_only", Description: "If true, the property value can be modified.", Type: proto.ColumnType_BOOL},
			{Name: "allow_empty_value", Description: "If true, it allows to set a empty value to the property.", Type: proto.ColumnType_BOOL},

			// info for number type data
			{Name: "min", Description: "Specify the minimum number that can be set to the property.", Type: proto.ColumnType_DOUBLE},
			{Name: "max", Description: "Specify the maximum number that can be set to the property.", Type: proto.ColumnType_DOUBLE},
			{Name: "multiple_of", Description: "Specify a numeric property's valid multiplier.", Type: proto.ColumnType_DOUBLE},

			// info for string type data
			{Name: "min_length", Description: "Specify the minimum length of a string type data.", Type: proto.ColumnType_DOUBLE},
			{Name: "max_length", Description: "Specify the maximum length of a string type data.", Type: proto.ColumnType_DOUBLE},
			{Name: "pattern", Description: "Specify the regex pattern for the property value.", Type: proto.ColumnType_STRING},

			// info for array type data
			{Name: "min_items", Description: "Specify the minimum number of items that can be added to the property of type array.", Type: proto.ColumnType_DOUBLE},
			{Name: "max_items", Description: "Specify the maximum number of items that can be added to the property of type array.", Type: proto.ColumnType_DOUBLE},
			{Name: "items", Description: "Specify the list of items defined in the property.", Type: proto.ColumnType_JSON},

			{Name: "required", Description: "If true, the property must be defined.", Type: proto.ColumnType_JSON},
			{Name: "properties", Description: "Describes the schema properties.", Type: proto.ColumnType_JSON},
			{Name: "path", Description: "Path to the file.", Type: proto.ColumnType_STRING},
		}),
	}
}

type openAPIComponentSchema struct {
	Path      string
	Name      string
	StartLine int
	EndLine   int
	openapi3.Schema
	Properties map[string]interface{}
}

//// LIST FUNCTION

func listOpenAPIComponentSchemas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	file, err := os.Open(path)
	if err != nil {
		plugin.Logger(ctx).Error("openapi_component_schema.listOpenAPIComponentSchemas", "file_open_error", err)
		return nil, err
	}

	// Get the parsed contents
	doc, err := getDoc(ctx, d, path)
	if err != nil {
		plugin.Logger(ctx).Error("openapi_component_schema.listOpenAPIComponentSchemas", "parse_error", err)
		return nil, err
	}

	// Return nil, if no schemas object defined
	if doc.Components == nil || doc.Components.Schemas == nil {
		return nil, nil
	}

	// For each schema, scan its arguments
	for k, v := range doc.Components.Schemas {

		// fetch start and end line for each schemas
		var startLine, endLine int
		if strings.HasSuffix(path, "json") {
			startLine, endLine = findBlockLinesFromJSON(file, "components", k)
		} else {
			startLine, endLine = findBlockLinesFromYML(file, "components", k)
		}

		properties := map[string]interface{}{}
		for i, j := range v.Value.Properties {
			properties[i] = j.Value
		}
		d.StreamListItem(ctx, openAPIComponentSchema{path, k, startLine, endLine, *v.Value, properties})

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
