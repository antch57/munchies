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

// Command logic
func addSnack(snack *string, count *int, timeInput *string) error {
	// Check if the snack and count are valid
	if *snack == "" || *count == 0 {
		return errors.New("gotta eat a snack to save a snack")
	}

	// If time is provided, validate the provided time format
	if *timeInput != "" {
		parsedTime, err := time.Parse("15:04", *timeInput)
		if err != nil {
			return fmt.Errorf("invalid time format: %v. Please use HH:MM format (e.g., 14:30)", err)
		}
		today := time.Now().Format("2006-01-02")
		*timeInput = fmt.Sprintf("%sT%s:00Z", today, parsedTime.Format("15:04"))

	} else {
		// If no time is provided, use the current time
		*timeInput = time.Now().Format(time.RFC3339)
	}
	// Get the path to the data file
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	dataFilePath := filepath.Join(homeDir, ".munchies", "data", "snack.json")

	// Ensure the directory exists
	dataDir := filepath.Dir(dataFilePath)
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		if mkdirErr := os.MkdirAll(dataDir, 0755); mkdirErr != nil {
			return mkdirErr
		}
	}

	// If it doesn't exist, create it
	if _, err := os.Stat(dataFilePath); err == nil {
		snacks := []models.Snack{
			{
				Snack: *snack,
				Count: *count,
				Time:  *timeInput,
			},
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
			{
				Snack: *snack,
				Count: *count,
				Time:  *timeInput,
			},
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
	timeInput := flagSet.String("time", "", "time of snack (optional, defaults to current time)")

	flagSet.Parse(args)

	flagSet.Usage = func() {
		flag.PrintDefaults()
	}

	err := addSnack(snack, count, timeInput)
	if err != nil {
		return err
	}

	return nil
}
