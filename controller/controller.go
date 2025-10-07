package controller

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/akhilbidhuri/TaskMaster/consts"
	"github.com/akhilbidhuri/TaskMaster/models"
	"github.com/akhilbidhuri/TaskMaster/repository"
)

type Controller struct {
	repo *repository.RepositoryI
}

func GetController(repoDep repository.RepositoryI) *Controller {
	return &Controller{
		repo: &repoDep,
	}
}

func (c *Controller) HandleRequest() {
	switch strings.Replace(os.Args[1], "-", "", 1) {
	case consts.LIST:
		consts.List.Parse(os.Args[2:])
		var tasks []models.Task
		if consts.All {
			tasks = (*c.repo).GetAllTasks()
		} else {
			tasks = (*c.repo).GetToDoTasks()
		}
		if len(tasks) == 0 {
			log.Fatalln("NO tasks present!")
		}
		for _, task := range tasks {
			fmt.Print(task)
		}
	case consts.NEW:
		consts.Add.Parse(os.Args[2:])
		if consts.Title == "" || consts.Desc == "" {
			panic("title and description are required for creating task!")
		}
		res := getResourceSlice(consts.Res)
		task := &models.Task{
			Title: consts.Title,
			Desc:  consts.Desc,
			Res:   res,
		}
		(*c.repo).AddTask(task)
	case consts.MODIFY:
		flag.Parse()
		id := consts.Update
		if !(*c.repo).TaskExists(id) {
			log.Fatal("Task dosen't exit!")
		}
		if consts.Desc == "" && consts.Title == "" && consts.Res == "" {
			log.Fatal("either title, desc or resources are required as parameter to update command.")
		}
		task := &models.Task{
			ID:    id,
			Title: consts.Title,
			Desc:  consts.Desc,
			Res:   getResourceSlice(consts.Res),
		}
		(*c.repo).ModifyTask(task)
	case consts.MARK_DONE:
		flag.Parse()
		id := consts.MarkDone
		if !(*c.repo).TaskExists(id) {
			log.Fatal("Task dosen't exit!")
		}
		(*c.repo).MarkTaskDone(id)
	case consts.DELETE:
		flag.Parse()
		id := consts.Remove
		if !(*c.repo).TaskExists(id) {
			log.Fatal("Task dosen't exit!")
		}
		(*c.repo).RemoveTask(id)
	case consts.CLEAN:
		(*c.repo).CleanUp()
	default:
	}
}

func getResourceSlice(res string) []string {
	resList := make([]string, 0)
	if consts.Res != "" {
		splitRes := strings.Split(consts.Res, ",")
		for _, r := range splitRes {
			resList = append(resList, strings.TrimSpace(r))
		}
	}
	return resList
}
