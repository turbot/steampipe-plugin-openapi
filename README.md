![image](https://hub.steampipe.io/images/plugins/turbot/openapi-social-graphic.png)

# OpenAPI Plugin for Steampipe

Use SQL to query introspection of the OpenAPI definition.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/openapi)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/openapi/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-openapi/issues)

## Quick start

### Install

Download and install the latest OpenAPI plugin:

```bash
steampipe plugin install openapi
```

Configure your [config file](https://hub.steampipe.io/plugins/turbot/openapi#configuration) to include directories with OpenAPI definition files.

```hcl
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
```

Run steampipe:

```shell
steampipe query
```

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-openapi.git
cd steampipe-plugin-openapi
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```sh
make
```

Configure the plugin:

```shell
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/openapi.spc
```

Try it!

```shell
steampipe query
> .inspect openapi
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-openapi/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [OpenAPI Plugin](https://github.com/turbot/steampipe-plugin-openapi/labels/help%20wanted)
