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

//// TABLE DEFINITION

func tableAnsibleTask(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "ansible_task",
		Description: "Tasks defined in an Ansible playbook",
		List: &plugin.ListConfig{
			ParentHydrate: resolveAnsiblePlaybookFilePaths,
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
				Description: "The name of the playbook where the task is defined.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The name of the playbook.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "any_errors_fatal",
				Description: "Force any un-handled task errors on any host to propagate to all hosts and end the play.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "async",
				Description: "Run a task asynchronously if the C(action) supports this; value is maximum runtime in seconds.",
				Type:        proto.ColumnType_INT,
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
			{
				Name:        "changed_when",
				Description: "Conditional expression that overrides the task's normal 'changed' status.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "check_mode",
				Description: "A boolean that controls if a task is executed in 'check' mode.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "connection",
				Description: "Allows you to change the connection plugin used for tasks to execute on the target.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "debugger",
				Description: "Enable debugging tasks based on state of the task result. Allowed values are: always, never, on_failed, on_unreachable, on_skipped.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "delay",
				Description: "Number of seconds to delay between retries.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "delegate_facts",
				Description: "Boolean that allows you to apply facts to a delegated host instead of inventory_hostname.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "delegate_to",
				Description: "Host to execute task instead of the target (inventory_hostname). Connection vars from the delegated host will also be used for the task.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "diff",
				Description: "Toggle to make tasks return 'diff' information or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "failed_when",
				Description: "Conditional expression that overrides the task's normal 'failed' status.",
				Type:        proto.ColumnType_STRING,
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
				Name:        "loop",
				Description: "Takes a list for the task to iterate over, saving each list element into the item variable (configurable via loop_control)",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "loop_action",
				Description: "Same as action but also implies delegate_to: localhost",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "no_log",
				Description: "Boolean that controls information disclosure.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "poll",
				Description: "Sets the polling interval in seconds for async tasks (default 10s).",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "port",
				Description: "Used to override the default port used in a connection.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "register",
				Description: "Name of variable that will contain task status and module return data.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "remote_user",
				Description: "User used to log into the target via the connection plugin.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "retries",
				Description: "Number of retries before giving up in a until loop.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "run_once",
				Description: "Boolean that will bypass the host loop, forcing the task to attempt to execute on the first host available and afterwards apply any results and facts to all active hosts in the same batch.",
				Type:        proto.ColumnType_BOOL,
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
			{
				Name:        "until",
				Description: "This keyword implies a 'retries loop' that will go on until the condition supplied here is met or we hit the retries limit.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "when",
				Description: "Conditional expression, determines if an iteration of a task is run or not.",
				Type:        proto.ColumnType_STRING,
			},
			// JSON columns
			{
				Name:        "collections",
				Description: "A section with tasks that are treated as handlers, these won't get executed normally, only when notified after each section of tasks is complete.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "loop_control",
				Description: "Several keys here allow you to modify/set loop behaviour in a task.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "module_defaults",
				Description: "Specifies default parameter values for modules.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "notify",
				Description: "A list of handlers to notify when the task returns a 'changed=True' status.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tags",
				Description: "A list of tags applied to the task or included tasks.",
				Type:        proto.ColumnType_JSON,
			},
			// Can't use 'group' as a column since it is a reserved word
			{
				Name:        "task_group",
				Description: "Specifies the group ownership of the task.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Group"),
			},
			// Can't use 'user' as a column since it is a reserved word
			{
				Name:        "task_user",
				Description: "Specifies the the user ownership for the task.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("User"),
			},
			{
				Name:        "vars",
				Description: "The dictionary/map of variables.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

// We include both `cty` and `yaml` struct tags to support multiple input formats.
// - `yaml` tags are used to unmarshal Ansible playbook files written in YAML format.
// - `cty` tags are used to support HCL parsing (e.g., when this struct is used in HCL-based configs).
// This allows the struct to be reused seamlessly across tools that rely on either YAML or HCL formats,
// such as Steampipe plugins that may support both Terraform-like configs and YAML playbooks.
type AnsiblePlaybookTask struct {
	Name  string        `cty:"name" yaml:"name"`
	Tasks []AnsibleTask `cty:"tasks" yaml:"tasks"`
}

type AnsibleTask struct {
	AnyErrorsFatal    string      `cty:"any_errors_fatal" yaml:"any_errors_fatal"`
	Async             int         `cty:"async" yaml:"async"`
	Become            bool        `cty:"become" yaml:"become"`
	BecomeFlags       string      `cty:"become_flags" yaml:"become_flags"`
	BecomeMethod      string      `cty:"become_method" yaml:"become_method"`
	BecomeUser        string      `cty:"become_user" yaml:"become_user"`
	ChangedWhen       string      `cty:"changed_when" yaml:"changed_when"`
	CheckMode         bool        `cty:"check_mode" yaml:"check_mode"`
	Collections       interface{} `cty:"collections" yaml:"collections"`
	Connection        interface{} `cty:"connection" yaml:"connection"`
	Debugger          string      `cty:"debugger" yaml:"debugger"`
	Delay             int         `cty:"delay" yaml:"delay"`
	DelegateFacts     bool        `cty:"delegate_facts" yaml:"delegate_facts"`
	DelegateTo        string      `cty:"delegate_to" yaml:"delegate_to"`
	Diff              bool        `cty:"diff" yaml:"diff"`
	FailedWhen        string      `cty:"failed_when" yaml:"failed_when"`
	Group             interface{} `cty:"group" yaml:"group"`
	IgnoreErrors      bool        `cty:"ignore_errors" yaml:"ignore_errors"`
	IgnoreUnreachable bool        `cty:"ignore_unreachable" yaml:"ignore_unreachable"`
	Loop              string      `cty:"loop" yaml:"loop"`
	LoopAction        string      `cty:"loop_action" yaml:"loop_action"`
	LoopControl       interface{} `cty:"loop_control" yaml:"loop_control"`
	ModuleDefaults    interface{} `cty:"module_defaults" yaml:"module_defaults"`
	Name              string      `cty:"name" yaml:"name"`
	NoLog             bool        `cty:"no_log" yaml:"no_log"`
	Notify            interface{} `cty:"notify" yaml:"notify"`
	Path              string      `cty:"-" yaml:"-"`
	PlaybookName      string      `cty:"-" yaml:"-"`
	Poll              int         `cty:"poll" yaml:"poll"`
	Port              int         `cty:"port" yaml:"port"`
	Register          string      `cty:"register" yaml:"register"`
	RemoteUser        string      `cty:"remote_user" yaml:"remote_user"`
	Retries           int         `cty:"retries" yaml:"retries"`
	RunOnce           bool        `cty:"run_once" yaml:"run_once"`
	Tags              []string    `cty:"tags" yaml:"tags"`
	Throttle          int         `cty:"throttle" yaml:"throttle"`
	Timeout           int         `cty:"timeout" yaml:"timeout"`
	Until             string      `cty:"until" yaml:"until"`
	User              interface{} `cty:"user" yaml:"user"`
	Vars              interface{} `cty:"vars" yaml:"vars"`
	When              string      `cty:"when" yaml:"when"`
}

//// LIST FUNCTION

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
