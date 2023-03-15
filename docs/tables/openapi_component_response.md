# Table: openapi_component_response

The table `openapi_component_response` describes the response definitions defined in the component object of an OpenAPI specification file.

The components response object provides a way to define reusable response objects that can be referenced in multiple operations. It allows users to define common responses that can be used across different API endpoints, reducing duplication and making it easier to maintain and update the API specification.

## Examples

### Basic info

```sql
select
  key,
  description,
  jsonb_pretty(content),
  path
from
  openapi_component_response;
```

### List unused response definitions

```sql
with list_used_response_defs as (
  select
    path,
    array_agg(distinct split_part(response_ref, '/', '4')) as resp
  from
    openapi_path_response
  where
    response_ref is not null
  group by
    path
),
all_responses_definition as (
  select
    path,
    array_agg(key) as resp_defs
  from
    openapi_component_response
  group by
    path
),
unused_response_definitions as (
  select path, unnest(resp_defs) as data from all_responses_definition
    except
  select path, unnest(resp) as data from list_used_response_defs
)
select
  path,
  concat('components.responses.', data) as response_ref
from
  unused_response_definitions;
```

### List response definitions without schema

```sql
select
  path,
  concat('components.responses.', key, '.content.', c ->> 'contentType') as paths
from
  openapi_component_response,
  jsonb_array_elements(content) as c
where
  c ->> 'schema' is null;
```
