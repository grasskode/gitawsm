// cmd_branch.go

package commands

import (
	"fmt"
	"strings"

	"github.com/grasskode/gitawsm/result"
	"github.com/grasskode/gitawsm/utils"
)

type BranchCmd struct {
	branch string
}

func NewBranchCmd(branch string) *BranchCmd {
	return &BranchCmd{
		branch: strings.Trim(branch, " "),
	}
}

func (c *BranchCmd) Help() string {
	return `
	usage: gitawsm branch <branch_name>
	`
}

func (c *BranchCmd) Run() *result.Result {
	// check if branch is specified
	if c.branch == "" {
		return &result.Result{
			Success: false,
			Message: "No branch name specified.",
		}
	}

	// check if branch is valid git branch
	if !utils.GitIsValidBranchName(c.branch) {
		return &result.Result{
			Success: false,
			Message: "Invalid branch name.",
		}
	}

	// check if branch is whitelisted/blacklisted
	whitelisted := false
	config := utils.ReadConfig()
	for _, wl := range config["whitelist"] {
		if utils.Matches(c.branch, wl) {
			utils.Print("Branch name whitelisted. Continuing.")
			whitelisted = true
			break
		}
	}

	if !whitelisted {
		// if not whitelisted, check if the branch name is blacklisted
		for _, bl := range config["blacklist"] {
			if utils.Matches(c.branch, bl) {
				return &result.Result{
					Success: false,
					Message: "Branch name blacklisted. Check your gitawsm config.",
				}
			}
		}
	}

	// check if branch already exists
	branches := utils.ReadBranches()
	_, exists := branches[c.branch]
	if exists {
		// branch already exists
		return &result.Result{
			Success: false,
			Message: "Branch already exists.",
		}
	}

	// add branch
	branches[c.branch] = []string{}
	err := utils.WriteBranches(branches)
	if err != nil {
		return &result.Result{
			Success: false,
			Message: fmt.Sprintf("Error creating branch : %s", err.Error()),
		}
	}

	return &result.Result{
		Success: true,
		Message: fmt.Sprintf("Created branch %q.", c.branch),
	}
}
