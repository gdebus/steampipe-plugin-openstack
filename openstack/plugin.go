package openstack

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-openstack",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			// identity
			"openstack_user":                  tableOpenstackUser(ctx),
			"openstack_role":                  tableOpenstackRole(ctx),
			"openstack_role_assignment":       tableOpenstackRoleAssignment(ctx),
			"openstack_group":                 tableOpenstackGroup(ctx),
			"openstack_project":               tableOpenstackProject(ctx),
			"openstack_domain":                tableOpenstackDomain(ctx),
			"openstack_applicationcredential": tableOpenstackApplicationcredential(ctx),
			// network
			"openstack_fip":          tableOpenstackFip(ctx),
			"openstack_secgroup":     tableOpenstackSecGroup(ctx),
			"openstack_secgrouprule": tableOpenstackSecGroupRule(ctx),
			"openstack_router":       tableOpenstackRouter(ctx),
			"openstack_network":      tableOpenstackNetwork(ctx),
			"openstack_subnet":       tableOpenstackSubnet(ctx),
			"openstack_port":         tableOpenstackPort(ctx),
			// volumes
			"openstack_volume":      tableOpenstackVolume(ctx),
			"openstack_snapshot":    tableOpenstackSnapshot(ctx),
			"openstack_volume_type": tableOpenstackVolumeType(ctx),
			// compute
			"openstack_server":            tableOpenstackServer(ctx),
			"openstack_compute_image":     tableOpenstackComputeImage(ctx),
			"openstack_keypair":           tableOpenstackKeypair(ctx),
			"openstack_servergroup":       tableOpenstackServergroup(ctx),
			"openstack_aggregate":         tableOpenstackAggregate(ctx),
			"openstack_availability_zone": tableOpenstackAvailabilityZone(ctx),
		},
	}

	return p
}
