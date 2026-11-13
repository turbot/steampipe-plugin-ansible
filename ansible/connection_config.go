package ansible

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type ansibleConfig struct {
	InventoryFilePaths []string `hcl:"inventory_file_paths,optional" steampipe:"watch"`
	PlayBookFilePaths  []string `hcl:"playbook_file_paths,optional" steampipe:"watch"`
}

func ConfigInstance() interface{} {
	return &ansibleConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) ansibleConfig {
	if connection == nil || connection.GetConfig() == nil {
		return ansibleConfig{}
	}
	config, _ := connection.GetConfig().(ansibleConfig)
	return config
}
