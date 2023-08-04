package ansible

import (
	"context"
	"errors"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	filehelpers "github.com/turbot/go-kit/files"
)

type filePath struct {
	Path string
}

func resolveAnsiblePlaybookFilePaths(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	// #1 - Path via qual

	// If the path was requested through qualifier then match it exactly. Globs
	// are not supported in this context since the output value for the column
	// will never match the requested value.
	quals := d.EqualsQuals
	if quals["path"] != nil {
		d.StreamListItem(ctx, filePath{Path: quals["path"].GetStringValue()})
		return nil, nil
	}

	// #2 - paths in config

	// Fail if no paths are specified
	ansibleConfig := GetConfig(d.Connection)
	if ansibleConfig.PlayBookFilePaths == nil {
		return nil, errors.New("playbook_file_paths must be configured")
	}

	// Gather file path matches for the glob
	var matches []string
	paths := ansibleConfig.PlayBookFilePaths
	for _, i := range paths {

		// List the files in the given source directory
		files, err := d.GetSourceFiles(i)
		if err != nil {
			return nil, err
		}
		matches = append(matches, files...)
	}

	// Sanitize the matches to ignore the directories
	for _, i := range matches {

		// Ignore directories
		if filehelpers.DirectoryExists(i) {
			continue
		}
		d.StreamListItem(ctx, filePath{Path: i})
	}

	return nil, nil
}

func resolveAnsibleInventoryFilePaths(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	// #1 - Path via qual

	// If the path was requested through qualifier then match it exactly. Globs
	// are not supported in this context since the output value for the column
	// will never match the requested value.
	quals := d.EqualsQuals
	if quals["path"] != nil {
		d.StreamListItem(ctx, filePath{Path: quals["path"].GetStringValue()})
		return nil, nil
	}

	// #2 - paths in config

	// Fail if no paths are specified
	ansibleConfig := GetConfig(d.Connection)
	if ansibleConfig.InventoryFilePaths == nil {
		return nil, errors.New("inventory_file_paths must be configured")
	}

	// Gather file path matches for the glob
	var matches []string
	paths := ansibleConfig.InventoryFilePaths
	for _, i := range paths {

		// List the files in the given source directory
		files, err := d.GetSourceFiles(i)
		if err != nil {
			return nil, err
		}
		matches = append(matches, files...)
	}

	// Sanitize the matches to ignore the directories
	for _, i := range matches {

		// Ignore directories
		if filehelpers.DirectoryExists(i) {
			continue
		}
		d.StreamListItem(ctx, filePath{Path: i})
	}

	return nil, nil
}
