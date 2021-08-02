package main

import (
	"strings"
	"testing"

	"github.com/Ralphbaer/ze-delivery/common"
)

func TestPackages(t *testing.T) {

	var data = `
packages:
- a
- b
`

	packages, err := ListPackages([]byte(data))

	if err != nil {
		t.Error(err)
	}

	if len(packages) != 2 {
		t.Error("Packages lenght must be 2")
	}

}

func TestP(t *testing.T) {

	output := common.CaptureOutput(func() {
		main()
	})

	if len(strings.Split(output, "\n")) != 3 {
		t.Error("Packages lenght must be 2 + 1(newline)")
	}

}
