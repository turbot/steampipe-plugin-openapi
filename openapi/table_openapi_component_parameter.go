package openapi

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/getkin/kin-openapi/openapi3"
)

//// TABLE DEFINITION

func tableOpenAPIComponentParameter(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openapi_component_parameter",
		Description: "The input parameters that are used in API requests",
		List: &plugin.ListConfig{
			ParentHydrate: listFiles,
			Hydrate:       listOpenAPIComponentParameters,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: []*plugin.Column{
			{Name: "name", Description: "The name of the parameter.", Type: proto.ColumnType_STRING},
			{Name: "location", Description: "The location of the parameter. Possible values are query, header, path or cookie.", Type: proto.ColumnType_STRING, Transform: transform.FromField("In")},
			{Name: "description", Description: "A brief description of the parameter.", Type: proto.ColumnType_STRING},
			{Name: "style", Description: "Describes how the parameter value will be serialized depending on the type of the parameter value. Default values (based on value of in): for query - form; for path - simple; for header - simple; for cookie - form.", Type: proto.ColumnType_STRING},
			{Name: "deprecated", Description: "True, if the parameter is deprecated.", Type: proto.ColumnType_BOOL},
			{Name: "explode", Description: "If true, parameter values of type array or object generate separate parameters for each value of the array or key-value pair of the map.", Type: proto.ColumnType_BOOL},
			{Name: "allow_empty_value", Description: "", Type: proto.ColumnType_BOOL},
			{Name: "allow_reserved", Description: "Determines whether the parameter value SHOULD allow reserved characters, as defined by RFC3986 (e.g. :/?#[]@!$&'()*+,;=) to be included without percent-encoding. This property only applies to parameters with an in value of query. The default value is false.", Type: proto.ColumnType_BOOL},
			{Name: "required", Description: "True, if the parameter is required.", Type: proto.ColumnType_BOOL},
			{Name: "schema", Description: "The schema of the parameter.", Type: proto.ColumnType_JSON, Transform: transform.FromField("Schema.Value")},
			{Name: "schema_ref", Description: "The schema reference of the parameter.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Schema.Ref")},
			{Name: "key", Description: "The key used to refer or search the parameter.", Type: proto.ColumnType_STRING},
			{Name: "path", Description: "Path to the file.", Type: proto.ColumnType_STRING},
		},
	}
}

type openAPIComponentParameter struct {
	Path string
	Key  string
	openapi3.Parameter
}

//// LIST FUNCTION

func listOpenAPIComponentParameters(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	doc, err := openapi3.NewLoader().LoadFromFile(path)
	if err != nil {
		plugin.Logger(ctx).Error("openapi_component_parameter.listOpenAPIComponentParameters", "file_error", err, "path", path)
		return nil, fmt.Errorf("failed to load file %s: %v", path, err)
	}

	// Return nil, if no parameters defined
	if doc.Components == nil || doc.Components.Parameters == nil {
		return nil, nil
	}

	for k, v := range doc.Components.Parameters {
		d.StreamListItem(ctx, openAPIComponentParameter{path, k, *v.Value})
	}

	return nil, nil
}
