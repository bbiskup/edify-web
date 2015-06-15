package defs

import (
	"fmt"
	"github.com/bbiskup/edify/edifact/spec/specparser"
	"github.com/bbiskup/edify/edifact/validation"
	"os"
	"path"
	"runtime"
)

var StaticDir string
var BowerDir string
var templateDir string
var SpecDir string
var SpecParser *specparser.FullSpecParser
var Validator *validation.MsgValidator

func init() {
	StaticDir = path.Join(projRootDir(), "static")
	BowerDir = path.Join(projRootDir(), "bower_components")
	templateDir = path.Join(StaticDir, "templates")
	SpecDir = path.Join(projRootDir(), ".edify/downloads/d14b")

	var err error
	Validator, SpecParser, err = validation.GetMsgValidator("14B", SpecDir)
	if err != nil {
		panic(fmt.Sprintf("Unable to create validator: %s", err))
	}
}

// projRootDir determines the root of the edify-web project
// e.g. to allow referencing test data
func projRootDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), "..")
}

func TemplatePaths(templateNames ...string) []string {
	result := make([]string, 0, len(templateNames))
	for _, templateName := range templateNames {
		result = append(result, templateDir+string(os.PathSeparator)+templateName)
	}
	return result
}
