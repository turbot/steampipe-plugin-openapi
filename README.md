![image](https://hub.steampipe.io/images/plugins/turbot/openapi-social-graphic.png)

# OpenAPI Plugin for Steampipe

Use SQL to query introspection of the OpenAPI definition.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/openapi)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/openapi/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-openapi/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install openapi
```

Configure your [config file](https://hub.steampipe.io/plugins/turbot/openapi#configuration) to include directories with OpenAPI definition files. If no directory is specified, the current working directory will be used.

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
