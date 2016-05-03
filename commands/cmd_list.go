// cmd_list.go

package commands

import (
	"fmt"

	"github.com/grasskode/gitawsm/result"
	"github.com/grasskode/gitawsm/utils"
)

type ListCmd struct {
	pattern string
}

func NewListCmd(pattern string) *ListCmd {
	return &ListCmd{
		pattern: pattern,
	}
}

func (c *ListCmd) help() string {
	return `
	usage: gitawsm list [pattern]
	`
}

func (c *ListCmd) run() *result.Result {
	// read all branches
	list := make(map[string][]string)
	branches := utils.ReadBranches()
	for b, projects := range branches {
		if utils.Matches(b, c.pattern) {
			list[b] = projects
		}
	}

	// format output
	output := "No matching branches found."
	if len(list) > 0 {
		output = ""
		for b, projects := range list {
			output += fmt.Sprintf("\n[%s]", b)
			if len(projects) == 0 {
				output += "\n\t-"
			} else {
				for _, p := range projects {
					output += fmt.Sprintf("\n\t%s", p)
				}
			}
			output += "\n"
		}
	}

	return &result.Result{
		Success: true,
		Message: fmt.Sprintf("\n%s\n", output),
	}
}
