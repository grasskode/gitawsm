// cmd_unknown.go

package commands

import (
	"fmt"

	"github.com/grasskode/gitawsm/result"
)

type UnknownCmd struct {
	command string
}

func (c *UnknownCmd) get_message() string {
	return fmt.Sprintf("Unknown command %q", c.command)
}

func (c *UnknownCmd) help() string {
	return fmt.Sprintf("\n\t%s\n", c.get_message())
}

func (c *UnknownCmd) run() *result.Result {
	return &result.Result{
		Success: false,
		Message: c.get_message(),
	}
}
