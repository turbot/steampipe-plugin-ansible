connection "ansible" {
  plugin = "ansible"

  # The plugin supports parsing both Ansible playbook files as well as inventory files.
  # For example:
  #   - To parse the Ansible playbook files, use `playbook_file_paths` argument to configure it.
  #   - Similarly, to parse the Ansible inventory files, use `inventory_file_paths`.

  # The above paths can be configured with a local directory, a remote Git repository URL, or an S3 bucket URL
  # Wildcard based searches are supported, including recursive searches
  # Local paths are resolved relative to the current working directory (CWD)

  # For example:
  #  - "*.yml" matches all Ansible playbook files in the CWD
  #  - "**/*.yml" matches all Ansible playbook files in the CWD and all sub-directories
  #  - "../*.yml" matches all Ansible playbook files in the CWD's parent directory
  #  - "steampipe*.yml" matches all Ansible playbook files starting with "steampipe" in the CWD
  #  - "/path/to/dir/*.yml" matches all Ansible playbook files in a specific directory
  #  - "/path/to/dir/main.yml" matches a specific file

  # If paths includes "*", all files (including non-Ansible playbook files) in
  # the CWD will be matched, which may cause errors if incompatible file types exist

  # Defaults to CWD
  playbook_file_paths  = [ "*.yml", "*.yaml" ]
  inventory_file_paths = [ "/etc/ansible/hosts", "~/.ansible/hosts" ]
}