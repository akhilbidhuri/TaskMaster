package repositoryfile

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/akhilbidhuri/TaskMaster/consts"
	"github.com/akhilbidhuri/TaskMaster/models"
	"github.com/akhilbidhuri/TaskMaster/repository"
	"github.com/google/uuid"
)

const storeFile = "tm_data.ndjson"
const indexFile = "tm_index.json"

var basePath = filepath.Clean("../store")

var fileStorePath = filepath.Join(basePath, storeFile)
var indexPath = filepath.Join(basePath, indexFile)

func exitOnIssue(msg interface{}) {
	log.Fatal(msg)
	os.Exit(0)
}

type FileStore struct {
	F     *os.File
	index repository.Index
}

// if dir and file store not present create
func init() {
	entries, err := os.ReadDir(filepath.Clean("../"))
	if err != nil {
		exitOnIssue(fmt.Errorf("Failed to read the base path:%s! %w\n", basePath, err))
	}
	foundDir, foundFile, foundIndex := false, false, false
	for _, entry := range entries {
		if entry.Name() == "store" {
			foundDir = true
			dirEntries, err := os.ReadDir(basePath)
			if err != nil {
				exitOnIssue(fmt.Errorf("Unable to read entries in store dir! %w\n", err))
			}
			for _, fEntry := range dirEntries {
				if fEntry.Name() == storeFile {
					foundFile = true
				} else if fEntry.Name() == indexFile {
					foundIndex = true
				}
				if foundFile && foundIndex {
					break
				}
			}
			if foundDir && foundFile {
				break
			}
		}
	}
	if !foundDir {
		if err := os.Mkdir(basePath, 0755); err != nil {
			exitOnIssue(fmt.Errorf("failed to create store!%w\n", err))
		}
	}
	if !foundFile {
		if f, err := os.Create(fileStorePath); err != nil {
			exitOnIssue(fmt.Errorf("failed to create store!%w\n", err))
		} else {
			defer f.Close()
		}
	}
	if !foundIndex {
		var indf *os.File
		var err error
		if indf, err = os.Create(indexPath); err != nil {
			exitOnIssue(fmt.Errorf("failed to create store!%w\n", err))
		} else {
			defer indf.Close()
		}
		store, err := os.Open(fileStorePath)
		if err != nil {
			log.Fatalln("failed to read store to create index")
		}
		indexMap := construcIndexFromStore(store)
		index := NdJsonIndex{
			Index: indexMap,
		}
		indexJson, err := json.Marshal(index)
		if err != nil {
			log.Fatal("failed to jsonify index for writing!, ", err)
		}
		writer := bufio.NewWriter(indf)
		writer.Write(indexJson)
		writer.Flush()
	}
}

func GetNewFileStore() *FileStore {
	fh, err := os.OpenFile(fileStorePath, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("Failed to open the file store!, ", err)
	}
	indf, err := os.OpenFile(fileStorePath, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("Failed to open the file index!, ", err)
	}
	indf.Close()
	index, err := getNdJSONINdex(indexPath)
	if err != nil {
		log.Fatal("failed to get index from file!, ", err)
	}
	return &FileStore{
		F:     fh,
		index: index,
	}
}

func (fs *FileStore) GetToDoTasks() []models.Task {
	return []models.Task{}
}

func (fs *FileStore) GetAllTasks() []models.Task {
	return []models.Task{}
}

func (fs *FileStore) AddTask(task *models.Task) *models.Task {
	task.ID = uuid.NewString()
	task.Status = consts.PENDING
	task.Created_At = time.Now()
	fs.F.Seek(0, io.SeekEnd)
	if taskJson, err := json.Marshal(task); err != nil {
		log.Fatal("Failed to store the task!", err)
		return nil
	} else {
		fs.F.Write(append(taskJson, byte('\n')))
	}
	return task
}

func (fs *FileStore) RemoveTask(id string) bool {
	return true
}

func (fs *FileStore) ModifyTask(*models.Task) *models.Task {
	return &models.Task{}
}
