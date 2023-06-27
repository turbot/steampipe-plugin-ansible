package ansible

import (
	"context"

	"github.com/relex/aini"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAnsibleHost(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "ansible_host",
		Description: "Host refers to a specific machine or device that Ansible can manage and operate on",
		List: &plugin.ListConfig{
			Hydrate: listAnsibleHosts,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the host.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Host.Name"),
			},
			{
				Name:        "port",
				Description: "The port that the host allows.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Host.Port"),
			},
			{
				Name:        "vars",
				Description: "A map of variables.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Host.Vars"),
			},
			{
				Name:        "groups",
				Description: "A list of groups where the host is located.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

type AnsibleHostInfo struct {
	Host   *aini.Host
	Groups []string
}

func listAnsibleHosts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// TODO: change the path to use paths given in the config
	path := "test_ansible/inventories/parentChild1"

	data, err := aini.ParseFile(path)
	if err != nil {
		plugin.Logger(ctx).Error("ansible_host.listAnsibleHosts", "read_file_error", err, "path", path)
		return nil, err
	}

	// Stream the data
	for _, host := range data.Hosts {
		var groups []string
		for _, group := range host.Groups {
			groups = append(groups, group.Name)
		}

		d.StreamListItem(ctx, AnsibleHostInfo{
			Host:   host,
			Groups: groups,
		})
	}

	return nil, nil
}
