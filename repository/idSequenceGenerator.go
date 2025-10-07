package repository

import (
	"encoding/json"
	"os"
	"strconv"
)

type Sequene struct {
	Current int `json:"current"`
}

type IDSequenceGenerator struct {
	seqFile *os.File
}

func (idsq *IDSequenceGenerator) Init(filePath string) error {
	if f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755); err != nil {
		return err
	} else {
		idsq.seqFile = f
		seq := Sequene{}
		defer idsq.seqFile.Sync()
		if err := json.NewDecoder(f).Decode(&seq); err != nil {
			seq.Current = 0
			f.Seek(0, 0)
			f.Truncate(0)
			json.NewEncoder(f).Encode(seq)
		} else {
			json.NewEncoder(f).Encode(seq)
		}
	}
	return nil
}

func (idsq *IDSequenceGenerator) GetNextID() (string, error) {
	var seq Sequene
	idsq.seqFile.Seek(0, 0)
	if err := json.NewDecoder(idsq.seqFile).Decode(&seq); err != nil {
		return "", err
	}
	id := seq.Current
	seq.Current++
	idsq.seqFile.Seek(0, 0)
	idsq.seqFile.Truncate(0)
	idsq.seqFile.Sync()
	if err := json.NewEncoder(idsq.seqFile).Encode(seq); err != nil {
		return "", err
	}
	idsq.seqFile.Sync()
	return strconv.Itoa(id), nil
}
