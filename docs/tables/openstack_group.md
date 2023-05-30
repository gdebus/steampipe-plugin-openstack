# Table: openstack_group

A group represents a container for one or multiple users. A group can only be part of one domain.

## Examples

### Basic group info

```sql
select
  group_name,
  description,
  domain_id,
  member_list,
  id
from
  openstack_group;
```

### All groups in default domain

```sql
select
  group_name,
  description,
  domain_id,
  member_list,
  id
from
  openstack_group
where
  domain_id = 'default';
```