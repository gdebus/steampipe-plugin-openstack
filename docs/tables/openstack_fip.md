# Table: openstack_fip

A floating IP is an IP address in a public network that can be attached to an instance to make it publicly accessible. Only one floating IP address can be attached to an instance.

## Examples

### Basic floating IP info

```sql
select
  id,
  description,
  floating_ip,
  port_id,
  fixed_ip,
  updated_at,
  created_at,
  project_id,
  status,
  router_id
from
  openstack_fip;
```

### All active floating IPs

```sql
select
  id,
  description,
  floating_ip,
  port_id,
  fixed_ip,
  updated_at,
  created_at,
  project_id,
  status,
  router_id
from
  openstack_fip
where
  status = 'ACTIVE';
```