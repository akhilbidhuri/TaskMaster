package main

import (
	"fmt"

	"github.com/mattn/go-colorable"

	"github.com/dimiro1/banner"
)

func init() {
	templ := `{{ .AnsiColor.BrightGreen }}{{ .Title "TaskMaster" "" 8 }}
   	   {{ .AnsiColor.BrightRed }}This is the task master where you register your tasks to be done{{.AnsiColor.Green}}`

	banner.InitString(colorable.NewColorableStdout(), true, true, templ)
	fmt.Println()
}
