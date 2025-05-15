package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
)

type Snack struct {
	Snack string
	Count int
}

type Command struct {
	Name string
	Help string
	Run  func(args []string) error
}

var commands = []Command{
	{Name: "help", Help: "Print this help", Run: printHelpCmd},
	{Name: "add", Help: "Record a snack that was eaten", Run: addSnackCmd},
	{Name: "list", Help: "list snacks that were eaten", Run: listSnackCmd},
}

func main() {
	var showDebugLog bool
	flag.BoolVar(&showDebugLog, "debug", false, "print debug messages")

	flag.Usage = usage // see below
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
	cmdIdx := slices.IndexFunc(commands, func(cmd Command) bool {
		return cmd.Name == name
	})

	if cmdIdx < 0 {
		fmt.Fprintf(os.Stderr, "command \"%s\" not found\n\n", name)
		flag.Usage()
		os.Exit(1)
	}

	if err := commands[cmdIdx].Run(args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}

// Command logic
func add_snack(snack *string, count *int) error {
	// Check if the snack and count are valid
	if *snack == "" || *count == 0 {
		return errors.New("gotta eat a snack to save a snack")
	}

	// Check if the file exists
	// If it doesn't exist, create it
	if _, err := os.Stat("data/snack.json"); err == nil {
		snacks := []Snack{
			{Snack: *snack, Count: *count},
		}

		// read the existing JSON file
		saved_snacks, err := read_data()
		if err != nil {
			return err
		}

		// Append the new snack to the slice
		list_of_snacks := append(saved_snacks, snacks...)

		// Write the updated JSON back to the file
		write_err := write_data(list_of_snacks)
		if write_err != nil {
			return write_err
		}

	} else {
		snacks := []Snack{
			{Snack: *snack, Count: *count},
		}

		// Write new file.
		err := write_data(snacks)
		if err != nil {
			return err
		}
	}
	return nil
}

func list_snack(snack *string) error {
	// Check if the file exists
	if _, err := os.Stat("data/snack.json"); err != nil {
		return errors.New("snack file does not exist")
	}

	// read the existing JSON file
	saved_snacks, err := read_data()
	if err != nil {
		return err
	}

	// If a snack is specified, filter the list
	if *snack != "" {
		var filtered_snacks []Snack
		for _, s := range saved_snacks {
			if s.Snack == *snack {
				filtered_snacks = append(filtered_snacks, s)
			}
		}
		saved_snacks = filtered_snacks
	}

	if len(saved_snacks) == 0 {
		fmt.Println("No snacks found.")
		return nil
	}

	for _, snack := range saved_snacks {
		fmt.Printf("Snack: %s, Count: %d\n", snack.Snack, snack.Count)
	}

	return nil
}

func usage() {
	intro := "munchies tracks your snacks."
	fmt.Fprintln(os.Stderr, "Usage: munchies <command> [command flags]")
	fmt.Fprintln(os.Stderr, intro)

	fmt.Fprintln(os.Stderr, "\nCommands:")
	for _, cmd := range commands {
		fmt.Fprintf(os.Stderr, "  %-8s %s\n", cmd.Name, cmd.Help)
	}

	fmt.Fprintln(os.Stderr, "\nFlags:")
	flag.PrintDefaults()

	fmt.Fprintln(os.Stderr)
	fmt.Fprintf(os.Stderr, "Run `munchies <command> -h` to get help for a specific command\n\n")
}

// Command functions
func addSnackCmd(args []string) error {
	// Define the command-line arg
	flagSet := flag.NewFlagSet("add", flag.ExitOnError)
	snack := flagSet.String("snack", "", "name to snack")
	count := flagSet.Int("count", 0, "number of snacks eaten")

	flagSet.Usage = func() {
		fmt.Fprintln(os.Stderr, "Add a munchie.")
		fmt.Fprintln(os.Stderr, "\nUsage: munchies add [flags]")
		fmt.Fprintln(os.Stderr, "\nFlags:")
		flag.PrintDefaults()

		fmt.Fprintln(os.Stderr)
	}

	flagSet.Parse(args)

	err := add_snack(snack, count)
	if err != nil {
		return err
	}

	return nil
}

func listSnackCmd(args []string) error {
	// Define the command-line arg
	flagSet := flag.NewFlagSet("list", flag.ExitOnError)
	snack := flagSet.String("snack", "", "name of snack to list")
	flagSet.Usage = func() {
		fmt.Fprintln(os.Stderr, "List munchies eaten.")
		fmt.Fprintln(os.Stderr, "\nUsage: munchies read [flags]")
		fmt.Fprintln(os.Stderr, "\nFlags:")
		flag.PrintDefaults()

		fmt.Fprintln(os.Stderr)
	}

	flagSet.Parse(args)

	err := list_snack(snack)
	if err != nil {
		return err
	}

	return nil
}

func printHelpCmd(_ []string) error {
	flag.Usage()
	return nil
}

// // Helper functions
// Write the data to a file
func write_data(snacks []Snack) error {
	// Marshal the data into JSON
	b, marshal_err := marshal_data(snacks)
	if marshal_err != nil {
		return marshal_err
	}

	// Create the data directory if it doesn't exist
	if _, err := os.Stat("data"); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir("data", 0755)
		if err != nil {
			log.Fatal("error creating data directory:", err)
		}
		log.Println("data directory created")
	}

	// Write the JSON data to the file
	write_err := os.WriteFile("data/snack.json", b, 0644)
	if write_err != nil {
		return write_err
	}

	log.Println("snack added")

	return nil
}

// read_data reads the data from the JSON file
func read_data() ([]Snack, error) {
	data, err := os.ReadFile("data/snack.json")
	if err != nil {
		return nil, err
	}

	// Unmarshal the existing data to check if it's empty
	saved_snacks, err := unmarshal_data(data)
	if err != nil {
		return nil, err
	}

	log.Println("snack file read")
	return saved_snacks, nil
}

// Marshal the data into JSON
func marshal_data(snacks []Snack) ([]byte, error) {
	b, err := json.Marshal(snacks)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Unmarshal the data from JSON
func unmarshal_data(data []byte) ([]Snack, error) {
	var snacks []Snack
	err := json.Unmarshal(data, &snacks)
	if err != nil {
		return nil, err
	}
	return snacks, nil
}
