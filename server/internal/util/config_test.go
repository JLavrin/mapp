package util

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	os.Setenv("STRING_VAR", "example")
	os.Setenv("INT_VAR", "123")
	os.Setenv("BOOL_VAR", "true")

	testCases := []struct {
		key          string
		defaultValue interface{}
		expected     interface{}
	}{
		{"STRING_VAR", "defaultString", "example"},
		{"INT_VAR", 0, 123},
		{"BOOL_VAR", false, true},
		{"NON_EXISTING_VAR", "defaultString", "defaultString"},
	}

	// Run tests
	for _, tc := range testCases {
		result := GetEnv(tc.key, tc.defaultValue)
		if result != tc.expected {
			t.Errorf("GetEnv(%s, %v) = %v, want %v", tc.key, tc.defaultValue, result, tc.expected)
		}
	}
}

func TestLoadEnv(t *testing.T) {
	// Create a mock .env file for testing
	envData := []byte(`KEY1=value1
# Commented line
KEY2=value2`)

	err := os.WriteFile(".env", envData, 0644)
	if err != nil {
		t.Fatalf("Failed to create .env file: %v", err)
	}
	defer os.Remove(".env")

	// Call the function
	LoadEnv()

	// Check if environment variables are set correctly
	tests := map[string]string{
		"KEY1": "value1",
		"KEY2": "value2",
	}

	for key, expected := range tests {
		actual := os.Getenv(key)
		if actual != expected {
			t.Errorf("Expected environment variable %s to be %s, got %s", key, expected, actual)
		}
	}
}
