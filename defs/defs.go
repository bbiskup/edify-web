package defs

import (
	"fmt"
	"github.com/bbiskup/edify/edifact/spec/specparser"
	"github.com/bbiskup/edify/edifact/validation"
	"os"
)

const (
	STATIC_DIR = "static"
	BOWER_DIR  = "bower_components"
	//  TPL_DIR    = strings.Join([]string{STATIC_DIR, "templates"}, string(os.PathSeparator))

	TPL_DIR = STATIC_DIR + string(os.PathSeparator) + "templates"

	SPEC_DIR = ".edify/downloads/d14b"
)

var SpecParser *specparser.FullSpecParser
var Validator *validation.MsgValidator

func init() {
	var err error
	Validator, SpecParser, err = validation.GetMsgValidator("14B", SPEC_DIR)
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
