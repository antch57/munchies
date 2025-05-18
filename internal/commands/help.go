package commands

import (
	"flag"
	"fmt"
	"os"
)

func Usage() {
	intro := "munchies tracks your snacks."
	fmt.Fprintln(os.Stderr, "Usage: munchies <command> [command flags]")
	fmt.Fprintln(os.Stderr, intro)

	fmt.Fprintln(os.Stderr, "\nCommands:")
	for _, cmd := range Registry {
		fmt.Fprintf(os.Stderr, "  %-8s %s\n", cmd.Name, cmd.Help)
	}

	fmt.Fprintln(os.Stderr, "\nFlags:")
	flag.PrintDefaults()

	fmt.Fprintln(os.Stderr)
	fmt.Fprintf(os.Stderr, "Run `munchies <command> -h` to get help for a specific command\n\n")
}

func PrintHelpCmd(_ []string) error {
	flag.Usage()
	return nil
}
