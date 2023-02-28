---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/openapi.svg"
brand_color: "#2483C0"
display_name: "OpenAPI"
short_name: "openapi"
description: "Steampipe plugin to query introspection of the OpenAPI definition."
og_description: "Query OpenAPI files with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/openapi-social-graphic.png"
---

# OpenAPI + Steampipe

A self-contained or composite resource which defines or describes an API or elements of an API.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query data using SQL.

Query all the endpoints available for an API:

```sql
select
  api_path,
  operation_id,
  summary,
  deprecated,
  tags
from
  openapi_path;
```

```sh
+----------+---------------------+--------------------------+------------+--------+
| api_path | operation_id        | summary                  | deprecated | tags   |
+----------+---------------------+--------------------------+------------+--------+
| /get     | listVersionsv2      | List API versions        | false      | <null> |
| /v2/get  | getVersionDetailsv2 | Show API version details | false      | <null> |
+----------+---------------------+--------------------------+------------+--------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/openapi/tables)**

## Get started

### Install

Download and install the latest OpenAPI plugin:

```bash
steampipe plugin install openapi
```

### Credentials

No credentials are required.

### Configuration

Installing the latest openAPI plugin will create a config file (`~/.steampipe/config/openapi.spc`) with a single connection named `openapi`:

```hcl
connection "openapi" {
  plugin = "openapi"

  # Paths is a list of locations to search for OpenAPI definition files
  # Paths can be configured with a local directory, a remote Git repository URL, or an S3 bucket URL
  # Wildcard based searches are supported, including recursive searches
  # Local paths are resolved relative to the current working directory (CWD)

  # For example:
  #  - "*.json" matches all OpenAPI definition files in the CWD
  #  - "**/*.json" matches all OpenAPI definition files in the CWD and all sub-directories
  #  - "../*.json" matches all OpenAPI definition files in the CWD's parent directory
  #  - "steampipe*.json" matches all OpenAPI definition files starting with "steampipe" in the CWD
  #  - "/path/to/dir/*.json" matches all OpenAPI definition files in a specific directory
  #  - "/path/to/dir/main.json" matches a specific file

  # If paths includes "*", all files (including non-OpenAPI definition files) in
  # the CWD will be matched, which may cause errors if incompatible file types exist

  # Defaults to CWD
  paths = [ "*.json" ]
}

```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-openapi
- Community: [Slack Channel](https://steampipe.io/community/join)
