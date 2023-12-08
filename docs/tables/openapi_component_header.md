---
title: "Steampipe Table: openapi_component_header - Query OpenAPI Components using SQL"
description: "Allows users to query OpenAPI Components, specifically component headers, providing insights into the structure and details of OpenAPI components."
---

# Table: openapi_component_header - Query OpenAPI Components using SQL

OpenAPI Components provide a standardized way to design and describe APIs, enabling interoperability and consistency across different services. They are a crucial part of the OpenAPI Specification, which is a widely adopted standard for designing and documenting APIs. Components, such as headers, provide reusable building blocks for the API, reducing redundancy and promoting reuse across the API.

## Table Usage Guide

The `openapi_component_header` table provides insights into the headers of OpenAPI components. As a developer or API designer, explore header-specific details through this table, including parameters, types, and associated metadata. Utilize it to uncover information about headers, such as their data types, required status, and description, aiding in the design, documentation, and understanding of APIs.

## Examples

### Basic info
Explore the key aspects of an OpenAPI component such as its name, location, and description to gain insights into its configuration and understand its role within the overall system. This can be particularly useful in identifying potential areas for optimization or troubleshooting.

```sql+postgres
select
  key,
  name,
  location,
  description,
  path
from
  openapi_component_header;
```

```sql+sqlite
select
  key,
  name,
  location,
  description,
  path
from
  openapi_component_header;
```

### List unused header definitions
Determine unused header definitions in your API's OpenAPI specification. This helps to maintain a lean and efficient API documentation by identifying and removing redundant header definitions.

```sql+postgres
with list_available_headers as (
  select
    path,
    array_agg(key) as header_refs
  from
    openapi_component_header
  group by
    path
),
list_used_headers as (
  select
    path,
    array_agg(distinct split_part(value ->> '$ref', '/', '4')) as headers
  from
    openapi_path_response,
    jsonb_each(headers)
  where
    (value ->> '$ref') is not null
  group by
    path
),
unused_headers_definitions as (
  select path, unnest(header_refs) as data from list_available_headers
    except
  select path, unnest(headers) as data from list_used_headers
)
select
  path,
  concat('components.headers.', data) as header_ref
from
  unused_headers_definitions;
```

```sql+sqlite
Error: SQLite does not support array_agg, split_part, and unnest functions.
```

### List headers with no schema
Explore which headers in your OpenAPI components are missing a schema, helping to identify potential areas for further definition and structure. This can be particularly useful for maintaining consistency and clarity in your API documentation.

```sql+postgres
select
  key,
  location,
  deprecated,
  description,
  path
from
  openapi_component_header
where
  schema is null
  and schema_ref is null;
```

```sql+sqlite
select
  key,
  location,
  deprecated,
  description,
  path
from
  openapi_component_header
where
  schema is null
  and schema_ref is null;
```

### List deprecated headers with no alternative mentioned in the description
Identify instances where deprecated headers are being used without any alternative mentioned. This can help in updating your API usage to avoid potential issues in the future.

```sql+postgres
select
  key,
  location,
  deprecated,
  description,
  path
from
  openapi_component_header
where
  deprecated
  and description is null;
```

```sql+sqlite
select
  key,
  location,
  deprecated,
  description,
  path
from
  openapi_component_header
where
  deprecated
  and description is null;
```