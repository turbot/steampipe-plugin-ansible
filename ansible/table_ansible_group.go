package ansible

import (
	"context"

	"github.com/relex/aini"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAnsibleGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "ansible_group",
		Description: "Lists logical grouping of hosts",
		List: &plugin.ListConfig{
			Hydrate: listAnsibleGroups,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Group.Name"),
			},
			{
				Name:        "hosts",
				Description: "A list of hosts listed under the group.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "parents",
				Description: "A list of parent groups under which this group is located.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "children",
				Description: "A list of child groups.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "vars",
				Description: "A map of group variables.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Group.Vars"),
			},
		},
	}
}

type AnsibleGroupInfo struct {
	Group    *aini.Group
	Parents  []string
	Children []string
	Hosts    []string
}

func listAnsibleGroups(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// TODO: change the path to use paths given in the config
	path := "test_ansible/inventories/parentChild1"

	data, err := aini.ParseFile(path)
	if err != nil {
		plugin.Logger(ctx).Error("ansible_group.listAnsibleGroups", "read_file_error", err, "path", path)
		return nil, err
	}

	// Even if you do not define any groups in your inventory file, Ansible creates two default groups: all and ungrouped . The all group contains every host. The ungrouped group contains all hosts that don't have another group aside from all .

	// Stream the data
	for _, group := range data.Groups {
		var hosts, parents, children []string

		for _, host := range group.Hosts {
			hosts = append(hosts, host.Name)
		}

		for _, parent := range group.Parents {
			parents = append(parents, parent.Name)
		}

		for _, child := range group.Children {
			children = append(children, child.Name)
		}

		d.StreamListItem(ctx, AnsibleGroupInfo{
			Group:    group,
			Parents:  parents,
			Children: children,
			Hosts:    hosts,
		})
	}

	return nil, nil
}
