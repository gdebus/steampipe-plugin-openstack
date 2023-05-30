# Table: openstack_compute_image

An image contains a virtual disk that holds a bootable operating system.

## Examples

### Basic image info

```sql
select
  name,
  status,
  id,
  created,
  min_disk,
  min_ram
from
  openstack_compute_image;
```

### All active images

```sql
select
  name,
  id,
  created,
  min_disk,
  min_ram
from
  openstack_compute_image
where
  status = 'ACTIVE';
```

### All images not yet built

```sql
select
  name,
  id,
  created,
  min_disk,
  min_ram
from
  openstack_compute_image
where
  progress < 100;
```