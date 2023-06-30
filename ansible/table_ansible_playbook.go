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
		Description: "Ansible playbook",
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

			// Become directives
			{
				Name:        "become",
				Description: "Controls if privilege escalation is used or not on task execution. If true, privilege escalation is activated.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "become_user",
				Description: "User that you 'become' after using privilege escalation.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "become_flags",
				Description: "A string of flag(s) to pass to the privilege escalation program when become is true.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "become_method",
				Description: "Specifies which method of privilege escalation to use (such as sudo or su).",
				Type:        proto.ColumnType_STRING,
			},
			//

			{
				Name:        "check_mode",
				Description: "A boolean that controls if a task is executed in 'check' mode.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "debugger",
				Description: "Enable debugging tasks based on state of the task result. Allowed values are: always, never, on_failed, on_unreachable, on_skipped.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "diff",
				Description: "Toggle to make tasks return ‘diff’ information or not.",
				Type:        proto.ColumnType_BOOL,
			},

			// JSON columns
			{
				Name:        "collections",
				Description: "A section with tasks that are treated as handlers, these won't get executed normally, only when notified after each section of tasks is complete.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tasks",
				Description: "The list of tasks to execute in the play.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "handlers",
				Description: "A section with tasks that are treated as handlers, these won't get executed normally, only when notified after each section of tasks is complete.",
				Type:        proto.ColumnType_JSON,
			},

			// variables
			{
				Name:        "vars",
				Description: "The dictionary/map of variables.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "vars_files",
				Description: "A list of files that contain vars to include in the play.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "vars_prompt",
				Description: "A list of variables to prompt for.",
				Type:        proto.ColumnType_JSON,
			},
			//
			{
				Name:        "path",
				Description: "Path to the file.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

type AnsiblePlaybookInfo struct {
	Become       bool        `cty:"become"`
	BecomeFlags  string      `cty:"become_flags"`
	BecomeMethod string      `cty:"become_method"`
	BecomeUser   string      `cty:"become_user"`
	CheckMode    bool        `cty:"check_mode"`
	Collections  interface{} `cty:"collections"`
	Debugger     string      `cty:"debugger"`
	Diff         bool        `cty:"diff"`
	Handlers     interface{} `cty:"handlers"`
	Hosts        string      `cty:"hosts"`
	Name         string      `cty:"name"`
	Path         string      `cty:"-"`
	Tasks        interface{} `cty:"tasks"`
	Vars         interface{} `cty:"vars"`
	VarsFiles    interface{} `cty:"vars_files"`
	VarsPrompt   interface{} `cty:"vars_prompt"`
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
