package main

import (
	"flag"
	"os"

	"github.com/akhilbidhuri/TaskMaster/consts"
)

// this init method assigns all the flags and subcommands
func init() {
	defer consts.PrintOpsonRecover()
	consts.List = flag.NewFlagSet(consts.LIST, flag.ExitOnError)
	consts.Add = flag.NewFlagSet(consts.NEW, flag.ExitOnError)
	flag.StringVar(&consts.Remove, consts.DELETE, "", "remove a task, provide the task id")
	flag.StringVar(&consts.MarkDone, consts.MARK_DONE, "", "mark the task as done")
	flag.StringVar(&consts.Update, consts.MODIFY, "", "update a task")
	consts.Add.StringVar(&consts.Title, consts.TITLE, "", "title of the task - for update and create operations")
	consts.Add.StringVar(&consts.Desc, consts.DESC, "", "description of the task - for update and create operations")
	consts.Add.StringVar(&consts.Res, consts.RESOURCES, "", "resources of the task - for uspdate and create operations")
	defer consts.PrintOpsonRecover()

	if len(os.Args) < 2 {
		panic(consts.No_Params)
	}
}
