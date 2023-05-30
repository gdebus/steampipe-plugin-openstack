# Table: openstack_volume_type

Volume types allow specify attributes for volumes.

## Examples

### Basic volume type info

```sql
select
  id,
  name,
  description,
  extra_specs,
  is_public,
  public_access
from
  openstack_volume_type;
```

### All public volume types

```sql
select
  id,
  name,
  description,
  extra_specs,
  is_public,
  public_access
from
  openstack_volume_type
where
  is_public = true;
```