# Table: openapi_path_response

The table `openapi_path_response` describes the response definition defined in the path object of an OpenAPI specification file.

The `response` object defined in a path object describes the possible responses for an API endpoint.

## Examples

### Basic info

```sql
select
  api_path,
  api_method,
  response_status,
  jsonb_pretty(content),
  path,
  start_line,
  end_line
from
  openapi_path_response;
```

### Get the response object for a specific API operation on success

```sql
select
  api_path,
  api_method,
  response_status,
  jsonb_pretty(content),
  path
from
  openapi_path_response
where
  api_path = '/app/installations/get'
  and response_status = '200';
```

### List response definitions without schema

```sql
select
  path,
  concat(api_path, '.responses.', response_status, '.content.', c ->> 'contentType') as paths
from
  openapi_path_response,
  jsonb_array_elements(content) as c
where
  c ->> 'schema' is null
  and response_ref is null;
```
