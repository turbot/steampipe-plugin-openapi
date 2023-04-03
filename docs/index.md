---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/openapi.svg"
brand_color: "#4C5B31"
display_name: "OpenAPI"
short_name: "openapi"
description: "Steampipe plugin to query introspection of the OpenAPI definition."
og_description: "Query OpenAPI files with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/openapi-social-graphic.png"
---

# OpenAPI + Steampipe

A self-contained or composite resource which defines or describes an API or elements of an API.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query data using SQL.

Query all the endpoints available for an API:

```sql
select
  api_path,
  operation_id,
  summary,
  deprecated,
  tags
from
  openapi_path;
```

```sh
+----------+---------------------+--------------------------+------------+--------+
| api_path | operation_id        | summary                  | deprecated | tags   |
+----------+---------------------+--------------------------+------------+--------+
| /get     | listVersionsv2      | List API versions        | false      | <null> |
| /v2/get  | getVersionDetailsv2 | Show API version details | false      | <null> |
+----------+---------------------+--------------------------+------------+--------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/openapi/tables)**

## Quick start

### Install

Download and install the latest OpenAPI plugin:

```bash
steampipe plugin install openapi
```

### Credentials

No credentials are required.

### Configuration

Installing the latest openAPI plugin will create a config file (`~/.steampipe/config/openapi.spc`) with a single connection named `openapi`:

```hcl
connection "openapi" {
  plugin = "openapi"

  # Paths is a list of locations to search for OpenAPI definition files
  # Paths can be configured with a local directory, a remote Git repository URL, or an S3 bucket URL
  # Refer https://hub.steampipe.io/plugins/turbot/openapi#supported-path-formats for more information
  # Wildcard based searches are supported, including recursive searches
  # Local paths are resolved relative to the current working directory (CWD)

  # For example:
  #  - "*.json" matches all OpenAPI JSON definition files in the CWD
  #  - "**/*.json" matches all OpenAPI JSON definition files in the CWD and all sub-directories
  #  - "../*.json" matches all OpenAPI JSON definition files in the CWD's parent directory
  #  - "*.yml" or "*.yaml" matches all OpenAPI YML or YAML definition files in the CWD
  #  - "**/*.yml" or "**/*.yaml" matches all OpenAPI YML or YAML definition files in the CWD and all sub-directories
  #  - "../*.yml" or "../*.yaml" matches all OpenAPI YML or YAML definition files in the CWD's parent directory
  #  - "steampipe*.json" or "steampipe*.yml" or "steampipe*.yaml" matches all OpenAPI definition files starting with "steampipe" in the CWD
  #  - "/path/to/dir/*.json" or "/path/to/dir/*.yml" or "/path/to/dir/*.yaml" matches all OpenAPI definition files in a specific directory
  #  - "/path/to/dir/main.json" or "/path/to/dir/main.yml" or "/path/to/dir/main.yaml" matches a specific file

  # If paths includes "*", all files (including non-OpenAPI definition files) in
  # the CWD will be matched, which may cause errors if incompatible file types exist

  # Defaults to CWD
  paths = [ "*.json", "*.yml", "*.yaml" ]
}
```

### Supported Path Formats

The `paths` config argument is flexible and can search for OpenAPI definition files from several different sources, e.g., local directory paths, Git, S3.

The following sources are supported:

- [Local files](#configuring-local-file-paths)
- [Remote Git repositories](#configuring-remote-git-repository-urls)
- [S3](#configuring-s3-urls)

Paths may [include wildcards](https://pkg.go.dev/path/filepath#Match) and support `**` for recursive matching. For example:

```hcl
connection "openapi" {
  plugin = "openapi"

  paths = [
    "*.json",
    "~/*.json",
    "github.com/OAI/OpenAPI-Specification//examples/v2.0/json//*.json",
    "github.com/OAI/OpenAPI-Specification//examples/v2.0/yaml//*.yaml",
    "s3::https://bucket.s3.us-east-1.amazonaws.com/test_folder//*.json"
  ]
}
```

**Note**: If any path matches on `*` without `.json` or `.yml` or `.yaml`, all files (including non-OpenAPI definition files) in the directory will be matched, which may cause errors if incompatible file types exist.

#### Configuring Local File Paths

You can define a list of local directory paths to search for OpenAPI specification files. Paths are resolved relative to the current working directory. For example:

- `*.json` or `*.yml` or `*.yaml` matches all OpenAPI definition files in the CWD.
- `**/*.json` or `**/*.yml` or `**/*.yaml` matches all OpenAPI definition files in the CWD and all sub-directories.
- `../*.json` or `../*.yml` or `../*.yaml` matches all OpenAPI definition files in the CWD's parent directory.
- `steampipe*.json` or `steampipe*.yml` or `steampipe*.yaml` matches all OpenAPI definition files starting with "steampipe" in the CWD.
- `/path/to/dir/*.json` or `/path/to/dir/*.yml` or `/path/to/dir/*.yaml` matches all OpenAPI definition files in a specific directory. For example:
  - `~/*.json` or `~/*.yml` or `~/*.yaml` matches all OpenAPI definition files in the home directory.
  - `~/**/*.json` or `~/**/*.yml` or `~/**/*.yaml` matches all OpenAPI definition files recursively in the home directory.
- `/path/to/dir/main.json` or `/path/to/dir/main.yml` or `/path/to/dir/main.yaml` matches a specific file.

```hcl
connection "openapi" {
  plugin = "openapi"

  paths = [ "*.json", "*.yml", "*.yaml", "~/*.json", "/path/to/dir/main.json" ]
}
```

#### Configuring Remote Git Repository URLs

You can also configure `paths` with any Git remote repository URLs, e.g., GitHub, BitBucket, GitLab. The plugin will then attempt to retrieve any OpenAPI definition files from the remote repositories.

For example:

- `github.com/OAI/OpenAPI-Specification//examples/v2.0/json//*.json` matches all top-level OpenAPI JSON definition files in the specified repository.
- `github.com/OAI/OpenAPI-Specification//examples/v2.0/yaml//*.yaml` matches all top-level OpenAPI YAML definition files in the specified repository.
- `github.com/OAI/OpenAPI-Specification//examples/v2.0//**/*.json` matches all OpenAPI definition files in the specified repository and all subdirectories.

You can specify a subdirectory after a double-slash (`//`) if you want to download only a specific subdirectory from a downloaded directory.

```hcl
connection "openapi" {
  plugin = "openapi"

  paths = [ 
    "github.com/OAI/OpenAPI-Specification//examples/v2.0/json//*.json",
    "github.com/OAI/OpenAPI-Specification//examples/v2.0/yaml//*.yaml"
  ]
}
```

Similarly, you can define a list of GitLab and BitBucket URLs to search for OpenAPI definition files:

```hcl
connection "openapi" {
  plugin = "openapi"

  paths = [
    "github.com/OAI/OpenAPI-Specification//examples/v2.0/json//*.json",
    "github.com/OAI/OpenAPI-Specification//examples/v2.0/yaml//*.yaml"
  ]
}
```

#### Configuring S3 URLs

You can also query all OpenAPI definition files stored inside an S3 bucket (public or private) using the bucket URL.

##### Accessing a Private Bucket

In order to access your files in a private S3 bucket, you will need to configure your credentials. You can use your configured AWS profile from local `~/.aws/config`, or pass the credentials using the standard AWS environment variables, e.g., `AWS_PROFILE`, `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_REGION`.

We recommend using AWS profiles for authentication.

**Note:** Make sure that `region` is configured in the config. If not set in the config, `region` will be fetched from the standard environment variable `AWS_REGION`.

You can also authenticate your request by setting the AWS profile and region in `paths`. For example:

```hcl
connection "openapi" {
  plugin = "openapi"

  paths = [
    "s3::https://bucket-2.s3.us-east-1.amazonaws.com//*.json?aws_profile=<AWS_PROFILE>",
    "s3::https://bucket-2.s3.us-east-1.amazonaws.com/test_folder//*.yaml?aws_profile=<AWS_PROFILE>"
  ]
}
```

**Note:**

In order to access the bucket, the IAM user or role will require the following IAM permissions:

- `s3:ListBucket`
- `s3:GetObject`
- `s3:GetObjectVersion`

If the bucket is in another AWS account, the bucket policy will need to grant access to your user or role. For example:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "ReadBucketObject",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::123456789012:user/YOUR_USER"
      },
      "Action": ["s3:ListBucket", "s3:GetObject", "s3:GetObjectVersion"],
      "Resource": ["arn:aws:s3:::test-bucket1", "arn:aws:s3:::test-bucket1/*"]
    }
  ]
}
```

##### Accessing a Public Bucket

Public access granted to buckets and objects through ACLs and bucket policies allows any user access to data in the bucket. We do not recommend making S3 buckets public, but if there are specific objects you'd like to make public, please see [How can I grant public read access to some objects in my Amazon S3 bucket?](https://aws.amazon.com/premiumsupport/knowledge-center/read-access-objects-s3-bucket/).

You can query any public S3 bucket directly using the URL without passing credentials. For example:

```hcl
connection "openapi" {
  plugin = "openapi"

  paths = [
    "s3::https://bucket-1.s3.us-east-1.amazonaws.com/test_folder//*.json",
    "s3::https://bucket-2.s3.us-east-1.amazonaws.com/test_folder//**/*.yaml"
  ]
}
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-openapi
- Community: [Slack Channel](https://steampipe.io/community/join)
