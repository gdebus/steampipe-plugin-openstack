# Table: openstack_volume

A volume is a virtual disk that holds data. Volumes can be attached to instances.

## Examples

### Basic volume info

```sql
select
  id,
  name,
  description,
  status,
  size,
  availability_zone,
  updated_at,
  created_at,
  attached_at,
  attachment_id,
  device,
  host_name,
  bootable,
  encrypted
from
  openstack_volume;
```

### All bootable volumes bigger than 10 GB

```sql
select
  id,
  name,
  description,
  status,
  size,
  availability_zone,
  updated_at,
  created_at,
  attached_at,
  attachment_id,
  device,
  host_name,
  bootable,
  encrypted
from
  openstack_volume
where
  size > 10;
```

### All bootable volumes

```sql
select
  id,
  name,
  description,
  status,
  size,
  availability_zone,
  updated_at,
  created_at,
  attached_at,
  attachment_id,
  device,
  host_name,
  bootable,
  encrypted
from
  openstack_volume
where
  bootable = true;
```

### All encrypted volumes

```sql
select
  id,
  name,
  description,
  status,
  size,
  availability_zone,
  updated_at,
  created_at,
  attached_at,
  attachment_id,
  device,
  host_name,
  bootable,
  encrypted
from
  openstack_volume
where
  encrypted = true;
```