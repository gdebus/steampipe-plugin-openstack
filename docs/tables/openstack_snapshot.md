# Table: openstack_snapshot

A snapshot is the backup of a volume.

## Examples

### Basic snapshot info

```sql
select
  snapshot_id,
  snapshot_name,
  description,
  created_at,
  updated_at,
  volume_id,
  status,
  size
from
  openstack_snapshot;
```

### All snapshots bigger than 1 GB

```sql
select
  snapshot_id,
  snapshot_name,
  description,
  created_at,
  updated_at,
  volume_id,
  status,
  size
from
  openstack_snapshot
where
  size > 1;
```

### All available snapshots

```sql
select
  snapshot_id,
  snapshot_name,
  description,
  created_at,
  updated_at,
  volume_id,
  status,
  size
from
  openstack_snapshot
where
  status = 'available';
```
