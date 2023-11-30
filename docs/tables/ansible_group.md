---
title: "Steampipe Table: ansible_group - Query Ansible Groups using SQL"
description: "Allows users to query Ansible Groups, specifically the group details and associated hosts, providing insights into the configuration and management of Ansible Groups."
---

# Table: ansible_group - Query Ansible Groups using SQL

Ansible is an open-source software provisioning, configuration management, and application-deployment tool. It provides large productivity gains to a wide variety of automation challenges. An Ansible Group is a collection of hosts, which can be managed collectively, rather than individually.

## Table Usage Guide

The `ansible_group` table provides insights into Ansible Groups within Ansible configuration management. As a DevOps engineer, explore group-specific details through this table, including group names, hosts, and associated metadata. Utilize it to uncover information about groups, such as associated hosts, to aid in the management and configuration of Ansible Groups.

## Examples

### Query a simple file
Discover the segments that are part of a specific group in an Ansible inventory, gaining insights into the hierarchical relationships and variable configurations. This can aid in understanding the structure and configuration of your Ansible deployments.
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

Query to retrieve the groups:


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
Explore the hierarchical structure within a group to understand its associations. This is useful for identifying the parent-child relationships within a group, which can help in managing or visualizing the group's organization.

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
Identify the groups that are associated with a specific server in your network. This can help in managing and organizing your resources effectively.

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
Discover the segments that have active hosts within a certain group. This is beneficial for efficiently managing resources and ensuring optimal utilization.

```sql
select
  name,
  hosts
from
  ansible_group
where
  hosts is not null;
```

### List groups with no children
Explore which Ansible groups do not have any child groups. This can be useful for identifying isolated groups, potentially simplifying your infrastructure management tasks.

```sql
select
  name,
  hosts
from
  ansible_group
where
  children is null;
```