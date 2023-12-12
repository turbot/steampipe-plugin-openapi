---
title: "Steampipe Table: openapi_server - Query OpenAPI Servers using SQL"
description: "Allows users to query OpenAPI Servers, providing detailed information about the API servers defined in the OpenAPI specification."
---

# Table: openapi_server - Query OpenAPI Servers using SQL

OpenAPI is a specification for machine-readable interface files for describing, producing, consuming, and visualizing RESTful web services. Previously known as the Swagger Specification, OpenAPI enables the development of APIs that can be understood and used by both machines and humans. The `openapi_server` resource provides details about the API servers defined in the OpenAPI specification.

## Table Usage Guide

The `openapi_server` table provides insights into API servers defined in the OpenAPI specification. As a developer or API manager, explore server-specific details through this table, including server URLs, descriptions, and variables. Utilize it to uncover comprehensive information about the API servers, such as their location, the environment they're configured for, and any additional metadata.

## Examples

### Basic info
Explore the basic information of an OpenAPI server, including its base URL, description, and variables. This can be helpful in understanding the server's configuration and identifying any potential issues.

```sql+postgres
select
  url as base_url,
  description,
  variables,
  path
from
  openapi_server;
```

```sql+sqlite
select
  url as base_url,
  description,
  variables,
  path
from
  openapi_server;
```

### Get the variables used for substitution in the server's URL
Explore which variables are used to modify the server's URL, allowing you to understand how different servers are set up and configured. This is useful for determining the flexibility and customization of your server URLs.

```sql+postgres
select
  url as base_url,
  jsonb_pretty(variables) as variables,
  path
from
  openapi_server;
```

```sql+sqlite
select
  url as base_url,
  variables,
  path
from
  openapi_server;
```