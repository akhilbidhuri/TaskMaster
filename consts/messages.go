package consts

import (
	"fmt"
	"os"
)

type Message string

const (
	ManStr Message = `operations:

	-l list tasks to be done
	-a list all tasks with status
	-new add new task
	-rm remove task`
	Wrong_Params Message = `Invalid params!`
	No_Params    Message = `No params provided!`
)

func PrintOpsonRecover() {
	if r := recover(); r != nil {
		Red.Fprintln(Output, fmt.Sprint(r))
		BlueBold.Fprintln(Output, ManStr)
		os.Exit(0)
	}
}
