# Table: openapi_component_security_scheme

The table `openapi_component_security_scheme` describes the way to define the security requirements for an API. A security scheme is a named set of security requirements that can be applied to one or more endpoints in the API.

Security schemes can be applied globally to the entire API, or to specific endpoints or operations. For each endpoint or operation, the security scheme(s) that are required to access that endpoint or operation are specified.

## Examples

### Basic info

```sql
select
  name,
  type,
  location,
  description,
  scheme,
  path,
  start_line,
  end_line
from
  openapi_component_security_scheme;
```

### List OpenAPI specs with no security scheme defined

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
