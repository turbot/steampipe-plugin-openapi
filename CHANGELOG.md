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
