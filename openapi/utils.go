package openapi

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

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

// findBlockLinesFromJSON locates the start and end lines of a specific block or nested element within a block.
// The file should contain structured data (e.g., JSON) and this function expects to search for blocks with specific names.
func findBlockLinesFromJSON(file *os.File, blockName string, pathName ...string) (int, int) {
	var currentLine, startLine, endLine int
	var bracketCounter int

	// These boolean flags indicate which part of the structured data we're currently processing.
	inBlock, inPath, inResponseStatus, inRequestBody, inServer, inComponent := false, false, false, false, false, false

	// Move the file pointer to the start of the file.
	_, _ = file.Seek(0, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currentLine++
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)

		// Detect the start of the desired block, path, response, etc.
		// Depending on the blockName and provided pathName, different conditions are checked.

		// Generic block detection
		if !inBlock && ((trimmedLine == fmt.Sprintf(`"%s": {`, blockName) || trimmedLine == fmt.Sprintf(`"%s": [`, blockName)) || trimmedLine == fmt.Sprintf(`%s:`, blockName)) {
			inBlock = true
			bracketCounter = 1
			startLine = currentLine
			continue
		} else if inBlock && blockName == "components" && trimmedLine == fmt.Sprintf(`"%s": {`, pathName[0]) {
			// Different component block detection within the "components" block
			inComponent = true
			bracketCounter = 1
			startLine = currentLine
			continue
		} else if inBlock && blockName == "servers" && strings.Contains(trimmedLine, fmt.Sprintf(`"url": "%s"`, pathName[0])) {
			// Server detection within the "servers" block
			inServer = true
			bracketCounter = 1
			startLine = currentLine - 1
			continue
		} else if inBlock && blockName == "paths" && len(pathName) > 0 && trimmedLine == fmt.Sprintf(`"%s": {`, pathName[0]) {
			// Path detection within the "paths" block
			inPath = true
			bracketCounter = 1
			startLine = currentLine
			continue
		} else if inPath && len(pathName) > 1 && pathName[1] == "requestBody" && trimmedLine == `"requestBody": {` {
			// Request body detection within a path
			inRequestBody = true
			bracketCounter = 1
			startLine = currentLine
			continue
		} else if inPath && len(pathName) > 1 && trimmedLine == fmt.Sprintf(`"%s": {`, pathName[1]) {
			// Response status detection within a path
			inResponseStatus = true
			bracketCounter = 1
			startLine = currentLine
			continue
		}

		// If we are within a block, we need to track the opening and closing brackets
		// to determine where the block ends.
		if (inBlock && !inServer) || (inBlock && !inComponent) || (inBlock && !inPath) || (inPath && !inResponseStatus) || (inPath && !inRequestBody) {
			bracketCounter += strings.Count(line, "{")
			bracketCounter -= strings.Count(line, "}")

			if bracketCounter == 0 {
				endLine = currentLine
				break
			}
		}
	}

	if startLine != 0 && endLine == 0 {
		// If we found the start but not the end, reset the start to indicate the block doesn't exist in entirety.
		startLine = 0
	}

	return startLine, endLine
}

// findBlockLinesFromYML locates the start and end lines of a specific block or nested element within a block.
// The file should contain structured data (e.g., YML/YAML) and this function expects to search for blocks with specific names.
func findBlockLinesFromYML(file *os.File, blockName string, pathName ...string) (int, int) {
	var currentLine, startLine, endLine, currentIndentLevel int
	var blockIndentLevel, pathIndentLevel, serverIndentLevel, requestBodyIndentLevel, componentIndentLevel, responseIndentLevel = -1, -1, -1, -1, -1, -1

	inBlock, inPath, inServer, inRequestBody, inComponent, inResponseStatus := false, false, false, false, false, false

	// Move the file pointer to the start of the file.
	_, _ = file.Seek(0, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currentLine++
		line := scanner.Text()

		// Determine the current indentation level by counting leading spaces.
		currentIndentLevel = len(line) - len(strings.TrimSpace(line))

		// Detect the start of the desired block.
		if !inBlock && strings.HasPrefix(strings.TrimSpace(line), blockName+":") {
			inBlock = true
			startLine = currentLine
			blockIndentLevel = currentIndentLevel
			continue
		} else if inBlock && blockName == "components" && len(pathName) > 0 && strings.HasPrefix(strings.TrimSpace(line), fmt.Sprintf("%s:", pathName[0])) {
			inComponent = true
			startLine = currentLine
			componentIndentLevel = currentIndentLevel
			continue
		} else if inBlock && blockName == "paths" && len(pathName) > 0 && strings.HasPrefix(strings.TrimSpace(line), fmt.Sprintf("%s:", pathName[0])) {
			inPath = true
			startLine = currentLine
			pathIndentLevel = currentIndentLevel
			continue
		} else if inPath && len(pathName) > 1 && pathName[1] == "requestBody" && strings.HasPrefix(strings.TrimSpace(line), "requestBody:") {
			inRequestBody = true
			startLine = currentLine
			requestBodyIndentLevel = currentIndentLevel
			continue
		} else if inPath && len(pathName) > 1 && strings.HasPrefix(strings.TrimSpace(line), fmt.Sprintf(`"%s":`, pathName[1])) {
			inResponseStatus = true
			startLine = currentLine
			responseIndentLevel = currentIndentLevel
			continue
		} else if inBlock && blockName == "servers" && len(pathName) > 0 && strings.Contains(strings.TrimSpace(line), fmt.Sprintf(`url: "%s"`, pathName[0])) {
			inServer = true
			startLine = currentLine
			serverIndentLevel = currentIndentLevel
			continue
		}

		// // If we are within a block, we need to track the closing
		if inComponent && currentIndentLevel <= componentIndentLevel {
			endLine = currentLine - 1
			break
		} else if inPath && currentIndentLevel <= pathIndentLevel {
			endLine = currentLine - 1
			break
		} else if inServer && currentIndentLevel <= serverIndentLevel {
			endLine = currentLine - 1
			break
		} else if inRequestBody && currentIndentLevel <= requestBodyIndentLevel {
			endLine = currentLine - 1
			break
		} else if inResponseStatus && currentIndentLevel <= responseIndentLevel {
			endLine = currentLine - 1
			break
		} else if inBlock && currentIndentLevel <= blockIndentLevel {
			endLine = currentLine - 1
			break
		}
	}

	if startLine != 0 && endLine == 0 {
		endLine = currentLine // Consider the end of the file as the end of the block or path
	}

	return startLine, endLine
}
