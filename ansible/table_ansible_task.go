package ansible

import (
	"context"
	"fmt"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"gopkg.in/yaml.v3"
)

func tableAnsibleTask(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "ansible_task",
		Description: "",
		List: &plugin.ListConfig{
			ParentHydrate: resolveAnsibleConfigPaths,
			Hydrate:       listAnsibleTasks,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "path",
				Description: "Path to the file.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "playbook_name",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The nae of the playbook.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "_group",
				Description: "",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Group"),
			},
			{
				Name:        "_user",
				Description: "",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("User"),
			},
			{
				Name:        "notify",
				Description: "",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "vars",
				Description: "The dictionary/map of variables.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

type AnsiblePlaybookTask struct {
	Name  string        `cty:"name"`
	Tasks []AnsibleTask `cty:"tasks"`
}

type AnsibleTask struct {
	Path         string      `cty:"-"`
	PlaybookName string      `cty:"-"`
	Name         string      `cty:"name"`
	Group        interface{} `cty:"group"`
	User         interface{} `cty:"user"`
	Notify       interface{} `cty:"notify"`
	Vars         interface{} `cty:"vars"`
}

func listAnsibleTasks(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	content, err := os.ReadFile(path)
	if err != nil {
		plugin.Logger(ctx).Error("ansible_task.listAnsibleTasks", "file_error", err, "path", path)
		return nil, fmt.Errorf("failed to read file %s: %v", path, err)
	}

	// Decoding the file content
	var data []AnsiblePlaybookTask
	err = yaml.Unmarshal(content, &data)
	if err != nil {
		plugin.Logger(ctx).Error("ansible_task.listAnsibleTasks", "parse_error", err, "path", path)
		return nil, fmt.Errorf("failed to unmarshal file content %s: %v", path, err)
	}

	for _, play := range data {
		for _, task := range play.Tasks {
			task.Path = path
			task.PlaybookName = play.Name

			d.StreamListItem(ctx, task)
		}
	}

	return nil, nil
}
