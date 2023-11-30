---
title: "Steampipe Table: openapi_component_request_body - Query OpenAPI Component Request Bodies using SQL"
description: "Allows users to query OpenAPI Component Request Bodies, specifically the details of each component's request body schema, providing insights into the structure and data requirements of API endpoints."
---

# Table: openapi_component_request_body - Query OpenAPI Component Request Bodies using SQL

OpenAPI is a specification for building APIs that offer a high level of interoperability. A key part of this specification is the Component Request Body, which defines the structure and data requirements for the body of a request when making API calls. This provides a standardized way for APIs to communicate what data they expect, helping to ensure consistency and reliability in API interactions.

## Table Usage Guide

The `openapi_component_request_body` table provides insights into the structure and data requirements of API endpoints within an OpenAPI Specification. As an API designer or developer, explore request body component details through this table, including data types, required fields, and associated metadata. Utilize it to uncover information about the request body components, such as their structure, the types of data they expect, and the constraints on that data.

## Examples

### Basic info
Explore the essential components of an OpenAPI request body. This query is particularly useful for understanding the key elements that are required for a successful API request, along with their descriptions and content.

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
Determine the areas in which request body definitions are unused in your OpenAPI paths. This is beneficial in identifying redundant components, helping to streamline your API documentation and maintain efficiency.

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
Explore instances where request body definitions lack a schema. This can help identify potential areas in your API where data validation might be missing, thus improving overall data quality and consistency.

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