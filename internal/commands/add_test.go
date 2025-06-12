package commands

import (
	"testing"
)

// Test for empty snack
func TestAddSnack_EmptySnackReturnsError(t *testing.T) {
	snack := ""
	count := 1
	timeInput := ""
	err := addSnack(&snack, &count, &timeInput)
	if err == nil || err.Error() != "gotta eat a snack to save a snack" {
		t.Errorf("expected error for empty snack, got %v", err)
	}
}

// Test for zero count
func TestAddSnack_ZeroCountReturnsError(t *testing.T) {
	snack := "chips"
	count := 0
	timeInput := ""
	err := addSnack(&snack, &count, &timeInput)
	if err == nil || err.Error() != "gotta eat a snack to save a snack" {
		t.Errorf("expected error for zero count, got %v", err)
	}
}

// Test for invalid time format
func TestAddSnack_InvalidTimeFormatReturnsError(t *testing.T) {
	snack := "chips"
	count := 2
	timeInput := "badtime"
	err := addSnack(&snack, &count, &timeInput)
	if err == nil || err.Error() == "" {
		t.Errorf("expected invalid time format error, got %v", err)
	}
}
