## v1.0.0 [2024-10-22]

There are intentionally no significant changes in this plugin version, but it has been released to coincide with the [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin follows [semantic versioning's specification](https://semver.org/#semantic-versioning-specification-semver) and preserves backward compatibility in each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#28](https://github.com/turbot/steampipe-plugin-ansible/pull/28))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#28](https://github.com/turbot/steampipe-plugin-ansible/pull/28))

## v0.2.1 [2023-12-12]

_Bug fixes_

- Fixed the missing optional tags on connection config parameters.

## v0.2.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/downloads), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#23](https://github.com/turbot/steampipe-plugin-ansible/pull/23))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension.
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-ansible/blob/main/docs/LICENSE). ([#23](https://github.com/turbot/steampipe-plugin-ansible/pull/23))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#21](https://github.com/turbot/steampipe-plugin-ansible/pull/21))

## v0.1.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#10](https://github.com/turbot/steampipe-plugin-ansible/pull/10))

## v0.1.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#7](https://github.com/turbot/steampipe-plugin-ansible/pull/7))
- Recompiled plugin with Go version `1.21`. ([#7](https://github.com/turbot/steampipe-plugin-ansible/pull/7))

## v0.0.1 [2023-08-04]

_What's new?_

- New tables added
  - [ansible_group](https://hub.steampipe.io/plugins/turbot/ansible/tables/ansible_group)
  - [ansible_host](https://hub.steampipe.io/plugins/turbot/ansible/tables/ansible_host)
  - [ansible_playbook](https://hub.steampipe.io/plugins/turbot/ansible/tables/ansible_playbook)
  - [ansible_task](https://hub.steampipe.io/plugins/turbot/ansible/tables/ansible_task)
