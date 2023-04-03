connection "openapi" {
  plugin = "openapi"

  # Paths is a list of locations to search for OpenAPI definition files
  # Paths can be configured with a local directory, a remote Git repository URL, or an S3 bucket URL
  # Refer https://hub.steampipe.io/plugins/turbot/openapi#supported-path-formats for more information
  # Wildcard based searches are supported, including recursive searches
  # Local paths are resolved relative to the current working directory (CWD)

  # For example:
  #  - "*.json" matches all OpenAPI JSON definition files in the CWD
  #  - "**/*.json" matches all OpenAPI JSON definition files in the CWD and all sub-directories
  #  - "../*.json" matches all OpenAPI JSON definition files in the CWD's parent directory
  #  - "*.yml" or "*.yaml" matches all OpenAPI YML or YAML definition files in the CWD
  #  - "**/*.yml" or "**/*.yaml" matches all OpenAPI YML or YAML definition files in the CWD and all sub-directories
  #  - "../*.yml" or "../*.yaml" matches all OpenAPI YML or YAML definition files in the CWD's parent directory
  #  - "steampipe*.json" or "steampipe*.yml" or "steampipe*.yaml" matches all OpenAPI definition files starting with "steampipe" in the CWD
  #  - "/path/to/dir/*.json" or "/path/to/dir/*.yml" or "/path/to/dir/*.yaml" matches all OpenAPI definition files in a specific directory
  #  - "/path/to/dir/main.json" or "/path/to/dir/main.yml" or "/path/to/dir/main.yaml" matches a specific file

  # If paths includes "*", all files (including non-OpenAPI definition files) in
  # the CWD will be matched, which may cause errors if incompatible file types exist

  # Defaults to CWD
  paths = [ "*.json", "*.yml", "*.yaml" ]
}
