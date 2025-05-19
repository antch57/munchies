package commands

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/antch57/munchies/internal/utils"
	"github.com/antch57/munchies/models"
)

func listSnack(snack *string) error {
	// Check if the file exists
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	dataFilePath := filepath.Join(homeDir, ".munchies", "data", "snack.json")

	if _, err := os.Stat(dataFilePath); err != nil {
		return errors.New("snack file does not exist")
	}

	// read the existing JSON file
	saved_snacks, err := utils.ReadData()
	if err != nil {
		return err
	}

	// If a snack is specified, filter the list
	if *snack != "" {
		var filtered_snacks []models.Snack
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

func ListSnackCmd(args []string) error {
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

	err := listSnack(snack)
	if err != nil {
		return err
	}

	return nil
}
