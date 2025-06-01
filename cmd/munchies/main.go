package main

import (
	"flag"
	"fmt"
	"os"
	"slices"

	"github.com/antch57/munchies/internal/commands"
	"github.com/antch57/munchies/models"
)

func main() {
	flag.Usage = commands.Usage
	flag.Parse()

	// user needs to provide a subcommand
	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	subCmd := flag.Arg(0)
	subCmdArgs := flag.Args()[1:]

	runCommand(subCmd, subCmdArgs)

}

func runCommand(name string, args []string) {
	cmdIdx := slices.IndexFunc(commands.Registry, func(cmd models.Command) bool {
		return cmd.Name == name
	})

	if cmdIdx < 0 {
		fmt.Fprintf(os.Stderr, "command \"%s\" not found\n\n", name)
		flag.Usage()
		os.Exit(1)
	}

	if err := commands.Registry[cmdIdx].Run(args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}
