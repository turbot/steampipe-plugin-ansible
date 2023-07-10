package ansible

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type ansibleConfig struct {
	InventoryFilePaths []string `cty:"inventory_file_paths" steampipe:"watch"`
	PlayBookFilePaths  []string `cty:"playbook_file_paths" steampipe:"watch"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"inventory_file_paths": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{Type: schema.TypeString},
	},
	"playbook_file_paths": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{Type: schema.TypeString},
	},
}

func ConfigInstance() interface{} {
	return &ansibleConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) ansibleConfig {
	if connection == nil || connection.Config == nil {
		return ansibleConfig{}
	}
	config, _ := connection.Config.(ansibleConfig)
	return config
}
