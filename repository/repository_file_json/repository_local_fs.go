package repositoryfile

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/akhilbidhuri/TaskMaster/models"
	"github.com/akhilbidhuri/TaskMaster/repository"
)

const storePath = "tm_data.json"

var basePath = filepath.Clean("../store")

var fileStorePath = filepath.Join(basePath, storePath)

func exitOnIssue(msg interface{}) {
	log.Fatal(msg)
	os.Exit(0)
}

type FileStore struct {
	f *os.File
}

// if dir and file store not present create
func init() {
	log.Println("init repo")
	entries, err := os.ReadDir(filepath.Clean("../"))
	if err != nil {
		exitOnIssue(fmt.Errorf("Failed to read the base path:%s! %w\n", basePath, err))
	}
	foundDir, foundFile := false, false
	for _, entry := range entries {
		if entry.Name() == "store" {
			foundDir = true
			dirEntries, err := os.ReadDir(basePath)
			if err != nil {
				exitOnIssue(fmt.Errorf("Unable to read entries in store dir! %w\n", err))
			}
			for _, fEntry := range dirEntries {
				if fEntry.Name() == storePath {
					foundFile = true
					break
				}
			}
			if foundDir && foundFile {
				break
			}
		}
	}
	log.Println(foundDir, foundFile)
	if !foundDir {
		if err := os.Mkdir(basePath, 0755); err != nil {
			exitOnIssue(fmt.Errorf("failed to create store!%w\n", err))
		}
	}
	if !foundFile {
		if _, err := os.Create(fileStorePath); err != nil {
			exitOnIssue(fmt.Errorf("failed to create store!%w\n", err))
		}
	}
}

func GetNewFileStore() repository.RepositoryI {
	fh, err := os.OpenFile(fileStorePath, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("Failed to open the file store!")
		os.Exit(0)
	}
	return &FileStore{
		f: fh,
	}
}

func (fs *FileStore) GetToDoTasks() []models.Task {
	return []models.Task{}
}
func (fs *FileStore) GetAllTasks() []models.Task {
	return []models.Task{}
}
func (fs *FileStore) AddTask(*models.Task) *models.Task {
	return &models.Task{}
}
func (fs *FileStore) RemoveTask(id string) bool {
	return true
}
func (fs *FileStore) ModifyTask(*models.Task) *models.Task {
	return &models.Task{}
}
