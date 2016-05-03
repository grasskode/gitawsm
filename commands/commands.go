// commands.go

package commands

import (
	"github.com/grasskode/gitawsm/result"
	"github.com/grasskode/gitawsm/utils"
)

type Command interface {
	Help() string
	Run() *result.Result
}

func CreateCommand(command string, args []string, parseArgs bool) Command {
	switch command {
	case "branch":
		branch := ""
		if parseArgs {
			if len(args) != 1 {
				panic("\"branch\" takes exactly 1 parameter.")
			}
			branch = args[0]
		}
		return NewBranchCmd(branch)
	case "add":
		branch := ""
		paths := []string{}
		if parseArgs {
			if len(args) < 1 {
				panic("\"add\" takes at least 1 parameter.")
			}
			branch = args[0]
			paths = []string{utils.GetWorkingDirectory()}
			if len(args) > 1 {
				paths = args[1:]
			}
		}
		return NewAddCmd(branch, paths)
	case "checkout":
		branch := ""
		if parseArgs {
			if len(args) != 1 {
				panic("\"checkout\" takes exactly 1 parameter.")
			}
			branch = args[0]
		}
		return NewCheckoutCmd(branch)
	case "list":
		pattern := "*"
		if parseArgs {
			if len(args) > 1 {
				panic("\"list\" takes at most 1 parameter.")
			}
			if len(args) == 1 {
				pattern = args[0]
			}
		}
		return NewListCmd(pattern)
	case "push":
		branch := ""
		if parseArgs {
			if len(args) != 1 {
				panic("\"push\" takes exactly 1 parameter.")
			}
			branch = args[0]
		}
		return NewPushCmd(branch)
	case "pull":
		branch := ""
		if parseArgs {
			if len(args) != 1 {
				panic("\"pull\" takes exactly 1 parameter.")
			}
			branch = args[0]
		}
		return NewPullCmd(branch)
	default:
		return &UnknownCmd{
			command: command,
		}
	}
}
