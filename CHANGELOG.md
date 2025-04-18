## v1.1.1 [2025-04-18]

_Bug fixes_

- Fixed Linux AMD64 plugin build failures for `Postgres 14 FDW`, `Postgres 15 FDW`, and `SQLite Extension` by upgrading GitHub Actions runners from `ubuntu-20.04` to `ubuntu-22.04`.

## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#71](https://github.com/turbot/steampipe-plugin-openapi/pull/71))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#71](https://github.com/turbot/steampipe-plugin-openapi/pull/71))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#51](https://github.com/turbot/steampipe-plugin-openapi/pull/51))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#51](https://github.com/turbot/steampipe-plugin-openapi/pull/51))

## v0.2.1 [2023-12-12]

_Bug fixes_

- Fixed the missing optional tag on the `Paths` connection config parameter.

## v0.2.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#31](https://github.com/turbot/steampipe-plugin-openapi/pull/31))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#31](https://github.com/turbot/steampipe-plugin-openapi/pull/31))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-openapi/blob/main/docs/LICENSE). ([#31](https://github.com/turbot/steampipe-plugin-openapi/pull/31))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#30](https://github.com/turbot/steampipe-plugin-openapi/pull/30))

## v0.1.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#18](https://github.com/turbot/steampipe-plugin-openapi/pull/18))

## v0.1.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#15](https://github.com/turbot/steampipe-plugin-openapi/pull/15))
- Recompiled plugin with Go version `1.21`. ([#15](https://github.com/turbot/steampipe-plugin-openapi/pull/15))

## v0.0.2 [2023-04-03]

_Bug fixes_

- Fixed the default value of the `paths` argument in the `config/openapi.spc` file to also include `.yml` and `.yaml` files in the current working directory. ([#4](https://github.com/turbot/steampipe-plugin-openapi/pull/4))
- Fixed the brand color.

## v0.0.1 [2023-03-31]

_What's new?_

- New tables added
  - [openapi_component_header](https://hub.steampipe.io/plugins/turbot/openapi/tables/openapi_component_header)
  - [openapi_component_parameter](https://hub.steampipe.io/plugins/turbot/openapi/tables/openapi_component_parameter)
  - [openapi_component_request_body](https://hub.steampipe.io/plugins/turbot/openapi/tables/openapi_component_request_body)
  - [openapi_component_response](https://hub.steampipe.io/plugins/turbot/openapi/tables/openapi_component_response)
  - [openapi_component_schema](https://hub.steampipe.io/plugins/turbot/openapi/tables/openapi_component_schema)
  - [openapi_component_security_scheme](https://hub.steampipe.io/plugins/turbot/openapi/tables/openapi_component_security_scheme)
  - [openapi_info](https://hub.steampipe.io/plugins/turbot/openapi/tables/openapi_info)
  - [openapi_path](https://hub.steampipe.io/plugins/turbot/openapi/tables/openapi_path)
  - [openapi_path_request_body](https://hub.steampipe.io/plugins/turbot/openapi/tables/openapi_path_request_body)
  - [openapi_path_response](https://hub.steampipe.io/plugins/turbot/openapi/tables/openapi_path_response)
  - [openapi_server](https://hub.steampipe.io/plugins/turbot/openapi/tables/openapi_server)
