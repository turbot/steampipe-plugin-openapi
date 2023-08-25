# Table: openapi_path

The table `openapi_path` describes the individual endpoints or operations that make up the API. It is a key part of the definition because it defines the specific resources and actions that the API supports.

The `paths` object in the OpenAPI definition is a mapping of endpoint paths (relative to the base URL of the API) to endpoint objects. Each endpoint object describes a specific HTTP method (e.g., GET, POST, PUT, DELETE) and the parameters, request body, and responses associated with that method.

## Examples

### Basic info

```sql
select
  api_path,
  deprecated,
  description,
  tags,
  path,
  start_line,
  end_line
from
  openapi_path;
```

### List all deprecated endpoints

```sql
select
  api_path,
  description,
  tags,
  path
from
  openapi_path
where
  deprecated;
```

### List all GET method endpoints

```sql
select
  api_path,
  description,
  parameters,
  path
from
  openapi_path
where
  method = 'GET';
```

### Get the parameters required for a specific endpoint

```sql
select
  api_path,
  p ->> 'name' as param,
  case
    when p -> 'required' is not null then true
    else false
  end as is_required,
  jsonb_pretty(p -> 'schema') as schema,
  path
from
  openapi_path,
  jsonb_array_elements(parameters) as p
where
  api_path = '/org/{org_handle}/audit_log/get';
```

### Get the success response schema of a specific endpoint

```sql
select
  api_path,
  jsonb_pretty(responses -> '200') as success_response,
  path
from
  openapi_path
where
  api_path = '/identity/{identity_handle}/get';
```
