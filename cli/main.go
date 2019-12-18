package cli

import (
	"fmt"

	"github.com/libidev/requtrap.go/cli/action"
	"github.com/libidev/requtrap.go/cli/config"
	"github.com/libidev/requtrap.go/cli/errors"
)

// Parse - parsing CLI command
func Parse(args []string) {
	var err error
	defer errors.IsError(err)

	if len(args) >= 1 {
		if args[0] == "help" {
			action.Help()
		} else if args[0] == "start" {
			if len(args) == 2 {
				config, err := config.Parse(args[1])
				if err != nil {
					return
				}
				action.Start(config)
			} else {
				fmt.Println("Config file not specified")
			}
		}
	} else {
		action.Help()
	}
}
