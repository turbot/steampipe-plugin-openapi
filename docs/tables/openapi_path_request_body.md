# Table: openapi_path_request_body

The table `openapi_path_request_body` describes the request body definition defined in the path object of an OpenAPI specification file.

The `requestBody` object defined in a path object describes the structure and content of a request body that's expected to be sent to the server for a specific API operation.

## Examples

### Basic info

```sql
select
  api_path,
  api_method,
  description,
  required,
  jsonb_pretty(content),
  path
from
  openapi_path_request_body;
```

### Get the request body object for a specific API operation

```sql
select
  api_path,
  api_method,
  description,
  required,
  jsonb_pretty(content),
  path
from
  openapi_path_request_body
where
  api_path = '/applications/{client_id}/token/post';
```

### List request body definitions without schema

```sql
select
  path,
  concat(api_path, '.requestBody.content.', c ->> 'contentType') as paths
from
  openapi_path_request_body,
  jsonb_array_elements(content) as c
where
  c ->> 'schema' is null
  and request_body_ref is null;
```
