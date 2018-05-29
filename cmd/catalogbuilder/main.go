package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Files is a map of files.
type Files map[string][]byte

// Glob searches the `/manifests` directory for files matching the pattern and returns them.
func Glob(pattern string) Files {
	matching := map[string][]byte{}
	files, err := filepath.Glob(pattern)
	if err != nil {
		panic(err)
	}

	for _, name := range files {
		bytes, err := ioutil.ReadFile(name)
		if err != nil {
			panic(err)
		}
		matching[name] = bytes
	}

	return matching
}

// Lines is used in the template to split a file's contents by newlines.
// Also removes comment lines and unescapes quotes.
func Lines(bytes []byte) []template.HTML {
	if bytes == nil {
		return []template.HTML{}
	}

	lines := strings.Split(string(bytes), "\n")

	var cleanedLines []template.HTML
	for _, line := range lines {
		if !strings.HasPrefix(line, "#") {
			cleanedLines = append(cleanedLines, template.HTML(line))
		}
	}

	return cleanedLines
}

func parseAndExecute(file string, data interface{}) error {
	tmpl, err := template.New(file).ParseFiles(file)
	if err != nil {
		return err
	}
	err = tmpl.Execute(os.Stdout, data)

	return err
}

// TemplateUtils contains (useful?) functions for templates.
type TemplateUtils struct{}

// Lines wraps the `Lines` function because it can't be passed to a template for some reason.
func (tu TemplateUtils) Lines(bytes []byte) []template.HTML {
	return Lines(bytes)
}

type catalogSourceData struct {
	CatalogNamespace string

	TemplateUtils
}

type configMapData struct {
	CatalogNamespace          string
	CustomResourceDefinitions Files
	ClusterServiceVersions    Files
	Packages                  Files

	TemplateUtils
}

func main() {
	catalogSourceData := catalogSourceData{
		CatalogNamespace: "tectonic-system",

		TemplateUtils: TemplateUtils{},
	}
	configMapData := configMapData{
		CatalogNamespace:          "tectonic-system",
		CustomResourceDefinitions: Glob("manifests/**/**.crd.yaml"),
		ClusterServiceVersions:    Glob("manifests/**/**.clusterserviceversion.yaml"),
		Packages:                  Glob("manifests/**/**.package.yaml"),

		TemplateUtils: TemplateUtils{},
	}

	err := parseAndExecute("operator-manifests.catalogsource.yaml", catalogSourceData)
	if err != nil {
		panic(err)
	}
	fmt.Print("\n---\n")
	err = parseAndExecute("operator-manifests.configmap.yaml", configMapData)
	if err != nil {
		panic(err)
	}
}
