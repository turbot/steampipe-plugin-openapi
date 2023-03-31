package openapi

import (
	"context"
	"errors"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	filehelpers "github.com/turbot/go-kit/files"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

var OperationTypes = []string{"connect", "delete", "get", "head", "options", "patch", "post", "put", "trace"}

type filePath struct {
	Path string
}

func listOpenAPIFiles(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
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

// getDoc returns the parsed contents of the specified file
func getDoc(ctx context.Context, d *plugin.QueryData, path string) (*openapi3.T, error) {
	// Create custom hydrate data to pass through the path. Hydrate data
	// is normally per-column, but we can hijack it for this case to pass
	// through the context we need.
	h := &plugin.HydrateData{Item: path}
	i, err := getDocCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	return i.(*openapi3.T), nil
}

// Cached form of getDoc, using the per-connection and parallel safe
// Memoize() method.
var getDocCached = plugin.HydrateFunc(getDocUncached).Memoize(memoize.WithCacheKeyFunction(getDocCacheKey))

// getDoc is per-path, but Memoize() is per-connection, so a setup
// a custom cache key with path information in it.
func getDocCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Extract the path from the hydrate data. This is not per-row data,
	// but a clever pass through of context for our case.
	path := h.Item.(string)
	key := fmt.Sprintf("getDoc-%s", path)
	return key, nil
}

// getDocUncached is the actual implementation of getDoc, which should
// be run only once per path per connection. Do not call this directly, use
// getDoc instead.
func getDocUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Extract the path from the hydrate data. This is not per-row data,
	// but a clever pass through of context for our case.
	path := h.Item.(string)

	doc, err := openapi3.NewLoader().LoadFromFile(path)
	if err != nil {
		plugin.Logger(ctx).Error("getDocUncached", "file_error", err, "path", path)
		return nil, fmt.Errorf("failed to load file %s: %v", path, err)
	}

	plugin.Logger(ctx).Debug("getDocUncached", "connection_name", d.Connection.Name, "path", path, "status", "done")

	return doc, nil
}
