package main

import (
	"github.com/turbot/steampipe-plugin-ansible/ansible"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: ansible.Plugin})
}
