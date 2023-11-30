---
title: "Steampipe Table: openapi_component_security_scheme - Query OpenAPI Component Security Schemes using SQL"
description: "Allows users to query OpenAPI Component Security Schemes, providing insights into the security requirements for different API operations."
---

# Table: openapi_component_security_scheme - Query OpenAPI Component Security Schemes using SQL

OpenAPI Component Security Schemes are part of the OpenAPI Specification (OAS), which is a standard, language-agnostic interface to RESTful APIs. These security schemes define the security requirements for different API operations, such as authentication and authorization methods. They are essential for ensuring the secure and controlled access to API resources.

## Table Usage Guide

The `openapi_component_security_scheme` table provides insights into the security requirements of RESTful APIs defined by the OpenAPI Specification. As an API developer or security analyst, explore this table to understand the security schemes applied to different API operations, including the types of schemes (e.g., HTTP, OAuth2, OpenID Connect) and their specific details. Utilize it to review and validate the security configurations of your APIs, and to identify potential security risks or misconfigurations.

## Examples

### Basic info
Explore the configuration details of security schemes in your OpenAPI components to understand their location, type, and purpose. This can help assess the elements within your API security design and ensure they are properly implemented.

```sql
select
  name,
  type,
  location,
  description,
  scheme,
  path
from
  openapi_component_security_scheme;
```

### List OpenAPI specs with no security scheme defined
Discover the segments of your OpenAPI specifications that lack defined security schemes. This is useful for identifying potential vulnerabilities and ensuring all parts of your API are secure.

```sql
select
  i.title,
  i.version,
  i.description,
  i.path
from
  openapi_info as i
  left join openapi_component_security_scheme as s on i.path = s.path
where
  s.path is null;
```

### List OAuth 1.0 security schemes
Explore the security schemes that utilize OAuth 1.0 for HTTP protocols. This can be useful in understanding the security mechanisms in place and identifying any potential areas for improvement.

```sql
select
  name,
  type,
  location,
  description,
  scheme,
  path
from
  openapi_security_scheme
where
  type = 'http'
  and scheme = 'oauth';
```

### List security schemes using basic HTTP authentication
Explore which security schemes utilize basic HTTP authentication to gain insights into potential vulnerabilities or areas requiring additional security measures. This is useful in identifying weak spots in your system's security and implementing necessary improvements.

```sql
select
  name,
  type,
  location,
  description,
  scheme,
  path
from
  openapi_security_scheme
where
  type = 'http'
  and scheme = 'basic';
```