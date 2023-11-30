---
title: "Steampipe Table: openapi_component_schema - Query OpenAPI Components using SQL"
description: "Allows users to query OpenAPI Components, specifically the schema details, providing insights into API structure and data models."
---

# Table: openapi_component_schema - Query OpenAPI Components using SQL

OpenAPI Components are reusable entities within the OpenAPI definition. They provide a way to reuse definitions and parameters across multiple endpoints, reducing duplication and promoting consistency. Components can include schemas, which describe the structure of an API's data model.

## Table Usage Guide

The `openapi_component_schema` table provides insights into the OpenAPI Components within an API specification. As an API developer or architect, explore schema-specific details through this table, including data types, properties, and associated metadata. Utilize it to understand the structure of your API's data models, such as complex object definitions, data validation rules, and the relationships between different schemas.

## Examples

### Basic info
Explore the basic elements of an OpenAPI component schema to understand its structure and contents. This can help determine which components might be deprecated or have default values, aiding in the maintenance and updating of the schema.

```sql
select
  name,
  type,
  format,
  deprecated,
  description,
  default_value,
  path
from
  openapi_component_schema;
```

### Get the properties returned by a specific API on success
Determine the details and structure of successful responses from a specific API endpoint. This can be useful for understanding what data is returned upon successful API calls, which can aid in further API integration and data management.

```sql
with get_schema_ref as (
  select
    api_path,
    r.key as response_status,
    r.value -> 'content' -> 'application/json' -> 'schema' ->> '$ref' as schema_ref
  from
    openapi_path,
    jsonb_each(responses) as r
  where
    api_path = '/repos/{owner}/{repo}/issues/post'
    and r.key::integer >= '201' and r.key::integer < 300
)
select
  r.api_path,
  s.name,
  jsonb_pretty(s.properties) as schema_property,
  jsonb_pretty(s.required) as required,
  s.description
from
  get_schema_ref as r
  join openapi_component_schema as s on r.schema_ref = concat('#/components/schemas/', s.name);
```

### Get the schema of a required parameter of a specific API
This example helps you understand the structure of a specific API's required parameter. It's useful when you need to know what information is necessary to successfully use or interact with a particular API.

```sql
select
  op.api_path,
  cp.required,
  jsonb_pretty(cp.schema) as schema
from
  openapi_path as op,
  jsonb_array_elements(op.parameters) as p
  join openapi_component_parameter as cp on (p ->> '$ref') = concat('#/components/parameters/', cp.name)
where
  op.api_path = '/orgs/{org}/members/{username}/delete'
  and cp.required;
```