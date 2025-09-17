package controller

import (
	"github.com/akhilbidhuri/TaskMaster/consts"
	"github.com/akhilbidhuri/TaskMaster/repository"
)

type Controller struct {
	repo *repository.RepositoryI
}

func GetController(repoDep *repository.RepositoryI) *Controller {
	return &Controller{
		repo: repoDep,
	}
}

func (c *Controller) HandleRequest() {
	switch {
	case consts.List:
		if consts.All {

		} else {

		}
	case consts.Add:
	case consts.Update != "":
	case consts.Remove != "":
	case consts.MarkDone != "":
	}
}
