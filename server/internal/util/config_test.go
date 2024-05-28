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
