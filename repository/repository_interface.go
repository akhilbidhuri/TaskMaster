package repository

import "github.com/akhilbidhuri/TaskMaster/models"

type RepositoryI interface {
	GetToDoTasks() []models.Task
	GetAllTasks() []models.Task
	AddTask(*models.Task) *models.Task
	RemoveTask(id string) bool
	ModifyTask(*models.Task) *models.Task
}
