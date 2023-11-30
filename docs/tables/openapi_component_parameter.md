---
title: "Steampipe Table: openapi_component_parameter - Query OpenAPI Components using SQL"
description: "Allows users to query OpenAPI Components, specifically the parameters of each component, providing insights into the structure and details of the API."
---

# Table: openapi_component_parameter - Query OpenAPI Components using SQL

OpenAPI is a specification for machine-readable interface files for describing, producing, consuming, and visualizing RESTful web services. It provides a way to describe and document RESTful APIs in a common language that everyone can understand. It is particularly useful for creating API documentation, generating code and ensuring that the APIs you build are simple, fast, and consistently well-structured.

## Table Usage Guide

The `openapi_component_parameter` table provides insights into the parameters of each component within OpenAPI. As a developer or API architect, explore parameter-specific details through this table, including their names, descriptions, and whether they are required. Utilize it to uncover information about parameters, such as their data types, default values, and whether they allow empty values, helping to ensure your APIs are well-structured and follow best practices.

## Examples

### Basic info
Explore the parameters of an OpenAPI component to understand their usage, including whether they are required or deprecated. This can assist in identifying areas for potential improvements or updates within your API structure.

```sql
select
  name,
  description,
  location,
  deprecated,
  required,
  schema,
  path
from
  openapi_component_parameter;
```

### List all supported parameters for a specific API endpoint
Explore the different specifications of parameters for a particular API endpoint. This is useful to understand the requirements and potential deprecations for each parameter, aiding in effective API usage and maintenance.

```sql
with list_parameters as (
  select
    api_path,
    p ->> '$ref' as parameter_ref
  from
    openapi_path,
    jsonb_array_elements(parameters) as p
  where
    api_path = '/repos/{owner}/{repo}/issues/post'
)
select
  l.api_path,
  p.name as parameter_name,
  p.required as is_required,
  p.deprecated as is_deprecated,
  jsonb_pretty(p.schema) as parameter_schema
from
  list_parameters as l
  join openapi_component_parameter as p on l.parameter_ref = concat('#/components/parameters/', p.name);
```

### List parameters with no schema
Discover the segments that have parameters without a defined schema, which can help identify potential areas of improvement or inconsistencies within your OpenAPI components. This could be particularly useful in maintaining or upgrading your API systems.

```sql
select
  name,
  description,
  location,
  deprecated,
  required,
  schema,
  path
from
  openapi_component_parameter
where
  schema is null
  and schema_ref is null;
```

### List deprecated parameters with no alternative mentioned in the description
Explore which parameters in your OpenAPI components are deprecated but lack a description, helping you to identify potential issues in your API documentation and ensure all deprecated parameters are correctly documented for future reference.

```sql
select
  name,
  description,
  location,
  deprecated,
  required,
  schema,
  path
from
  openapi_component_parameter
where
  deprecated
  and description is null;
```