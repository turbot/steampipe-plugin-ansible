# Table: ansible_group

An Ansible Inventory Group is a logical collection of hosts that share common properties, attributes, and purposes within an Ansible playbook. It allows users to organize the infrastructure into meaningful segments, making it easier to manage and configure multiple hosts simultaneously.

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
  hosts,
  parents,
  vars
from
  ansible_group;
```

```sh
+-----------+---------------------------+---------------------------+--------------------------------------------------------------------------------------------------------------------------+
| name      | hosts                     | parents                   | vars                                                                                                                     |
+-----------+---------------------------+---------------------------+--------------------------------------------------------------------------------------------------------------------------+
| southeast | <null>                    | ["usa","all"]             | {"escape_pods":"2","halon_system_timeout":"30","self_destruct_countdown":"60","some_server":"foo.southeast.example.com"} |
| northwest | <null>                    | ["usa","all"]             | {}                                                                                                                       |
| all       | ["host1","host2","host3"] | <null>                    | {}                                                                                                                       |
| ungrouped | <null>                    | ["all"]                   | {}                                                                                                                       |
| Raleigh   | <null>                    | ["southeast","usa","all"] | {"escape_pods":"2","halon_system_timeout":"30","self_destruct_countdown":"60","some_server":"foo.southeast.example.com"} |
| northeast | <null>                    | ["usa","all"]             | {}                                                                                                                       |
| raleigh   | ["host2","host3"]         | ["all"]                   | {}                                                                                                                       |
| southwest | <null>                    | ["usa","all"]             | {}                                                                                                                       |
| atlanta   | ["host1","host2"]         | ["all"]                   | {}                                                                                                                       |
| Atlanta   | <null>                    | ["southeast","usa","all"] | {"escape_pods":"2","halon_system_timeout":"30","self_destruct_countdown":"60","some_server":"foo.southeast.example.com"} |
| usa       | <null>                    | ["all"]                   | {}                                                                                                                       |
+-----------+---------------------------+---------------------------+--------------------------------------------------------------------------------------------------------------------------+
```

### List groups along with its parents and children

```sql
select
  name,
  parents,
  children
from
  ansible_group;
```

```sh
+-----------+---------------------------+-------------------------------------------------------------------------------------------------------------+
| name      | parents                   | children                                                                                                    |
+-----------+---------------------------+-------------------------------------------------------------------------------------------------------------+
| southeast | ["usa","all"]             | ["Atlanta","Raleigh"]                                                                                       |
| northwest | ["usa","all"]             | <null>                                                                                                      |
| all       | <null>                    | ["Atlanta","Raleigh","northeast","raleigh","southwest","usa","ungrouped","atlanta","southeast","northwest"] |
| ungrouped | ["all"]                   | <null>                                                                                                      |
| Raleigh   | ["southeast","usa","all"] | <null>                                                                                                      |
| northeast | ["usa","all"]             | <null>                                                                                                      |
| raleigh   | ["all"]                   | <null>                                                                                                      |
| southwest | ["usa","all"]             | <null>                                                                                                      |
| atlanta   | ["all"]                   | <null>                                                                                                      |
| Atlanta   | ["southeast","usa","all"] | <null>                                                                                                      |
| usa       | ["all"]                   | ["southwest","northwest","Atlanta","Raleigh","northeast","southeast"]                                       |
+-----------+---------------------------+-------------------------------------------------------------------------------------------------------------+
```

### Find groups with a specific variable key-value pair

```sql
select
  name,
  parents,
  vars
from
  ansible_group
where
  vars ->> 'some_server' = 'foo.southeast.example.com';
```

### List hosts per group

```sql
select
  name,
  hosts
from
  ansible_group
where
  hosts is not null;
```

### List group with no children

```sql
select
  name,
  hosts
from
  ansible_group
where
  children is null;
```
