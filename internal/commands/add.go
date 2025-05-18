package commands

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/antch57/munchies/internal/utils"
	"github.com/antch57/munchies/models"
)

// Command logic
func addSnack(snack *string, count *int) error {
	// Check if the snack and count are valid
	if *snack == "" || *count == 0 {
		return errors.New("gotta eat a snack to save a snack")
	}

	// Check if the file exists
	// If it doesn't exist, create it
	if _, err := os.Stat("data/snack.json"); err == nil {
		snacks := []models.Snack{
			{Snack: *snack, Count: *count},
		}

		// read the existing JSON file
		saved_snacks, err := utils.ReadData()
		if err != nil {
			return err
		}

		// Append the new snack to the slice
		list_of_snacks := append(saved_snacks, snacks...)

		// Write the updated JSON back to the file
		write_err := utils.WriteData(list_of_snacks)
		if write_err != nil {
			return write_err
		}

	} else {
		snacks := []models.Snack{
			{Snack: *snack, Count: *count},
		}

		// Write new file.
		err := utils.WriteData(snacks)
		if err != nil {
			return err
		}
	}
	return nil
}

func AddSnackCmd(args []string) error {
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

	err := addSnack(snack, count)
	if err != nil {
		return err
	}

	return nil
}
