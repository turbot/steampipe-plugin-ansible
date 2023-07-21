# Table: ansible_task

A task is the smallest unit of action you can automate using an Ansible playbook. Playbooks typically contain a series of tasks that serve a goal, such as to set up a web server, or to deploy an application to remote environments. Ansible executes tasks in the same order they are defined inside a playbook.

The table `ansible_task` lists all the tasks defined inside a playbook.

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
