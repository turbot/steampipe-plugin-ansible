package ansible

import (
	"context"

	"github.com/relex/aini"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAnsibleGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "ansible_group",
		Description: "Lists logical grouping of hosts",
		List: &plugin.ListConfig{
			ParentHydrate: resolveAnsibleInventoryFilePaths,
			Hydrate:       listAnsibleGroups,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
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
			{
				Name:        "path",
				Description: "Path to the file.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

type AnsibleGroupInfo struct {
	Children []string
	Group    *aini.Group
	Hosts    []string
	Parents  []string
	Path     string
}

//// LIST FUNCTION

func listAnsibleGroups(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	data, err := aini.ParseFile(path)
	if err != nil {
		plugin.Logger(ctx).Error("ansible_group.listAnsibleGroups", "read_file_error", err, "path", path)
		return nil, err
	}

	// Even if you do not define any groups in your inventory file, Ansible creates two default groups: all and ungrouped. The all group contains every host. The ungrouped group contains all hosts that don't have another group aside from all.

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
			Children: children,
			Group:    group,
			Hosts:    hosts,
			Parents:  parents,
			Path:     path,
		})
	}

	return nil, nil
}
