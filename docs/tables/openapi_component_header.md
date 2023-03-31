# Table: openapi_component_header

The table `openapi_component_header` describes the header definitions defined in the component object of an OpenAPI specification file.

The components header object is used to define reusable header objects that can be used across multiple API operations.

## Examples

### Basic info

```sql
select
  key,
  name,
  location,
  description,
  path
from
  openapi_component_header;
```

### List unused header definitions

```sql
with list_available_headers as (
  select
    path,
    array_agg(key) as header_refs
  from
    openapi_component_header
  group by
    path
),
list_used_headers as (
  select
    path,
    array_agg(distinct split_part(value ->> '$ref', '/', '4')) as headers
  from
    openapi_path_response,
    jsonb_each(headers)
  where
    (value ->> '$ref') is not null
  group by
    path
),
unused_headers_definitions as (
  select path, unnest(header_refs) as data from list_available_headers
    except
  select path, unnest(headers) as data from list_used_headers
)
select
  path,
  concat('components.headers.', data) as header_ref
from
  unused_headers_definitions;
```

### List headers with no schema

```sql
select
  key,
  location,
  deprecated,
  description,
  path
from
  openapi_component_header
where
  schema is null
  and schema_ref is null;
```

### List deprecated headers with no alternative mentioned in the description

```sql
select
  key,
  location,
  deprecated,
  description,
  path
from
  openapi_component_header
where
  deprecated
  and description is null;
```
