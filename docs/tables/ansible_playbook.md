# Table: ansible_playbook

Ansible Playbooks offer a repeatable, re-usable, simple configuration management and multi-machine deployment system, one that is well suited to deploying complex applications.

A playbook is composed of one or more `plays` in an ordered list. Each play executes part of the overall goal of the playbook, running one or more tasks. Each task calls an Ansible module.

Playbooks are expressed in YAML format with a minimum of syntax. The table `ansible_playbook` reads all the plays defined in a configured YAML files, and showcase the data in a table format.

## Examples

### Retrieve all playbooks

```sql
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

```sql
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

```sql
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

### List playbooks with no handlers

```sql
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

```sql
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
