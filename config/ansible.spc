connection "ansible" {
  plugin = "ansible"

  # Paths is a list of locations to search for Ansible playbook files
  # Paths can be configured with a local directory, a remote Git repository URL, or an S3 bucket URL
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
  paths = [ "*.yml", ""*.yaml"" ]
}