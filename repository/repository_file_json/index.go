package repositoryfile

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/akhilbidhuri/TaskMaster/models"
)

type NdJsonIndex struct {
	Index map[string]int64 `json:"index"`
	f     *os.File         `json:"-"`
}

func getNdJSONINdex(indexPath string) (*NdJsonIndex, error) {
	if indexF, err := os.Open(indexPath); err != nil {
		log.Fatal("unable to get index file, ", err)
	} else {
		var index NdJsonIndex
		if err := json.NewDecoder(indexF).Decode(&index); err != nil {
			log.Fatal("not able to decode index, ", err)
		}
		index.f = indexF
		return &index, nil
	}
	return nil, errors.New("could not get index")
}

func (i *NdJsonIndex) Add(id string, offset int64) error {
	i.Index[id] = offset
	i.f.Seek(0, 0)
	i.f.Truncate(0)
	defer seekStart(i.f)
	json.NewEncoder(i.f).Encode(i)
	return nil
}

func (i *NdJsonIndex) Remove(id string) error {
	if _, err := i.Find(id); err != nil {
		return errors.New("task is not Present!")
	}
	delete(i.Index, id)
	i.f.Seek(0, 0)
	i.f.Truncate(0)
	defer seekStart(i.f)
	json.NewEncoder(i.f).Encode(i)
	return nil
}

func (i *NdJsonIndex) Find(id string) (int64, error) {
	offest, ok := i.Index[id]
	if !ok {
		return -1, errors.New("task not present!")
	}
	return offest, nil
}

func construcIndexFromStore(f *os.File) map[string]int64 {
	var index = make(map[string]int64)
	fScanner := bufio.NewScanner(f)
	offset := 0
	newLineOffset := len([]byte("\n"))
	for fScanner.Scan() {
		record := fScanner.Bytes()
		var task models.Task
		if err := json.Unmarshal(record, &task); err != nil {
			log.Fatalln("failed to read store! for constructing index")
		}
		id := task.ID
		index[id] = int64(offset)
		offset += len(record) + newLineOffset
	}
	return index
}
