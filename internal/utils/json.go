package utils

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/antch57/munchies/models"
)

// Marshal the data into JSON
func marshalData(snacks []models.Snack) ([]byte, error) {
	b, err := json.Marshal(snacks)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Unmarshal the data from JSON
func unmarshalData(data []byte) ([]models.Snack, error) {
	var snacks []models.Snack
	err := json.Unmarshal(data, &snacks)
	if err != nil {
		return nil, err
	}
	return snacks, nil
}

func WriteData(snacks []models.Snack) error {
	// Marshal the data into JSON
	b, marshal_err := marshalData(snacks)
	if marshal_err != nil {
		return marshal_err
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
		log.Println("data directory created")
	}

	// Write the JSON data to the file
	write_err := os.WriteFile(dataFilePath, b, 0644)
	if write_err != nil {
		return write_err
	}

	log.Println("snack added")

	return nil
}

// read_data reads the data from the JSON file
func ReadData() ([]models.Snack, error) {
	// Get the path to the data file
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	dataFilePath := filepath.Join(homeDir, ".munchies", "data", "snack.json")

	data, err := os.ReadFile(dataFilePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal the existing data to check if it's empty
	saved_snacks, err := unmarshalData(data)
	if err != nil {
		return nil, err
	}

	log.Println("snack file read")
	return saved_snacks, nil
}
