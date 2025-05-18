package utils

import (
	"encoding/json"
	"errors"
	"log"
	"os"

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
func ReadData() ([]models.Snack, error) {
	data, err := os.ReadFile("data/snack.json")
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
