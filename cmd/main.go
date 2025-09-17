package main

import (
	"os"

	"github.com/akhilbidhuri/TaskMaster/consts"
	"github.com/akhilbidhuri/TaskMaster/controller"
	repositoryfile "github.com/akhilbidhuri/TaskMaster/repository/repository_file_json"
)

func main() {
	// operations:

	// -l list tasks to be done
	//  -a list all tasks with status
	// -new add new task
	// -rm remove task
	defer consts.PrintOpsonRecover()

	if len(os.Args) < 2 {
		panic(consts.No_Params)
	}
	fsRepo := repositoryfile.GetNewFileStore()
	cmdController := controller.GetController(&fsRepo)
	cmdController.HandleRequest()
}
