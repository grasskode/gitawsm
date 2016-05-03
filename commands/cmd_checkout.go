// cmd_checkout.go

package commands

import (
	"fmt"
	"os"

	"github.com/grasskode/gitawsm/result"
	"github.com/grasskode/gitawsm/utils"
)

type CheckoutCmd struct {
	branch string
}

func NewCheckoutCmd(branch string) *CheckoutCmd {
	return &CheckoutCmd{
		branch: branch,
	}
}

func (c *CheckoutCmd) Help() string {
	return `
	usage: gitawsm checkout <branch_name>
	`
}

func (c *CheckoutCmd) Run() *result.Result {
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

	// check for clean tree in all projects
	unclean := []string{}
	for _, p := range projects {
		if !utils.GitIsCleanWorkingTree(p) {
			unclean = append(unclean, p)
		}
	}
	if len(unclean) > 0 {
		return &result.Result{
			Success: false,
			Message: fmt.Sprintf("There are uncommitted changes in the following project(s). Please stash or commit them before changing branch :\n%v", unclean),
		}
	}

	// checkout branch in all projects
	err_projects := []string{}
	success_projects := []string{}
	for _, p := range projects {
		gerr := utils.GitCreateBranchIfDoesNotExist(p, c.branch)
		if gerr != nil {
			err_projects = append(err_projects, p)
			continue
		}
		gerr = utils.GitCheckoutBranch(p, c.branch)
		if gerr != nil {
			err_projects = append(err_projects, p)
			continue
		}
		success_projects = append(success_projects, p)
	}

	if len(err_projects) > 0 {
		return &result.Result{
			Success: false,
			Message: fmt.Sprintf("Partially checked out branch %q in project(s).\nSuccess : %v\nFailure : %v", c.branch, success_projects, err_projects),
		}
	}

	return &result.Result{
		Success: true,
		Message: fmt.Sprintf("Checked out branch %q in project(s) %v.", c.branch, projects),
	}
}
