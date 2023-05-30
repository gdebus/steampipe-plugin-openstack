---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/openstack.svg"
brand_color: "#ed1944"
display_name: OpenStack
name: openstack
description: Steampipe plugin to query OpenStack deployments.
og_description: Query OpenStack with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/turbot/openstack-graphic.png"
---

# OpenStack + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

[OpenStack](https://www.openstack.org/) is the most widely deployed open source cloud software in the world.

Example query:

```sql
select
  name,
  description,
  email,
  enabled
from
  openstack_user;
```

```
+-------------------+-----------------------------+---------+
| name              | email                       | enabled |
+-------------------+-----------------------------+---------+
| demo              | demo@example.com            | true    |
| admin             | admin@testproject.com       | true    |
| reader            | reader@testproject.com      | true    |
+-------------------+-----------------------------+---------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/openstack/tables)**

## Get started

### Install

Download and install the latest OpenStack plugin:

```bash
steampipe plugin install openstack
```

### Configuration

Installing the latest openstack plugin will create a config file (`~/.steampipe/config/openstack.spc`) with a single connection named `openstack`:

```hcl
connection "openstack" {
    plugin    = "local/openstack"

    # IdentityEndpoint specifies the HTTP endpoint that is required to work with
    # the Identity API of the appropriate version. While it's ultimately needed by
    # all of the identity services, it will often be populated by a provider-level
    # function.
    # Can also be set with the environment variable "OS_AUTH_URL"
    # identity_endpoint = "http://example.com/identity/v3"

    # Username is required if using Identity V2 API. Consult with your provider's
    # control panel to discover your account's username. In Identity V3, either
    # UserID or a combination of Username and DomainID or DomainName are needed.
    # Can also be set with the environment variable "OS_USERNAME"
    # username = "admin"
    # Can also be set with the environment variable "OS_USER_ID"
    # user_id = "d8e8fca2dc0f896fd7cb4cb0031ba249"

    # Password specifies the password of the user
    # Can also be set with the environment variable "OS_PASSWORD"
    # password = "changeme"

    # Passcode is used in TOTP authentication method
    # Can also be set with the environment variable "OS_PASSCODE"
    # passcode = "123456"

    # At most one of DomainID and DomainName must be provided if using Username
    # with Identity V3. Otherwise, either are optional.
    # Can also be set with the environment variable "OS_DOMAIN_ID"
    # domain_id = "default"
    # Can also be set with the environment variable "OS_DOMAIN_NAME"
    # domain_name = "Default"

    # The TenantID and TenantName fields are optional for the Identity V2 API.
    # The same fields are known as project_id and project_name in the Identity
    # V3 API, but are collected as TenantID and TenantName here in both cases.
    # Some providers allow you to specify a TenantName instead of the TenantId.
    # Some require both. Your provider's authentication policies will determine
    # how these fields influence authentication.
    # If DomainID or DomainName are provided, they will also apply to TenantName.
    # It is not currently possible to authenticate with Username and a Domain
    # and scope to a Project in a different Domain by using TenantName. To
    # accomplish that, the ProjectID will need to be provided as the TenantID
    # option.
    # Can also be set with the environment variable "OS_PROJECT_ID"
    # project_id = "3e666015f769bf30cda73a1a1e9b794a"
    # Can also be set with the environment variable "OS_PROJECT_NAME"
    # project_name = "my_project"

    # AllowReauth should be set to true if you grant permission for Gophercloud to
    # cache your credentials in memory, and to allow Gophercloud to attempt to
    # re-authenticate automatically if/when your token expires.  If you set it to
    # false, it will not cache these settings, but re-authentication will not be
    # possible.  This setting defaults to false.
    # allow_reauth = false
}
```

Environment variables are also available as an alternate configuration method:

- `OS_AUTH_URL`
- `OS_USERNAME`
- `OS_USER_ID`
- `OS_PASSWORD`
- `OS_PASSCODE`
- `OS_DOMAIN_ID`
- `OS_DOMAIN_NAME`
- `OS_PROJECT_ID`
- `OS_PROJECT_NAME`

## Get involved

* Open source: https://github.com/turbot/steampipe-plugin-openstack
* Community: [Slack Channel](https://steampipe.io/community/join)