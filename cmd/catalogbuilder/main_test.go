package main

import (
	"io/ioutil"
	"path"
	"strings"
	"testing"
)

func TestLinesNil(t *testing.T) {
	lines := Lines(nil)

	if len(lines) != 0 {
		t.Fatalf("Expected empty string slice")
	}
}

func TestLines(t *testing.T) {
	bytes, err := ioutil.ReadFile(path.Join("../../manifests", "etcd-operator", "etcd.package.yaml"))
	if err != nil {
		t.Fatalf("Error reading file")
	}
	lines := Lines(bytes)

	if len(lines) != 5 {
		t.Fatalf("Expected string slice of length %d, received %d", 5, len(lines))
	}
	if lines[len(lines)-1] != "" {
		t.Fatalf("Expected last item in string slice to be empty string, received %s", lines[5])
	}
}

func TestLinesRemovesComments(t *testing.T) {
	bytes, err := ioutil.ReadFile(path.Join("../../manifests", "etcd-operator", "etcd.package.yaml"))
	if err != nil {
		t.Fatalf("Error reading file")
	}
	lines := Lines(bytes)

	for _, line := range lines {
		if strings.HasPrefix(string(line), "#!") {
			t.Fatalf("Expected no comment lines to exist, received %s", line)
		}
	}
}

func TestLinesUnescapesQuotes(t *testing.T) {
	bytes, err := ioutil.ReadFile(path.Join("../../manifests", "etcd-operator", "etcdoperator.clusterserviceversion.yaml"))
	if err != nil {
		t.Fatalf("Error reading file")
	}
	lines := Lines(bytes)

	for _, line := range lines {
		if strings.Contains(string(line), "&#39;") {
			t.Fatalf("Expected no escaped quotes to exist, received %s", line)
		}
	}
}
