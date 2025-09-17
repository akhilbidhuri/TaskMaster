package models

import (
	"time"

	"github.com/akhilbidhuri/TaskMaster/consts"
)

type Resources []string

type Task struct {
	ID           string          `json:"id"`
	Title        string          `json:"title"`
	Desc         string          `json:"desc"`
	Res          Resources       `json:"res"`
	Status       consts.T_Status `json:"status"`
	Created_At   time.Time       `json:"created_at"`
	Completed_At time.Time       `json:"completed_at,omitempty"`
}
