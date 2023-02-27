package openapi

import (
	"context"
	"errors"

	filehelpers "github.com/turbot/go-kit/files"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type filePath struct {
	Path string
}

func listFiles(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// #1 - Path via qual

	// If the path was requested through qualifier then match it exactly. Globs
	// are not supported in this context since the output value for the column
	// will never match the requested value.
	quals := d.EqualsQuals
	if quals["path"] != nil {
		d.StreamListItem(ctx, filePath{Path: quals["path"].GetStringValue()})
		return nil, nil
	}

	// #2 - paths in config

	// Glob paths in config
	// Fail if no paths are specified
	openAPIConfig := GetConfig(d.Connection)
	if openAPIConfig.Paths == nil {
		return nil, errors.New("paths must be configured")
	}

	// Gather file path matches for the glob
	var matches []string
	paths := openAPIConfig.Paths
	for _, i := range paths {

		// List the files in the given source directory
		files, err := d.GetSourceFiles(i)
		if err != nil {
			return nil, err
		}
		matches = append(matches, files...)
	}

	// Sanitize the matches to ignore the directories
	for _, i := range matches {

		// Ignore directories
		if filehelpers.DirectoryExists(i) {
			continue
		}
		d.StreamListItem(ctx, filePath{Path: i})
	}

	return nil, nil
}
