package repository

import "github.com/akhilbidhuri/TaskMaster/models"

type RepositoryI interface {
	GetToDoTasks() []models.Task
	GetAllTasks() []models.Task
	AddTask(*models.Task) *models.Task
	RemoveTask(id string) bool
	ModifyTask(*models.Task) *models.Task
	MarkTaskDone(id string) bool
	TaskExists(id string) bool
	CleanUp()
}

type Index interface {
	Add(string, int64) error
	Remove(string) error
	Find(string) (int64, error)
}
