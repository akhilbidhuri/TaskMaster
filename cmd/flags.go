package main

import (
	"flag"

	"github.com/akhilbidhuri/TaskMaster/consts"
)

// this init method assigns all the flags and subcommands
func init() {
	defer consts.PrintOpsonRecover()
	flag.BoolVar(&consts.List, consts.LIST, false, "list the tasks, (to be done), for all add -a")
	flag.BoolVar(&consts.Add, consts.NEW, false, "create new task")
	flag.StringVar(&consts.Remove, consts.DELETE, "", "remove a task, provide the task id")
	flag.StringVar(&consts.MarkDone, consts.MARK_DONE, "", "mark the task as done")
	flag.Parse()
	if consts.List {
		flag.BoolVar(&consts.All, consts.ALL, false, "list all the tasks")
		flag.Parse()
	} else if consts.Add {
		flag.StringVar(&consts.Title, consts.TITLE, "", "title of the task")
		flag.StringVar(&consts.Desc, consts.DESC, "", "description of the task")
		flag.StringVar(&consts.Res, consts.RESOURCES, "", "resources for the task, comma or space sepatated")
		flag.Parse()
		if consts.Title == "" || consts.Desc == "" {
			panic("for add a task title and description are required!")
		}
	} else if consts.Update != "" {
		flag.StringVar(&consts.Title, consts.TITLE, "", "new value of title")
		flag.StringVar(&consts.Desc, consts.DESC, "", "new value of the descriptoin")
		flag.Parse()
		if consts.Title == "" && consts.Desc == "" {
			panic("some value needs to be updated, title and desc empty!")
		}

	} else if consts.MarkDone == "" && consts.Remove == "" {
		panic("\ninvalid operation params!")
	}
}
