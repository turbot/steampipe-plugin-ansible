# Table: ansible_task

Ansible Playbooks offer a repeatable, re-usable, simple configuration management and multi-machine deployment system, one that is well suited to deploying complex applications.

A playbook is composed of one or more `plays` in an ordered list. Each play executes part of the overall goal of the playbook, running one or more tasks. Each task calls an Ansible module.

Playbooks are expressed in YAML format with a minimum of syntax. The table `ansible_playbook` reads all the plays defined in a configured YAML files, and showcase the data in a table format.

## Examples

### Retrieve all tasks in a playbook

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
