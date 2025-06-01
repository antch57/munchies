package commands

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/antch57/munchies/internal/utils"
	"github.com/antch57/munchies/models"
)

func listSnack(snack *string, date *string) error {
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

	// If a date or snack is specified, filter the list
	var filtered_snacks []models.Snack

	if *date != "" && *snack != "" {
		for _, s := range saved_snacks {
			// Parse the time from the snack struct
			parsedTime, err := time.Parse(time.RFC3339, s.Time)
			if err != nil {
				return fmt.Errorf("error parsing time for snack %s: %v", s.Snack, err)
			}
			// Format the date to YYYY-MM-DD
			formattedDate := parsedTime.Format("01/02/06")
			if formattedDate == *date && s.Snack == *snack {
				filtered_snacks = append(filtered_snacks, s)
			}
		}
	} else if *date != "" {
		// If a date is specified, filter the list
		if *date != "" {
			for _, s := range saved_snacks {
				// Parse the time from the snack struct
				parsedTime, err := time.Parse(time.RFC3339, s.Time)
				if err != nil {
					return fmt.Errorf("error parsing time for snack %s: %v", s.Snack, err)
				}
				// Format the date to YYYY-MM-DD
				formattedDate := parsedTime.Format("01/02/06")
				if formattedDate == *date {
					filtered_snacks = append(filtered_snacks, s)
				}
			}
		}
	} else if *snack != "" {
		// If a snack is specified, filter the list
		for _, s := range saved_snacks {
			if s.Snack == *snack {
				filtered_snacks = append(filtered_snacks, s)
			}
		}
	}

	// If no snacks match the criteria return.
	if len(filtered_snacks) == 0 {
		fmt.Println(*snack, "not found.")
		return nil
	}

	// Print to the console
	for _, snack := range filtered_snacks {
		// Parse the time from the snack struct
		parsedTime, err := time.Parse(time.RFC3339, snack.Time)
		if err != nil {
			return fmt.Errorf("error parsing time for snack %s: %v", snack.Snack, err)
		}

		// Format the time into MM/DD/YY - HH:MM
		formattedTime := parsedTime.Format("01/02/06 - 15:04")

		fmt.Printf("Snack: %s, Count: %d, Date: %s\n", snack.Snack, snack.Count, formattedTime)
	}

	return nil
}

func ListSnackCmd(args []string) error {
	// Define the command-line arg
	flagSet := flag.NewFlagSet("list", flag.ExitOnError)
	snack := flagSet.String("snack", "", "name of snack to list")
	date := flagSet.String("date", "", "date to filter snacks (MM/DD/YY format)")

	flagSet.Parse(args)

	flagSet.Usage = func() {
		flag.PrintDefaults()
	}

	err := listSnack(snack, date)
	if err != nil {
		return err
	}

	return nil
}
