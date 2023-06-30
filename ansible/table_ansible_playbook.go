package ansible

import (
	"context"
	"fmt"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"gopkg.in/yaml.v3"
)

//// TABLE DEFINITION

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
				Description: "Toggle to make tasks return 'diff' information or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "force_handlers",
				Description: "Will force notified handler execution for hosts even if they failed during the play.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "gather_facts",
				Description: "A boolean that controls if the play will automatically run the 'setup' task to gather facts for the hosts.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "ignore_errors",
				Description: "Boolean that allows you to ignore task failures and continue with play.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "ignore_unreachable",
				Description: "Boolean that allows you to ignore task failures due to an unreachable host and continue with the play.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "max_fail_percentage",
				Description: "It can be used to abort the run after a given percentage of hosts in the current batch has failed.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "no_log",
				Description: "Boolean that controls information disclosure.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "order",
				Description: "Controls the sorting of hosts as they are used for executing the play. Possible values are inventory (default), sorted, reverse_sorted, reverse_inventory and shuffle.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "remote_user",
				Description: "User used to log into the target via the connection plugin.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "run_once",
				Description: "Boolean that will bypass the host loop, forcing the task to attempt to execute on the first host available and afterwards apply any results and facts to all active hosts in the same batch.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "serial",
				Description: "Explicitly define how Ansible batches the execution of the current play on the play's target.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "strategy",
				Description: "Allows you to choose the connection plugin to use for the play.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "tags",
				Description: "Tags applied at the level of play.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "throttle",
				Description: "Limit number of concurrent task runs on task, block and playbook level.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "timeout",
				Description: "Time limit for task to execute in, if exceeded Ansible will interrupt and fail the task.",
				Type:        proto.ColumnType_INT,
			},

			// JSON columns
			{
				Name:        "collections",
				Description: "A section with tasks that are treated as handlers, these won't get executed normally, only when notified after each section of tasks is complete.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "environment",
				Description: "A dictionary that gets converted into environment vars to be provided for the task upon execution.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "gather_subset",
				Description: "Allows you to pass subset options to the fact gathering plugin controlled by gather_facts.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "handlers",
				Description: "A section with tasks that are treated as handlers, these won't get executed normally, only when notified after each section of tasks is complete.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "module_defaults",
				Description: "Specifies default parameter values for modules.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "post_tasks",
				Description: "A list of tasks to execute after the tasks section.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "pre_tasks",
				Description: "A list of tasks to execute before roles.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "roles",
				Description: "The list of roles to be imported into the play.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tasks",
				Description: "The list of tasks to execute in the play.",
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
	Become            bool        `cty:"become"`
	BecomeFlags       string      `cty:"become_flags"`
	BecomeMethod      string      `cty:"become_method"`
	BecomeUser        string      `cty:"become_user"`
	CheckMode         bool        `cty:"check_mode"`
	Collections       interface{} `cty:"collections"`
	Debugger          string      `cty:"debugger"`
	Diff              bool        `cty:"diff"`
	Environment       interface{} `cty:"environment"`
	ForceHandlers     bool        `cty:"force_handlers"`
	GatherFacts       bool        `cty:"gather_facts"`
	GatherSubset      interface{} `cty:"gether_subset"`
	Handlers          interface{} `cty:"handlers"`
	Hosts             string      `cty:"hosts"`
	IgnoreErrors      bool        `cty:"ignore_errors"`
	IgnoreUnreachable bool        `cty:"ignore_unreachable"`
	MaxFailPercentage int         `cty:"max_fail_percentage"`
	ModuleDefaults    interface{} `cty:"module_defaults"`
	Name              string      `cty:"name"`
	NoLog             bool        `cty:"no_log"`
	Order             string      `cty:"order"`
	Path              string      `cty:"-"`
	PostTasks         interface{} `cty:"post_tasks"`
	PreTasks          interface{} `cty:"pre_tasks"`
	RemoteUser        string      `cty:"remote_user"`
	Roles             interface{} `cty:"roles"`
	RunOnce           bool        `cty:"run_once"`
	Serial            int         `cty:"serial"`
	Strategy          string      `cty:"strategy"`
	Tags              string      `cty:"tags"`
	Tasks             interface{} `cty:"tasks"`
	Throttle          int         `cty:"throttle"`
	Timeout           int         `cty:"timeout"`
	Vars              interface{} `cty:"vars"`
	VarsFiles         interface{} `cty:"vars_files"`
	VarsPrompt        interface{} `cty:"vars_prompt"`
}

//// LIST FUNCTION

func listAnsiblePlaybooks(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// The path comes from a parent hydrate, defaulting to the config paths or
	// available by the optional key column
	path := h.Item.(filePath).Path

	content, err := os.ReadFile(path)
	if err != nil {
		plugin.Logger(ctx).Error("ansible_playbook.listAnsiblePlaybooks", "file_error", err, "path", path)
		return nil, fmt.Errorf("failed to read file %s: %v", path, err)
	}

	// Decoding the file content
	var data []AnsiblePlaybookInfo
	err = yaml.Unmarshal(content, &data)
	if err != nil {
		plugin.Logger(ctx).Error("ansible_playbook.listAnsiblePlaybooks", "parse_error", err, "path", path)
		return nil, fmt.Errorf("failed to unmarshal file content %s: %v", path, err)
	}

	for _, play := range data {
		play.Path = path

		d.StreamListItem(ctx, play)
	}

	return nil, nil
}
