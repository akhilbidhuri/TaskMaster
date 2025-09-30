package repository

import "github.com/akhilbidhuri/TaskMaster/models"

type RepositoryI interface {
	GetToDoTasks() []models.Task
	GetAllTasks() []models.Task
	AddTask(*models.Task) *models.Task
	RemoveTask(id string) bool
	ModifyTask(*models.Task) *models.Task
	TaskExists(id string) bool
}

type Index interface {
	Add(string, int64) error
	Remove(string) error
	Find(string) (int64, error)
}
