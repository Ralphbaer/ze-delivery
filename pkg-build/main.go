package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

const defaultFilePath = "./pkg-config.yml"

// T represents the root of yml file.
type T struct {
	Packages []string
}

// ListPackages list all packages to be build.
func ListPackages(data []byte) ([]string, error) {

	t := T{}

	err := yaml.Unmarshal(data, &t)

	if err != nil {
		return nil, err
	}

	return t.Packages, nil
}

func main() {

	filePath := defaultFilePath

	if len(os.Args) == 2 {
		filePath = os.Args[1]
	}

	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	packages, err := ListPackages(file)

	if err != nil {
		panic(err)
	}

	fmt.Println(strings.Join(packages, "\n"))
}
