# Table: openapi_component_request_body

The table `openapi_component_request_body` describes the request body definition defined in the component object of an OpenAPI specification file.

The `requestBody` object describes the payload of an HTTP request. It is used to specify the data format and structure of the request body for a particular operation.

## Examples

### Basic info

```sql
select
  key,
  description,
  required,
  jsonb_pretty(content)
from
  openapi_component_request_body;
```

### List unused request body definitions

```sql
with list_used_request_bodies as (
  select
    path,
    array_agg(distinct split_part(request_body_ref, '/', 4)) as req_bodies
  from
    openapi_path_request_body
  where
    request_body_ref is not null
  group by
    path
),
-- List all available request body definitions
all_request_body_definition as (
  select
    path,
    array_agg(key) as req_body_defs
  from
    openapi_component_request_body
  group by
    path
),
-- List all unused request body definitons
unused_request_body_definitions as (
  select path, unnest(req_body_defs) as data from all_request_body_definition
    except
  select path, unnest(req_bodies) as data from list_used_request_bodies
)
select
  path,
  concat('components.requestBodies.', data) as request_body_ref
from
  unused_request_body_definitions;
```

### List request body definitions without schema

```sql
select
  path,
  concat('components.requestBodies.', key, '.content.', c ->> 'contentType') as request_body_ref
from
  openapi_component_request_body,
  jsonb_array_elements(content) as c
where
  c ->> 'schema' is null;
```
