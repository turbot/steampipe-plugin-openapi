---
title: "Steampipe Table: openapi_path_response - Query OpenAPI Path Responses using SQL"
description: "Allows users to query OpenAPI Path Responses, specifically the details of the responses returned by the API paths, providing insights into the API's behavior and potential issues."
---

# Table: openapi_path_response - Query OpenAPI Path Responses using SQL

OpenAPI is a specification for building APIs that offer a high degree of interoperability. The Path Response in OpenAPI provides detailed information about the responses that an API path can return. It is a critical component in understanding the behavior of the API and diagnosing potential issues.

## Table Usage Guide

The `openapi_path_response` table provides insights into the responses returned by the API paths in an OpenAPI specification. As an API developer or tester, explore response-specific details through this table, including the status codes, descriptions, and associated schema. Utilize it to uncover information about the API's behavior, such as the responses it can return, their structure, and the status codes they are associated with.

## Examples

### Basic info
Explore the status of responses from various API paths and methods. This can help in identifying any irregularities or issues in the API's responses, improving overall system monitoring and troubleshooting.

```sql
select
  api_path,
  api_method,
  response_status,
  jsonb_pretty(content),
  path
from
  openapi_path_response;
```

### Get the response object for a specific API operation on success
Determine the success status of a specific API operation by analyzing the response object. This is useful for troubleshooting and optimizing API calls to ensure they are functioning as expected.

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
Explore which paths in the OpenAPI path response lack a defined schema. This is useful in identifying areas of your API that may need further definition or refinement.

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