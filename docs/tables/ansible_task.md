---
title: "Steampipe Table: ansible_task - Query Ansible Tasks using SQL"
description: "Allows users to query Ansible Tasks, specifically data about the tasks executed in Ansible playbooks, providing insights into task details and execution status."
---

# Table: ansible_task - Query Ansible Tasks using SQL

Ansible is an open-source software provisioning, configuration management, and application-deployment tool. It provides large productivity gains to a wide variety of automation challenges. A key component of Ansible is its Tasks, which are units of action in Ansible.

## Table Usage Guide

The `ansible_task` table provides insights into tasks within Ansible. As a DevOps engineer, explore task-specific details through this table, including the task name, host, status, and associated metadata. Utilize it to uncover information about tasks, such as their execution status, the hosts they are associated with, and the specific details of each task.

## Examples

### Retrieve all tasks in a playbook
Explore which tasks within a playbook require escalated privileges. This can help identify areas where potential security risks may exist, allowing for a more secure configuration of your playbook.

```sql
select
  name as task_name,
  tags,
  become,
  become_user
  path
from
  ansible_task
where
  playbook_name = 'Playbook';
```

### List tasks that use privilege escalation
Discover the segments that use privilege escalation in Ansible tasks. This is beneficial to identify areas where elevated permissions are granted, allowing for a review of security practices.

```sql
select
  name as task_name,
  tags,
  become,
  become_user
  path
from
  ansible_task
where
  become;
```

### Lists tasks with a specific tag
Explore which tasks are associated with a specific tag in Ansible to better manage and organize your automation scripts.

```sql
select
  name as task_name,
  tags,
  become,
  become_user
  path
from
  ansible_task
where
  tags ?| array['create_user'];
```

### Lists tasks with a specific connection type
Explore which tasks within your Ansible setup are utilizing SSH as their connection type. This can be useful in identifying potential security vulnerabilities or for routine auditing of your network connections.

```sql
select
  name as task_name,
  tags,
  become,
  become_user
  path
from
  ansible_task
where
  connection = 'ssh';
```

### List tasks that use `root` privilege
Identify instances where tasks are using elevated privileges, such as 'root', within Ansible. This can help in assessing security risks and ensuring adherence to best practices.

```sql
select
  name as task_name,
  tags,
  become,
  become_user
  path
from
  ansible_task
where
  become
  and (
    become_user is null
    or become_user = 'root'
  );
```