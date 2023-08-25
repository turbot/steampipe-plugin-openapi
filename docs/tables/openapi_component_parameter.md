# Table: openapi_component_parameter

The table `openapi_component_parameter` describes the input parameters that are used in API requests. These parameters can be used to specify the values that should be sent as part of the request, such as query parameters, headers, or request body data.

Parameters object in OpenAPI definition can be used to define path parameters, query parameters, header parameters, request body etc.

## Examples

### Basic info

```sql
select
  name,
  description,
  location,
  deprecated,
  required,
  schema,
  path,
  start_line,
  end_line
from
  openapi_component_parameter;
```

### List all supported parameters for a specific API endpoint

```sql
with list_parameters as (
  select
    api_path,
    p ->> '$ref' as parameter_ref
  from
    openapi_path,
    jsonb_array_elements(parameters) as p
  where
    api_path = '/repos/{owner}/{repo}/issues/post'
)
select
  l.api_path,
  p.name as parameter_name,
  p.required as is_required,
  p.deprecated as is_deprecated,
  jsonb_pretty(p.schema) as parameter_schema
from
  list_parameters as l
  join openapi_component_parameter as p on l.parameter_ref = concat('#/components/parameters/', p.name);
```

### List parameters with no schema

```sql
select
  name,
  description,
  location,
  deprecated,
  required,
  schema,
  path
from
  openapi_component_parameter
where
  schema is null
  and schema_ref is null;
```

### List deprecated parameters with no alternative mentioned in the description

```sql
select
  name,
  description,
  location,
  deprecated,
  required,
  schema,
  path
from
  openapi_component_parameter
where
  deprecated
  and description is null;
```
