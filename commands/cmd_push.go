// cmd_push.go

package commands

import (
	"fmt"
	"os"

	"github.com/grasskode/gitawsm/result"
	"github.com/grasskode/gitawsm/utils"
)

type PushCmd struct {
	branch string
}

func NewPushCmd(branch string) *PushCmd {
	return &PushCmd{
		branch: branch,
	}
}

func (c *PushCmd) help() string {
	return `
	usage: gitawsm push <branch_name>
	`
}

func (c *PushCmd) run() *result.Result {
	// check if branch exists
	branches := utils.ReadBranches()
	projects, exists := branches[c.branch]
	if !exists {
		// branch does not exist
		return &result.Result{
			Success: false,
			Message: "Branch does not exist. Use the \"branch\" command to create the branch first.",
		}
	}

	// check if there are associated projects
	if len(projects) == 0 {
		return &result.Result{
			Success: true,
			Message: "No projects found. Use the \"add\" command to add projects to the branch.",
		}
	}

	// check all projects for the branch
	invalid := []string{}
	for _, p := range projects {
		stat, err := os.Stat(fmt.Sprintf("%s/.git", p))
		if !(err == nil && stat.IsDir()) {
			invalid = append(invalid, p)
		}
	}

	// return if there are invalid projects
	if len(invalid) > 0 {
		return &result.Result{
			Success: false,
			Message: fmt.Sprintf("Found invalid git project(s) : %v", invalid),
		}
	}

	// ensure that all projects are on the intended branch
	invalid = []string{}
	for _, p := range projects {
		if utils.GitGetBranch(p) != c.branch {
			invalid = append(invalid, p)
		}
	}
	if len(invalid) > 0 {
		return &result.Result{
			Success: false,
			Message: fmt.Sprintf("The following projects are not on the mentioned branch : %v\nPlease checkout the branch before pushing. Run %q.", invalid, "git checkout "+c.branch),
		}
	}

	// push to remote
	for _, p := range projects {
		if utils.GitPush(p, c.branch) != nil {
			return &result.Result{
				Success: false,
				Message: fmt.Sprintf("Unexpected error while pushing branch %q in project %q.", c.branch, p),
			}
		}
	}

	return &result.Result{
		Success: true,
		Message: fmt.Sprintf("Pushed branch %q.", c.branch),
	}
}
