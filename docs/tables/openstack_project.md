# Table: openstack_project

A project (tenant) is a container for cloud resources. OpenStack software components split their resources into multiple projects. A project can be part of only one domain.

## Examples

### Basic project info

```sql
select
  name,
  description,
  is_domain,
  domain_id,
  enabled,
  id
from
  openstack_project;
```

### All active projects

```sql
select
  name,
  description,
  is_domain,
  domain_id,
  enabled,
  id
from
  openstack_project
where
  enabled = true;
```

### All projects that are a domain

```
select
  name,
  description,
  is_domain,
  domain_id,
  enabled,
  id
from
  openstack_project
where
  is_domain = true;
```

### All projects of the default domain

```
select
  name,
  description,
  is_domain,
  domain_id,
  enabled,
  id
from
  openstack_project
where
  domain_id = 'default';
```