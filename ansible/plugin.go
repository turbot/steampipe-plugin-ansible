package ansible

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-ansible"

// Plugin creates this (ansible) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel().NullIfZero(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		TableMap: map[string]*plugin.Table{
			"ansible_group":    tableAnsibleGroup(ctx),
			"ansible_host":     tableAnsibleHost(ctx),
			"ansible_playbook": tableAnsiblePlaybook(ctx),
			"ansible_task":     tableAnsibleTask(ctx),
		},
	}

	return p
}
