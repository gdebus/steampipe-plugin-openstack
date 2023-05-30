# Table: openstack_server

A server is a VM that runs on a hypervisor. A server is launched from an image.

## Examples

### Basic server info

```sql
select
  name,
  updated,
  created,
  status,
  addresses,
  security_groups,
  attached_volumes
from
  openstack_server;
```

### All shutoff servers

```sql
select
  name,
  updated,
  created,
  status,
  addresses,
  security_groups,
  attached_volumes
from
  openstack_server
where
  status = 'SHUTOFF';
```

### All shelved servers

```sql
select
  name,
  updated,
  created,
  status,
  addresses,
  security_groups,
  attached_volumes
from
  openstack_server
where
  status = 'SHELVED_OFFLOADED';
```