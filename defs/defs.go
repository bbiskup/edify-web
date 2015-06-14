package defs

import (
	"os"
)

const (
	STATIC_DIR = "static"
	BOWER_DIR  = "bower_components"
	//  TPL_DIR    = strings.Join([]string{STATIC_DIR, "templates"}, string(os.PathSeparator))

	TPL_DIR = STATIC_DIR + string(os.PathSeparator) + "templates"

	SPEC_DIR = ".edify/downloads/d14b"
)

func TemplatePaths(templateNames ...string) []string {
	result := make([]string, 0, len(templateNames))
	for _, templateName := range templateNames {
		result = append(result, TPL_DIR+string(os.PathSeparator)+templateName)
	}
	return result
}
