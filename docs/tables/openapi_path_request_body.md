---
title: "Steampipe Table: openapi_path_request_body - Query OpenAPI Request Bodies using SQL"
description: "Allows users to query Request Bodies in OpenAPI, specifically details about each request body that is defined within an OpenAPI specification."
---

# Table: openapi_path_request_body - Query OpenAPI Request Bodies using SQL

OpenAPI is a specification for machine-readable interface files for describing, producing, consuming, and visualizing RESTful web services. Request Bodies in OpenAPI are used to send and receive data via the REST API. They provide detailed information about the type of data that an API can accept or return.

## Table Usage Guide

The `openapi_path_request_body` table provides insights into Request Bodies within OpenAPI. As a developer or API designer, you can explore details about each request body defined within an OpenAPI specification through this table, including the content type, schema, and description. Utilize it to understand the data requirements of your API endpoints and ensure they are correctly documented and implemented.

## Examples

### Basic info
Explore the API paths and methods, along with their descriptions and requirements. This will help you understand the structure and functionality of the API, including the details of the content and paths associated with each method.

```sql+postgres
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

```sql+sqlite
select
  api_path,
  api_method,
  description,
  required,
  content,
  path
from
  openapi_path_request_body;
```

### Get the request body object for a specific API operation
Explore the details of a specific API operation to understand its requirements and contents. This is particularly useful in scenarios where you need to understand the structure and requirements of an API call before making it.

```sql+postgres
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

```sql+sqlite
select
  api_path,
  api_method,
  description,
  required,
  content,
  path
from
  openapi_path_request_body
where
  api_path = '/applications/{client_id}/token/post';
```

### List request body definitions without schema
Discover the segments that lack schema in your API request body definitions. This aids in identifying potential areas for schema inclusion, improving data validation and consistency within your API.

```sql+postgres
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

```sql+sqlite
select
  path,
  api_path || '.requestBody.content.' || json_extract(c.value, '$.contentType') as paths
from
  openapi_path_request_body,
  json_each(content) as c
where
  json_extract(c.value, '$.schema') is null
  and request_body_ref is null;
```