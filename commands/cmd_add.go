// cmd_add.go

package commands

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/grasskode/gitawsm/result"
	"github.com/grasskode/gitawsm/utils"
)

type AddCmd struct {
	branch string
	paths  []string
}

func NewAddCmd(branch string, paths []string) *AddCmd {
	return &AddCmd{
		branch: strings.Trim(branch, " "),
		paths:  paths,
	}
}

func (c *AddCmd) Help() string {
	return `
	usage: gitawsm add <branch_name> [path...]
	`
}

func (c *AddCmd) Run() *result.Result {
	// check if branch exists
	branches := utils.ReadBranches()
	_, exists := branches[c.branch]
	if !exists {
		// branch does not exist
		return &result.Result{
			Success: false,
			Message: "Branch does not exist. Use the \"branch\" command to create the branch first.",
		}
	}

	// check if all given paths are valid git projects
	invalid := []string{}
	valid := []string{}
	for _, p := range c.paths {
		stat, err := os.Stat(fmt.Sprintf("%s/.git", p))
		if !(err == nil && stat.IsDir()) {
			invalid = append(invalid, p)
		} else {
			if !path.IsAbs(p) {
				p = path.Clean(fmt.Sprintf("%s/%s", utils.GetWorkingDirectory(), p))
			}
			pathIsNew := true
			for _, bp := range branches[c.branch] {
				if bp == p {
					pathIsNew = false
					break
				}
			}
			if pathIsNew {
				valid = append(valid, p)
			}
		}
	}

	// return if there are invalid paths
	if len(invalid) > 0 {
		return &result.Result{
			Success: false,
			Message: fmt.Sprintf("Found invalid git project(s) : %v", invalid),
		}
	}

	// check if there are valid paths to add
	if len(valid) == 0 {
		return &result.Result{
			Success: true,
			Message: "No new paths to add.",
		}
	}

	// add valid path(s) to branch
	branches[c.branch] = append(branches[c.branch], valid...)
	utils.WriteBranches(branches)

	return &result.Result{
		Success: true,
		Message: fmt.Sprintf("Added %d new paths(s) to branch.", len(valid)),
	}
}
