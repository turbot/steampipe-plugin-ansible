---
title: "Steampipe Table: ansible_playbook - Query Ansible Playbooks using SQL"
description: "Allows users to query Ansible Playbooks, specifically the details of each playbook, providing insights into the automation scripts and potential issues."
---

# Table: ansible_playbook - Query Ansible Playbooks using SQL

Ansible Playbook is a set of instructions that Ansible will execute on the target host or hosts. It is the primary mechanism for system configuration management in Ansible and is written in YAML. Playbooks can declare configurations, orchestrate steps of any manual ordered process, and even interact with other tools and services.

## Table Usage Guide

The `ansible_playbook` table provides insights into playbooks within Ansible. As a DevOps engineer, explore playbook-specific details through this table, including the tasks, handlers, and associated metadata. Utilize it to uncover information about playbooks, such as those with errors, the sequence of tasks, and the verification of handlers.

## Examples

### Retrieve all playbooks
Explore which playbooks are available in your Ansible configuration. This allows you to gain insights into the tasks, variables, and hosts associated with each playbook, and understand their respective paths.

```sql+postgres
select
  name,
  hosts,
  tasks,
  vars,
  path
from
  ansible_playbook;
```

```sql+sqlite
select
  name,
  hosts,
  tasks,
  vars,
  path
from
  ansible_playbook;
```

### List playbooks targeting specific hosts
Explore which ansible playbooks are specifically targeting your web servers. This can help you manage and optimize the deployment of updates or changes across your server infrastructure.

```sql+postgres
select
  name,
  hosts,
  tasks,
  vars,
  path
from
  ansible_playbook
where
  hosts = 'web_servers';
```

```sql+sqlite
select
  name,
  hosts,
  tasks,
  vars,
  path
from
  ansible_playbook
where
  hosts = 'web_servers';
```

### List playbooks that use privilege escalation
Explore which Ansible playbooks are using privilege escalation. This can be helpful to assess security practices and identify potential areas of risk in your infrastructure setup.

```sql+postgres
select
  name,
  hosts,
  tasks,
  vars,
  path
from
  ansible_playbook
where
  become;
```

```sql+sqlite
select
  name,
  hosts,
  tasks,
  vars,
  path
from
  ansible_playbook
where
  become = 1;
```

### List playbooks with no handlers
Explore which Ansible playbooks lack handlers, providing a way to identify potential areas for adding error or event handling to improve playbook robustness and reliability.

```sql+postgres
select
  name,
  hosts,
  tasks,
  vars,
  path
from
  ansible_playbook
where
  handlers is null;
```

```sql+sqlite
select
  name,
  hosts,
  tasks,
  vars,
  path
from
  ansible_playbook
where
  handlers is null;
```

### List playbooks that use `root` privilege
Explore which playbooks are utilizing root privileges. This can be beneficial to identify potential security risks and ensure best practices are adhered to.

```sql+postgres
select
  name,
  hosts,
  tasks,
  vars,
  path
from
  ansible_playbook
where
  become
  and (
    become_user is null
    or become_user = 'root'
  );
```

```sql+sqlite
select
  name,
  hosts,
  tasks,
  vars,
  path
from
  ansible_playbook
where
  become = 1
  and (
    become_user is null
    or become_user = 'root'
  );
```