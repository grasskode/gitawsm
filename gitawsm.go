// gitawsm.go

package main

import (
	"fmt"
	"os"

	"github.com/grasskode/gitawsm/commands"
	"github.com/grasskode/gitawsm/utils"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "help" {
		help := get_help()
		if len(args) > 1 {
			cmd := commands.CreateCommand(args[1], []string{}, false)
			help = cmd.Help()
		}
		fmt.Println(help)
	} else {
		command_args := []string{}
		if len(args) > 1 {
			command_args = args[1:]
		}
		cmd := commands.CreateCommand(args[0], command_args, true)
		result := cmd.Run()
		utils.Print(fmt.Sprintf("Executed command.\nSuccess : %v\nMessage : %v", result.Success, result.Message))
	}
}

func get_help() string {
	return `
	usage: gitawsm [help | COMMAND]
	COMMAND
		branch 		: Create a new gitawsm branch
 		add 		: Add a git project into a gitawsm branch
		list 		: List all gitawsm branches by patterns
 		checkout 	: Checkout a gitawsm branch
		pull		: Pull a gitawsm branch from upstream
		push 		: Push a gitawsm branch to configured remotes
	
	Run "gitawsm help COMMAND" for detailed help of the gitawsm command
`
}
