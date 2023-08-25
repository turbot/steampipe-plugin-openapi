package openapi

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func openAPICommonColumns(columns []*plugin.Column) []*plugin.Column {
	allColumns := definitionResourceColumns()
	allColumns = append(allColumns, columns...)
	return allColumns
}

func definitionResourceColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "start_line", Type: proto.ColumnType_INT, Description: "The start line of the block.", Transform: transform.FromField("StartLine").NullIfZero()},
		{Name: "end_line", Type: proto.ColumnType_INT, Description: "The end line of the block.", Transform: transform.FromField("EndLine").NullIfZero()},
	}
}
