package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	output := os.Args[1]
	files := os.Args[2:]
	manifest := map[string]string{}

	for _, file := range files {
		hash, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		serviceName := filepath.Base(filepath.Dir(file)) + "-service"
		manifest[serviceName] = string(hash)
	}
	j, err := json.MarshalIndent(manifest, "", "  ")

	if err != nil {
		log.Fatalf("Error generating manifest: %s\n", err.Error())
	}

	err = ioutil.WriteFile(output, j, 0644)
	if err != nil {
		log.Fatalf("Error writing manifest to file: %s\n", err.Error())
	}
}
