package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/bazelbuild/rules_go/go/tools/bazel"
)

var cli = flag.String("cli", "", "The CLI binary")

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestCLIManifest(t *testing.T) {
	t.Parallel()

	path, err := bazel.Runfile(*cli)
	if err != nil {
		t.Fatalf("Could not find runfile %s: %q", *cli, err)
	}

	if _, err = os.Stat(path); os.IsNotExist(err) {
		t.Fatalf("Missing binary %v", path)
	}
	file, err := filepath.EvalSymlinks(path)
	if err != nil {
		t.Fatalf("Invalid filename %v", path)
	}

	args := []string{"manifest.json"}
	args = append(args, os.Args[2:]...)
	cmd := exec.Command(file, args...)
	cmd.Stderr = os.Stderr
	res, err := cmd.Output()
	if err != nil {
		t.Fatalf("failed running '%v': %v", path, err)
	}
	output := strings.TrimSpace(string(res))

	if output != "" {
		t.Error("Expected", "", "got", output)
	}

	// load manifest and check contents
	data, err := ioutil.ReadFile("manifest.json")
	if err != nil {
		t.Error("Error loading manifest.json output file: ", err)
	}
	var manifest map[string]string
	err = json.Unmarshal(data, &manifest)
	if err != nil {
		t.Error("Error Unmarshaling manifest.json output file: ", err)
	}
	if manifest["test1-service"] != "1234" {
		t.Error("Expected", "1234", "got", manifest["test1-service"])
	}

	if manifest["test2-service"] != "1234" {
		t.Error("Expected", "1234", "got", manifest["test2-service"])
	}
}
