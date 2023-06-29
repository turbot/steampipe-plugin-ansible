# Table: ansible_host

Ansible automates tasks on managed nodes or `hosts` in your infrastructure, using a list or group of lists known as inventory. The table `ansible_host` lists all the host defined in the inventory files.

**Note:**

- Even if you do not define any groups in your inventory file, Ansible creates two default groups: `all` and `ungrouped`.
- The `all` group contains every host. The `ungrouped` group contains all hosts that don’t have another group aside from all.
- Every host will always belong to at least 2 groups (`all` and `ungrouped` or `all` and some other group).

## Examples

### Query a simple file

Given the inventory file `/etc/ansible/hosts` with following configuration:

```bash
[atlanta]
host1
host2

[raleigh]
host2
host3

[southeast:children]
Atlanta
Raleigh

[southeast:vars]
some_server=foo.southeast.example.com
halon_system_timeout=30
self_destruct_countdown=60
escape_pods=2

[usa:children]
southeast
northeast
southwest
northwest
```

and the query is:

```sql
select
  name,
  port,
  vars
from
  ansible_host;
```

```sh
+-------+------+------+
| name  | port | vars |
+-------+------+------+
| host3 | 22   | {}   |
| host2 | 22   | {}   |
| host1 | 22   | {}   |
+-------+------+------+
```

or, if the same host is used in more than one group, you can easily identify the list of groups where it is being used:

```sql
select
  name,
  port,
  group
from
  ansible_host;
```

```sh
+-------+------+-----------------------------+
| name  | port | groups                      |
+-------+------+-----------------------------+
| host3 | 22   | ["all","raleigh"]           |
| host2 | 22   | ["raleigh","all","atlanta"] |
| host1 | 22   | ["all","atlanta"]           |
+-------+------+-----------------------------+
```

### Casting column data for analysis

Text columns can be easily cast to other types:

```sql
select
  section,
  key,
  value::bool
from
  ini_key_value
where
  path = '/Users/myuser/defaults.ini'
  and section = 'analytics'
  and key = 'check_for_updates'
  and not value::bool;
```

```sh
+-----------+-------------------+-------+
| section   | key               | value |
+-----------+-------------------+-------+
| analytics | check_for_updates | false |
+-----------+-------------------+-------+
```
