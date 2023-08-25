# Table: openapi_server

The table `openapi_server` describes the network addresses (e.g. URLs) where the API is available. It provides information about the base URL(s) for the API.

## Examples

### Basic info

```sql
select
  url as base_url,
  description,
  variables,
  path,
  start_line,
  end_line
from
  openapi_server;
```

### Get the variables used for substitution in the server's URL

```sql
select
  url as base_url,
  jsonb_pretty(variables) as variables,
  path
from
  openapi_server;
```
