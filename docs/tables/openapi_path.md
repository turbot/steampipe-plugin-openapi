---
title: "Steampipe Table: openapi_path - Query OpenAPI Paths using SQL"
description: "Allows users to query OpenAPI Paths, providing insights into the available API endpoints and their configurations."
---

# Table: openapi_path - Query OpenAPI Paths using SQL

OpenAPI is a specification for machine-readable interface files for describing, producing, consuming, and visualizing RESTful web services. It provides a way to describe the capabilities of a service in a standard, language-agnostic manner. This allows both humans and computers to discover and understand the capabilities of a service without requiring access to source code, additional documentation, or inspection of network traffic.

## Table Usage Guide

The `openapi_path` table provides insights into the paths defined within an OpenAPI specification. As a developer or API designer, explore path-specific details through this table, including the available operations, parameters, and responses. Utilize it to uncover information about the API's structure, such as the available endpoints, the HTTP methods they support, and the expected request and response formats.

## Examples

### Basic info
Explore the API paths in a project to determine if any are outdated or no longer in use. This can assist in maintaining the efficiency of your code by identifying and removing unnecessary elements.

```sql
select
  api_path,
  deprecated,
  description,
  tags,
  path
from
  openapi_path;
```

### List all deprecated endpoints
Uncover the details of outdated API endpoints in your application. This can help you identify areas that may need updates or replacement, ensuring your application stays current and secure.

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
Explore all endpoints that use the GET method to understand how your API interacts with data. This can help optimize the performance and security of your API by identifying potential areas for improvement.

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
Explore the required parameters for a specific endpoint to better understand its necessary inputs and structure. This is particularly useful for ensuring proper API calls and data retrieval.

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
Explore the structure of a successful response from a specific API endpoint. This can be useful to understand the data format and fields returned upon successful API calls, aiding in the development of applications that interact with this endpoint.

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