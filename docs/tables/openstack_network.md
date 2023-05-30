# Table: openstack_network

A network contains one or more subnets.

## Examples

### Basic network info

```sql
select
  name,
  description,
  status,
  subnets,
  created_at,
  updated_at,
  project_id,
  shared
from
  openstack_network;
```

### All active networks

```sql
select
  name,
  description,
  subnets,
  created_at,
  updated_at,
  project_id,
  shared
from
  openstack_network
where
  status = 'ACTIVE';
```

### All shared networks

```sql
select
  name,
  description,
  status,
  subnets,
  created_at,
  updated_at,
  project_id,
  shared
from
  openstack_network
where
  shared = true;
```