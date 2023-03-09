package openapi

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/getkin/kin-openapi/openapi3"
)

//// TABLE DEFINITION

func tableOpenAPIComponentHeader(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openapi_component_header",
		Description: "The headers information defined in the components",
		List: &plugin.ListConfig{
			ParentHydrate: listFiles,
			Hydrate:       listOpenAPIComponentHeaders,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: []*plugin.Column{
			{Name: "key", Description: "The key used to refer or search the header.", Type: proto.ColumnType_STRING},
			{Name: "name", Description: "The name of the header.", Type: proto.ColumnType_STRING},
			{Name: "location", Description: "The location of the header. Possible values are query, header, path or cookie.", Type: proto.ColumnType_STRING, Transform: transform.FromField("In")},
			{Name: "description", Description: "A brief description of the header.", Type: proto.ColumnType_STRING},
			{Name: "style", Description: "Describes how the header value will be serialized depending on the type of the header value. Default values (based on value of in): for query - form; for path - simple; for header - simple; for cookie - form.", Type: proto.ColumnType_STRING},
			{Name: "deprecated", Description: "True, if the header is deprecated.", Type: proto.ColumnType_BOOL},
			{Name: "explode", Description: "If true, header values of type array or object generate separate headers for each value of the array or key-value pair of the map.", Type: proto.ColumnType_BOOL},
			{Name: "allow_empty_value", Description: "", Type: proto.ColumnType_BOOL},
			{Name: "allow_reserved", Description: "Determines whether the header value SHOULD allow reserved characters, as defined by RFC3986 (e.g. :/?#[]@!$&'()*+,;=) to be included without percent-encoding. This property only applies to headers with an in value of query. The default value is false.", Type: proto.ColumnType_BOOL},
			{Name: "required", Description: "True, if the header is required.", Type: proto.ColumnType_BOOL},
			{Name: "schema", Description: "The schema of the header.", Type: proto.ColumnType_JSON, Transform: transform.FromField("Schema.Value")},
			{Name: "schema_ref", Description: "The schema reference of the header.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Schema.Ref")},
			{Name: "path", Description: "Path to the file.", Type: proto.ColumnType_STRING},
		},
	}
}

type openAPIComponentHeader struct {
	Path string
	Key  string
	openapi3.Header
}

//// LIST FUNCTION

func listOpenAPIComponentHeaders(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	doc, err := getDoc(ctx, d, path)
	if err != nil {
		return nil, err
	}

	// Return nil, if no schema defined
	if doc.Components == nil || doc.Components.Headers == nil {
		return nil, nil
	}

	for k, v := range doc.Components.Headers {
		d.StreamListItem(ctx, openAPIComponentHeader{path, k, *v.Value})

		// Context may get cancelled due to manual cancellation or if the limit has been reached
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
