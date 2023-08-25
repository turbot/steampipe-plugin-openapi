# Table: openapi_component_schema

The table `openapi_component_schema` describes the reusable schemas that can be referenced throughout the API definition. The schemas allows users to define data models or schemas that describe the structure of the API data.

A schema in OpenAPI definition can describe the schema for request or response payload, error messages, headers, authentication etc.

## Examples

### Basic info

```sql
select
  name,
  type,
  format,
  deprecated,
  description,
  default_value,
  path,
  start_line,
  end_line
from
  openapi_component_schema;
```

### Get the properties returned by a specific API on success

```sql
with get_schema_ref as (
  select
    api_path,
    r.key as response_status,
    r.value -> 'content' -> 'application/json' -> 'schema' ->> '$ref' as schema_ref
  from
    openapi_path,
    jsonb_each(responses) as r
  where
    api_path = '/repos/{owner}/{repo}/issues/post'
    and r.key::integer >= '201' and r.key::integer < 300
)
select
  r.api_path,
  s.name,
  jsonb_pretty(s.properties) as schema_property,
  jsonb_pretty(s.required) as required,
  s.description
from
  get_schema_ref as r
  join openapi_component_schema as s on r.schema_ref = concat('#/components/schemas/', s.name);
```

### Get the schema of a required parameter of a specific API

```sql
select
  op.api_path,
  cp.required,
  jsonb_pretty(cp.schema) as schema
from
  openapi_path as op,
  jsonb_array_elements(op.parameters) as p
  join openapi_component_parameter as cp on (p ->> '$ref') = concat('#/components/parameters/', cp.name)
where
  op.api_path = '/orgs/{org}/members/{username}/delete'
  and cp.required;
```
