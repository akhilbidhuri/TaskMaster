package main

import (
	"log"

	"github.com/akhilbidhuri/TaskMaster/controller"
	repositoryfile "github.com/akhilbidhuri/TaskMaster/repository/repository_file_json"
)

func main() {
	// operations:

	// -l list tasks to be done
	//  -a list all tasks with status
	// -new add new task
	// -rm remove task

	fsRepo := repositoryfile.GetNewFileStore()
	defer fsRepo.F.Close()
	cmdController := controller.GetController(fsRepo)
	log.Println("calling req handler")
	cmdController.HandleRequest()
}
