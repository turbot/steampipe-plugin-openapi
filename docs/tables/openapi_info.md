---
title: "Steampipe Table: openapi_info - Query OpenAPI Information using SQL"
description: "Allows users to query OpenAPI Information, specifically the version, title, description, and terms of service of the OpenAPI specification."
---

# Table: openapi_info - Query OpenAPI Information using SQL

OpenAPI is a specification for machine-readable interface files for describing, producing, consuming, and visualizing RESTful web services. The OpenAPI Information resource provides details about the OpenAPI specification, including its version, title, description, and terms of service. This information is useful for understanding the specifics of the RESTful web service described by the OpenAPI specification.

## Table Usage Guide

The `openapi_info` table provides insights into the OpenAPI specification of a RESTful web service. As a software developer or a DevOps engineer, explore specification-specific details through this table, including the version, title, description, and terms of service of the OpenAPI specification. Utilize it to uncover information about the specification, such as its title and description, and the version and terms of service associated with it.

## Examples

### Basic info
Explore the general information of an OpenAPI specification to gain insights into its title, description, version, contact details, and paths. This can be useful for understanding the overall structure and details of the API at a glance.

```sql
select
  title,
  description,
  version,
  contact,
  path
from
  openapi_info;
```

### Get the maintainer's contact information
Explore the contact information of the maintainer to ensure effective communication and collaboration. This could be particularly beneficial for troubleshooting, seeking clarifications, or discussing potential enhancements.

```sql
select
  title,
  version,
  contact ->> 'name' as maintainer_name,
  contact ->> 'email' as maintainer_email,
  path
from
  openapi_info;
```

### List API specifications not using any license
Explore which API specifications are not currently using any license. This can help identify potential compliance issues or areas where licensing needs to be updated or added.

```sql
select
  title,
  version,
  contact ->> 'name' as maintainer_name,
  contact ->> 'email' as maintainer_email,
  path
from
  openapi_info
where
  license is null;
```