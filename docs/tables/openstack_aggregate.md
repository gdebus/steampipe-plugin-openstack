# Table: openstack_aggregate

A host aggregate partitions the hypervisor hosts in an OpenStack cloud. This can be used to partition the hypervisors, e.g., depending on their hardware capabilities.

## Examples

### Basic aggregate info

```sql
select
  availability_zone,
  hosts,
  id,
  name
from
  openstack_aggregate;
```

### All deleted aggregates

```sql
select
  availability_zone,
  hosts,
  hosts,
  metadata,
  name,
  created_at,
  updated_at,
  deleted_at
from
  openstack_aggregate
where
  deleted = true;
```