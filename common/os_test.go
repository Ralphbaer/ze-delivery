package common

import (
	"os"
	"testing"
)

func TestGetenvBool(t *testing.T) {

	if err := os.Setenv("KEY_TRUE", "true"); err != nil {
		t.Error(err)
	}
	if err := os.Setenv("KEY_FALSE", "false"); err != nil {
		t.Error(err)
	}

	if err := os.Setenv("INVALID_KEY", "invalid"); err != nil {
		t.Error(err)
	}

	validTrue := GetenvBoolOrDefault("KEY_TRUE", false)
	validFalse := GetenvBoolOrDefault("KEY_FALSE", true)
	InvalidKey := GetenvBoolOrDefault("INVALID_KEY", true)

	if validTrue != true {
		t.Errorf("got: %v want: %v", validTrue, true)
	}

	if validFalse != false {
		t.Errorf("got: %v want: %v", validFalse, false)
	}

	if InvalidKey != true {
		t.Errorf("got: %v want: %v", InvalidKey, true)
	}
}
