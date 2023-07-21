![image](https://hub.steampipe.io/images/plugins/turbot/ansible-social-graphic.png)

# Ansible Plugin for Steampipe

Use SQL to query configurations from the Ansible playbooks.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/ansible)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/ansible/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-ansible/issues)

## Quick start

### Install

Download and install the latest Ansible plugin:

```bash
steampipe plugin install ansible
```

Configure your [config file](https://hub.steampipe.io/plugins/turbot/ansible#configuration).

Configure your file paths in `~/.steampipe/config/ansible.spc`:

```hcl
connection "ansible" {
  plugin = "ansible"

  # Defaults to CWD
  playbook_file_paths  = [ "*.yml", "*.yaml" ]
  inventory_file_paths = [ "/etc/ansible/hosts", "~/.ansible/hosts" ]
}
```

Run steampipe:

```shell
steampipe query
```

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-ansible.git
cd steampipe-plugin-ansible
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/ansible.spc
```

Try it!

```
steampipe query
> .inspect ansible
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-ansible/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Ansible Plugin](https://github.com/turbot/steampipe-plugin-ansible/labels/help%20wanted)
