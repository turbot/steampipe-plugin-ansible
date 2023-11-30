---
title: "Steampipe Table: ansible_host - Query Ansible Hosts using SQL"
description: "Allows users to query Ansible Hosts, specifically the host details, providing insights into the configuration and status of each host."
---

# Table: ansible_host - Query Ansible Hosts using SQL

Ansible is an open-source software provisioning, configuration management, and application-deployment tool. It provides large productivity gains to a wide variety of automation challenges. This tool is very simple to use yet powerful enough to automate complex multi-tier IT application environments.

## Table Usage Guide

The `ansible_host` table provides insights into hosts within Ansible. As a DevOps engineer, explore host-specific details through this table, including host names, groups, variables, and facts. Utilize it to uncover information about hosts, such as their configuration, status, and the groups they belong to.

## Examples

### Query a simple file
Explore which hosts are being used in your network and their respective ports. This can help you manage and monitor your network infrastructure more effectively by identifying where each host is being used.
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

Query to retrieve the hosts:


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
Identify instances where automatic updates have been turned off in the analytics section of a configuration file. This is useful for ensuring that all systems are set to receive the latest updates and features.
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