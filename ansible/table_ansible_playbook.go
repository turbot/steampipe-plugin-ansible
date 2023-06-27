package ansible

import (
	"context"
	"fmt"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"gopkg.in/yaml.v3"
)

func tableAnsiblePlaybook(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "ansible_playbook",
		Description: "",
		List: &plugin.ListConfig{
			ParentHydrate: resolveAnsibleConfigPaths,
			Hydrate:       listAnsiblePlaybooks,
			KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The nae of the playbook.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "hosts",
				Description: "A list of groups, hosts or host pattern that translates into a list of hosts that are the play's target.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "become",
				Description: "Controls if privilege escalation is used or not on Task execution.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "become_user",
				Description: "User that you 'become' after using privilege escalation.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "tasks",
				Description: "The list of tasks to execute in the play.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "vars",
				Description: "The dictionary/map of variables.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "handlers",
				Description: "A section with tasks that are treated as handlers, these won't get executed normally, only when notified after each section of tasks is complete.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "path",
				Description: "Path to the file.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

type AnsiblePlaybookInfo struct {
	Path       string      `cty:"-"`
	Name       string      `cty:"name"`
	Hosts      string      `cty:"hosts"`
	Tasks      interface{} `cty:"tasks"`
	Vars       interface{} `cty:"vars"`
	Handlers   interface{} `cty:"handlers"`
	Become     string      `cty:"become"`
	BecomeUser string      `cty:"become_user"`
}

func listAnsiblePlaybooks(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	content, err := os.ReadFile(path)
	if err != nil {
		plugin.Logger(ctx).Error("yml_file.listYMLFileWithPath", "file_error", err, "path", path)
		return nil, fmt.Errorf("failed to read file %s: %v", path, err)
	}

	// Decoding the file content
	var data []AnsiblePlaybookInfo
	err = yaml.Unmarshal(content, &data)
	if err != nil {
		plugin.Logger(ctx).Error("yml_file.listYMLFileWithPath", "parse_error", err, "path", path)
		return nil, fmt.Errorf("failed to unmarshal file content %s: %v", path, err)
	}

	for _, play := range data {
		play.Path = path

		d.StreamListItem(ctx, play)
	}

	return nil, nil
}
