package main

import (
	"github.com/turbot/steampipe-plugin-openstack/openstack"
    "github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{PluginFunc: openstack.Plugin})
}