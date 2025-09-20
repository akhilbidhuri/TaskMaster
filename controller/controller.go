package controller

import (
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
	log.Println("came here")
	fmt.Println("args1: ", os.Args[1])
	switch strings.Replace(os.Args[1], "-", "", 1) {
	case consts.LIST:
		consts.List.Parse(os.Args[2:])
	case consts.NEW:
		consts.Add.Parse(os.Args[2:])
		if consts.Title == "" || consts.Desc == "" {
			panic("title and description are required for creating task!")
		}
		res := make([]string, 0, 5)
		if consts.Res != "" {
			splitRes := strings.Split(consts.Res, ",")
			for _, r := range splitRes {
				res = append(res, strings.TrimSpace(r))
			}
		}
		task := &models.Task{
			Title: consts.Title,
			Desc:  consts.Desc,
			Res:   res,
		}
		(*c.repo).AddTask(task)
	default:
	}
}
