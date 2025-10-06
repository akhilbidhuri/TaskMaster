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
)

var basePath = filepath.Clean("../store")

var fileStorePath = filepath.Join(basePath, consts.StoreFile)
var indexPath = filepath.Join(basePath, consts.IndexFile)

func exitOnIssue(msg interface{}) {
	log.Fatal(msg)
	os.Exit(0)
}

type FileStore struct {
	F     *os.File
	index repository.Index
	sqGen *repository.IDSequenceGenerator
}

// if dir and file store not present create
func init() {
	entries, err := os.ReadDir(filepath.Clean("../"))
	if err != nil {
		exitOnIssue(fmt.Errorf("Failed to read the base path:%s! %w\n", basePath, err))
	}
	foundDir, foundFile, foundIndex, foundSeqGen := false, false, false, false
	for _, entry := range entries {
		if entry.Name() == "store" {
			foundDir = true
			dirEntries, err := os.ReadDir(basePath)
			if err != nil {
				exitOnIssue(fmt.Errorf("Unable to read entries in store dir! %w\n", err))
			}
			for _, fEntry := range dirEntries {
				if fEntry.Name() == consts.StoreFile {
					foundFile = true
				} else if fEntry.Name() == consts.IndexFile {
					foundIndex = true
				} else if fEntry.Name() == consts.IdSequenceFile {
					foundSeqGen = true
				}
				if foundFile && foundIndex && foundSeqGen {
					break
				}
			}
			if foundDir && foundFile && foundIndex && foundSeqGen {
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
	if !foundSeqGen {
		var seqGenFile *os.File
		var err error
		if seqGenFile, err = os.Create(filepath.Join(basePath, consts.IdSequenceFile)); err != nil {
			exitOnIssue(fmt.Errorf("failed to create id sequence file!%w\n", err))
		} else {
			defer seqGenFile.Close()
		}
		seqGen := repository.Sequene{
			Current: 0,
		}
		seqGenJson, err := json.Marshal(seqGen)
		if err != nil {
			log.Fatal("failed to jsonify index for writing!, ", err)
		}
		writer := bufio.NewWriter(seqGenFile)
		writer.Write(seqGenJson)
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
	sqGen := &repository.IDSequenceGenerator{}
	if err := sqGen.Init(filepath.Join(basePath, consts.IdSequenceFile)); err != nil {
		log.Fatal("failed to init sequence generator!, ", err)
	}
	return &FileStore{
		F:     fh,
		index: index,
		sqGen: sqGen,
	}
}

func (fs *FileStore) GetToDoTasks() []models.Task {
	var tasks = make([]models.Task, 0)
	defer seekStart(fs.F)
	reader := bufio.NewScanner(fs.F)
	for reader.Scan() {
		taskMarshalled := reader.Bytes()
		var task models.Task
		json.Unmarshal([]byte(taskMarshalled), &task)
		if task.Status == consts.PENDING {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

func (fs *FileStore) GetAllTasks() []models.Task {
	var tasks = make([]models.Task, 0)
	defer seekStart(fs.F)
	reader := bufio.NewScanner(fs.F)
	for reader.Scan() {
		taskMarshalled := reader.Bytes()
		var task models.Task
		json.Unmarshal([]byte(taskMarshalled), &task)
		tasks = append(tasks, task)
	}
	return tasks
}

func (fs *FileStore) AddTask(task *models.Task) *models.Task {
	id, err := fs.sqGen.GetNextID()
	if err != nil {
		log.Fatal("Failed to generate ID for task!, ", err)
	}
	task.ID = id
	task.Status = consts.PENDING
	task.Created_At = time.Now()
	offset, _ := fs.F.Seek(0, io.SeekEnd)
	defer seekStart(fs.F)
	if taskJson, err := json.Marshal(task); err != nil {
		log.Fatal("Failed to store the task!", err)
		return nil
	} else {
		fs.F.Write(append(taskJson, byte('\n')))
	}
	if err := fs.index.Add(task.ID, offset); err != nil {
		log.Fatal("Failed to add task to index!, ", err)
	}
	return task
}

func (fs *FileStore) RemoveTask(id string) bool {
	return true
}

func (fs *FileStore) ModifyTask(task *models.Task) *models.Task {
	offest, err := fs.index.Find(task.ID)
	if err != nil {
		log.Fatal("Task dosen't exit!")
	}
	defer seekStart(fs.F)
	fs.F.Seek(offest, io.SeekStart)
	reader := bufio.NewScanner(fs.F)
	reader.Scan()
	taskMarshalled := reader.Bytes()
	var existingTask models.Task
	json.Unmarshal([]byte(taskMarshalled), &existingTask)
	if task.Title != "" {
		existingTask.Title = task.Title
	}
	if task.Desc != "" {
		existingTask.Desc = task.Desc
	}
	if len(task.Res) != 0 {
		existingTask.Res = task.Res
	}
	fs.index.Remove(task.ID)
	fs.AddTask(&existingTask)
	return &models.Task{}
}

func (fs *FileStore) TaskExists(id string) bool {
	if _, err := fs.index.Find(id); err != nil {
		return false
	}
	return true
}

func (fs *FileStore) MarkTaskDone(id string) bool {
	offset, err := fs.index.Find(id)
	if err != nil {
		log.Fatal("Task dosen't exit!")
	}
	fs.F.Seek(offset, io.SeekStart)
	var task models.Task
	if err := json.NewDecoder(fs.F).Decode(&task); err != nil {
		log.Fatal("Failed to decode the task from store!, ", err)
	}
	task.Status = consts.DONE
	fs.F.Seek(offset, io.SeekStart)
	if err := json.NewEncoder(fs.F).Encode(&task); err != nil {
		log.Fatal("Failed to encode the task to store!, ", err)
	}
	log.Println("Marked task as done!", task.ID, task.Title)
	return true
}

func seekStart(f *os.File) {
	f.Seek(0, io.SeekStart)
}
