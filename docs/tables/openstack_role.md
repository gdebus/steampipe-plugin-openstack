# Table: openstack_role

A role grants authorization to an end-user and dictates the authorization level assigned to a user.

## Examples

### Basic role info

```sql
select
  role_id,
  role_name,
  extra_description
from
  openstack_role;
```