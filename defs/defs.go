package defs

import (
	"fmt"
	"github.com/bbiskup/edify/edifact/validation"
	"os"
)

var Validator *validation.MsgValidator

const (
	STATIC_DIR = "static"
	BOWER_DIR  = "bower_components"
	//  TPL_DIR    = strings.Join([]string{STATIC_DIR, "templates"}, string(os.PathSeparator))

	TPL_DIR = STATIC_DIR + string(os.PathSeparator) + "templates"

	SPEC_DIR = ".edify/downloads/d14b"
)

func init() {
	var err error
	Validator, err = validation.GetMsgValidator("14B", SPEC_DIR)
	if err != nil {
		panic(fmt.Sprintf("Unable to create validator: %s", err))
	}
}

func TemplatePaths(templateNames ...string) []string {
	result := make([]string, 0, len(templateNames))
	for _, templateName := range templateNames {
		result = append(result, TPL_DIR+string(os.PathSeparator)+templateName)
	}
	return result
}
