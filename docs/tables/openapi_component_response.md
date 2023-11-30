---
title: "Steampipe Table: openapi_component_response - Query OpenAPI Component Responses using SQL"
description: "Allows users to query OpenAPI Component Responses, specifically the details of responses defined in the 'components' section of an OpenAPI specification."
---

# Table: openapi_component_response - Query OpenAPI Component Responses using SQL

OpenAPI Component Responses are part of the OpenAPI specification, a standard, language-agnostic interface to RESTful APIs. They provide a structured way of describing the responses an API can return, allowing for better understanding and validation of the API responses. Component responses are defined under the 'components' section of an OpenAPI specification.

## Table Usage Guide

The `openapi_component_response` table provides insights into the responses defined within an OpenAPI specification. As a developer or API designer, you can explore response-specific details through this table, including response codes, descriptions, and associated metadata. Utilize it to examine and validate the structure and consistency of your API responses as defined in the OpenAPI specification.

## Examples

### Basic info
Explore the key details and descriptions within an API's response components to gain a better understanding of its structure and data. This can assist in planning how to interact with the API or troubleshoot issues.

```sql
select
  key,
  description,
  jsonb_pretty(content),
  path
from
  openapi_component_response;
```

### List unused response definitions
Determine areas in your API where certain response definitions are not being utilized. This could help streamline your API by removing unnecessary definitions, making it easier to manage and understand.

```sql
with list_used_response_defs as (
  select
    path,
    array_agg(distinct split_part(response_ref, '/', '4')) as resp
  from
    openapi_path_response
  where
    response_ref is not null
  group by
    path
),
all_responses_definition as (
  select
    path,
    array_agg(key) as resp_defs
  from
    openapi_component_response
  group by
    path
),
unused_response_definitions as (
  select path, unnest(resp_defs) as data from all_responses_definition
    except
  select path, unnest(resp) as data from list_used_response_defs
)
select
  path,
  concat('components.responses.', data) as response_ref
from
  unused_response_definitions;
```

### List response definitions without schema
Discover the segments that lack a defined schema in the response components of your OpenAPI specifications. This can aid in identifying potential inconsistencies or gaps in your API documentation.

```sql
select
  path,
  concat('components.responses.', key, '.content.', c ->> 'contentType') as paths
from
  openapi_component_response,
  jsonb_array_elements(content) as c
where
  c ->> 'schema' is null;
```