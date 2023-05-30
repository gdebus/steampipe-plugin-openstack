# Table: openstack_router

A router is a logical component used to forward packages between different networks. It provides NAT forwarding to allow external access to instances.

## Examples

### Basic router info

```sql
select
  name,
  description,
  status,
  id,
  project_id
from
  openstack_router;
```

### All active routers

```sql
select
  name,
  description,
  status,
  id,
  project_id
from
  openstack_router
where
  status = 'ACTIVE';
```