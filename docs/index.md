---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/ansible.svg"
brand_color: "#1A1918"
display_name: "Ansible"
short_name: "ansible"
description: "Steampipe plugin to query configurations from the Ansible playbooks."
og_description: "Query Ansible playbooks files with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/ansible-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Ansible + Steampipe

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

[Ansible](https://www.ansible.com) offers open-source automation that is simple, flexible, and powerful.

The Ansible plugin makes it simpler to query the configured Ansible playbook files, and the various tasks defined in it. Apart from scanning the playbook files, the plugin also supports scanning the Ansible inventory files from different sources.

List all playbooks that use privilege escalation:

```sql
select
  name,
  hosts,
  jsonb_pretty(tasks) as tasks
from
  ansible_playbook
where
  become;
```

```
+----------+-------------+----------------------------------------------------------+
| name     | hosts       | tasks                                                    |
+----------+-------------+----------------------------------------------------------+
| Playbook | web_servers | [                                                        |
|          |             |     {                                                    |
|          |             |         "yum": {                                         |
|          |             |             "name": "httpd",                             |
|          |             |             "state": "latest"                            |
|          |             |         },                                               |
|          |             |         "name": "ensure apache is at the latest version" |
|          |             |     },                                                   |
|          |             |     {                                                    |
|          |             |         "name": "ensure apache is running",              |
|          |             |         "service": {                                     |
|          |             |             "name": "httpd",                             |
|          |             |             "state": "started"                           |
|          |             |         }                                                |
|          |             |     }                                                    |
|          |             | ]                                                        |
+----------+-------------+----------------------------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/ansible/tables)**

## Quick start

### Install

Download and install the latest Ansible plugin:

```sh
steampipe plugin install ansible
```

### Credentials

No credentials are required.

### Configuration

Installing the latest ansible plugin will create a config file (`~/.steampipe/config/ansible.spc`) with a single connection named `ansible`:

Configure your file paths in `~/.steampipe/config/ansible.spc`:

```hcl
connection "ansible" {
  plugin = "ansible"

  # The plugin supports parsing both Ansible playbook files as well as inventory files.
  # For example:
  #  - To parse the Ansible playbook files, use `playbook_file_paths` argument to configure it.
  #  - Similarly, to parse the Ansible inventory files, use `inventory_file_paths`.

  # The above paths can be configured with a local directory, a remote Git repository URL, or an S3 bucket URL
  # Wildcard based searches are supported, including recursive searches
  # Local paths are resolved relative to the current working directory (CWD)

  # For example:
  #  - "*.yml" matches all Ansible playbook files in the CWD
  #  - "**/*.yml" matches all Ansible playbook files in the CWD and all sub-directories
  #  - "../*.yml" matches all Ansible playbook files in the CWD's parent directory
  #  - "steampipe*.yml" matches all Ansible playbook files starting with "steampipe" in the CWD
  #  - "/path/to/dir/*.yml" matches all Ansible playbook files in a specific directory
  #  - "/path/to/dir/main.yml" matches a specific file

  # If paths includes "*", all files (including non-Ansible playbook files) in
  # the CWD will be matched, which may cause errors if incompatible file types exist

  # Defaults to CWD
  playbook_file_paths  = [ "*.yml", "*.yaml" ]
  inventory_file_paths = [ "/etc/ansible/hosts", "~/.ansible/hosts" ]
}
```

## Configuring File Paths

The plugin supports scanning both Ansible playbook files and the inventory files. For scanning the files, configure the plugin config file with the desired file paths. For example:

- For scanning the Ansible playbook files, use `playbook_file_paths` argument to configure it.
- For scanning the Ansible inventory files, use `inventory_file_paths` argument to configure it.

Both `playbook_file_paths` and `inventory_file_paths` config arguments are flexible and can search for Ansible playbook files from various sources (e.g., [Local files](#configuring-local-file-paths), [Git](#configuring-remote-git-repository-urls), [S3](#configuring-s3-urls) etc.).

Paths may [include wildcards](https://pkg.go.dev/path/filepath#Match) and support `**` for recursive matching. For example:

```hcl
connection "ansible" {
  plugin = "ansible"

  playbook_file_paths = [
    "*.yml",
    "~/*.yaml",
    "github.com/ansible-community/molecule//playbooks//*.yaml",
    "s3::https://bucket.s3.us-east-1.amazonaws.com/test_folder//*.yaml"
  ]

  inventory_file_paths = ["*.ini", "~/*.ini"]
}
```

**Note**: If any path matches on `*` with `.yml` or `.yaml`, all files (including non-Ansible playbook files) in the directory will be matched, which may cause errors if incompatible file types exist.

### Configuring Local File Paths

You can define a list of local directory paths to search for Ansible playbook files. Paths are resolved relative to the current working directory. For example:

- `*.yml` or `*.yaml` matches all Ansible playbook files in the CWD.
- `**/*.yml` or `**/*.yaml` matches all Ansible playbook files in the CWD and all sub-directories.
- `../*.yml` or `../*.yaml` matches all Ansible playbook files in the CWD's parent directory.
- `steampipe*.yml` or `steampipe*.yaml` matches all Ansible playbook files starting with "steampipe" in the CWD.
- `/path/to/dir/*.yml` or `/path/to/dir/*.yaml` matches all Ansible playbook files in a specific directory. For example:
  - `~/*.yml` or `~/*.yaml` matches all Ansible playbook files in the home directory.
  - `~/**/*.yml` or `~/**/*.yaml` matches all Ansible playbook files recursively in the home directory.
- `/path/to/dir/main.yml` or `/path/to/dir/main.yaml` matches a specific file.

```hcl
connection "ansible" {
  plugin = "ansible"

  playbook_file_paths = [ "*.yml", "*.yaml", "/path/to/dir/playbook.yaml" ]
}
```

### Configuring Remote Git Repository URLs

You can also configure `paths` with any Git remote repository URLs, e.g., GitHub, BitBucket, GitLab. The plugin will then attempt to retrieve any Ansible playbook files from the remote repositories.

For example:

- `github.com/ansible-community/molecule//playbooks//*.yaml` matches all top-level Ansible playbook files in the specified repository.
- `github.com/ansible-community/molecule//playbooks//**/*.yaml` matches all Ansible playbook files in the specified repository and all subdirectories.

You can specify a subdirectory after a double-slash (`//`) if you want to download only a specific subdirectory from a downloaded directory.

```hcl
connection "ansible" {
  plugin = "ansible"

  playbook_file_paths = ["github.com/ansible-community/molecule//playbooks//*.yaml"]
}
```

Similarly, you can define a list of GitLab and BitBucket URLs to search for Ansible playbook files.

### Configuring S3 URLs

You can also query all Ansible playbook files stored inside an S3 bucket (public or private) using the bucket URL.

#### Accessing a Private Bucket

In order to access your files in a private S3 bucket, you will need to configure your credentials. You can use your configured AWS profile from local `~/.aws/config`, or pass the credentials using the standard AWS environment variables, e.g., `AWS_PROFILE`, `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_REGION`.

We recommend using AWS profiles for authentication.

**Note:** Make sure that `region` is configured in the config. If not set in the config, `region` will be fetched from the standard environment variable `AWS_REGION`.

You can also authenticate your request by setting the AWS profile and region in `paths`. For example:

```hcl
connection "ansible" {
  plugin = "ansible"

  playbook_file_paths = [
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

#### Accessing a Public Bucket

Public access granted to buckets and objects through ACLs and bucket policies allows any user access to data in the bucket. We do not recommend making S3 buckets public, but if there are specific objects you'd like to make public, please see [How can I grant public read access to some objects in my Amazon S3 bucket?](https://aws.amazon.com/premiumsupport/knowledge-center/read-access-objects-s3-bucket/).

You can query any public S3 bucket directly using the URL without passing credentials. For example:

```hcl
connection "ansible" {
  plugin = "ansible"

  playbook_file_paths = [
    "s3::https://bucket-1.s3.us-east-1.amazonaws.com/test_folder//*.json",
    "s3::https://bucket-2.s3.us-east-1.amazonaws.com/test_folder//**/*.yaml"
  ]
}
```
