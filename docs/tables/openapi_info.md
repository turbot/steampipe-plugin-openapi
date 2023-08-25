# Table: openapi_info

The table `openapi_info` describes general information about the API, such as its title, version, and contact information for the API's maintainers. It provides context and metadata about the API that can help developers understand its purpose and functionality, as well as any terms or conditions that apply to its use.

## Examples

### Basic info

```sql
select
  title,
  description,
  version,
  contact,
  path,
  start_line,
  end_line
from
  openapi_info;
```

### Get the maintainer's contact information

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
