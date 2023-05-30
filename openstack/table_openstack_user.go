package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/users"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableOpenstackUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openstack_group",
		Description: "Table of all users.",
		List: &plugin.ListConfig{
			Hydrate: listUser,
		},
		Columns: []*plugin.Column{
			{Name: "default_project_id", Type: proto.ColumnType_STRING, Description: "DefaultProjectID is the ID of the default project of the user."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description is the description of the user."},
			{Name: "domain_id", Type: proto.ColumnType_STRING, Description: "DomainID is the domain ID the user belongs to."},
			{Name: "enabled", Type: proto.ColumnType_BOOL, Description: "Enabled is whether or not the user is enabled."},
			{Name: "email", Type: proto.ColumnType_STRING, Transform: transform.FromField("Extra.email"), Description: "Email is the email configured for the user."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID is the unique ID of the user."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name is the name of the user."},
			{Name: "lock_password", Type: proto.ColumnType_STRING, Transform: transform.FromField("Options.lock_password"), Description: "Disables the ability for a user to change its password through self-service APIs if set to true."},
			{Name: "password_expires_at", Type: proto.ColumnType_TIMESTAMP, Description: "PasswordExpiresAt is the timestamp when the user's password expires."},
		},
	}
}

func listUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider, err := connect(ctx, d)
	if err != nil {
		logger.Error("openstack_group.listUser", "connection_error", err)
		return nil, err
	}

	// get identity client from provider
	identityClient, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})

	if err != nil {
		logger.Error("openstack_group.listUser", "connection_error", err)
		return nil, err
	}

	// get users
	allPages, err := users.List(identityClient, users.ListOpts{}).AllPages()
	if err != nil {
		logger.Error("openstack_group.listUser", "query_error", err)
		return nil, err
	}

	allUsers, err := users.ExtractUsers(allPages)

	for _, user := range allUsers {
		d.StreamListItem(ctx, user)
	}

	return nil, nil
}
