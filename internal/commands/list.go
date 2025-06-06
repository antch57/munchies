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

func listSnack(snack *string, start *string, end *string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	dataFilePath := filepath.Join(homeDir, ".munchies", "data", "snack.json")
	if _, err := os.Stat(dataFilePath); err != nil {
		return errors.New("snack file does not exist")
	}

	saved_snacks, err := utils.ReadData()
	if err != nil {
		return err
	}

	// Parse date range
	layout := "01/02/06"
	var startDate, endDate time.Time
	now := time.Now()

	if *start == "" {
		*start = now.Format(layout)
	}
	if *end == "" {
		*end = *start
	}
	startDate, err = time.Parse(layout, *start)
	if err != nil {
		return fmt.Errorf("invalid start date: %v", err)
	}
	endDate, err = time.Parse(layout, *end)
	if err != nil {
		return fmt.Errorf("invalid end date: %v", err)
	}

	// Filter snacks by date range and snack name
	var filtered_snacks []models.Snack
	summary := make(map[string]int)
	for _, s := range saved_snacks {
		parsedTime, err := time.Parse(time.RFC3339, s.Time)
		if err != nil {
			continue
		}
		dateOnly := parsedTime.Format(layout)
		dt, _ := time.Parse(layout, dateOnly)
		if (dt.Equal(startDate) || dt.After(startDate)) && (dt.Equal(endDate) || dt.Before(endDate)) {
			if *snack == "" || s.Snack == *snack {
				filtered_snacks = append(filtered_snacks, s)
				summary[s.Snack] += s.Count
			}
		}
	}

	if len(filtered_snacks) == 0 {
		fmt.Println("No snacks found matching the criteria.")
		return nil
	}

	// Print detailed section
	fmt.Printf("Snacks for %s", startDate.Format(layout))
	if !startDate.Equal(endDate) {
		fmt.Printf(" to %s", endDate.Format(layout))
	}
	fmt.Println("\n\nDetailed Log:")
	fmt.Println("────────────────────────────────────────────")
	fmt.Printf("  %-8s | %-10s | %-5s\n", "Time", "Snack", "Count")
	fmt.Println("───────────┼────────────┼───────")
	for _, snack := range filtered_snacks {
		parsedTime, _ := time.Parse(time.RFC3339, snack.Time)
		fmt.Printf("  %-8s | %-10s | %-5d\n", parsedTime.Format("15:04"), snack.Snack, snack.Count)
	}

	// Print summary
	fmt.Println("\nSummary:")
	fmt.Println("─────────────────────────────")
	for snackName, total := range summary {
		bar := ""
		for i := 0; i < total; i++ {
			bar += "█"
		}
		fmt.Printf("%-8s: %-10s %d\n\n", snackName, bar, total)
	}

	return nil
}

func ListSnackCmd(args []string) error {
	flagSet := flag.NewFlagSet("list", flag.ExitOnError)
	snack := flagSet.String("snack", "", "name of snack to list")
	start := flagSet.String("start", "", "start date (MM/DD/YY)")
	end := flagSet.String("end", "", "end date (MM/DD/YY)")

	flagSet.Parse(args)
	flagSet.Usage = func() {
		flag.PrintDefaults()
	}

	return listSnack(snack, start, end)
}
