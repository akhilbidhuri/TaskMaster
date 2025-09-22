package models

import (
	"fmt"
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

func (t Task) String() string {
	return fmt.Sprintf("\n----------------\n\tID: %v\n\tTitle: %v\n\tdesc: %v\n\tresources: %v\n\tstatus: %v\n",
		t.ID,
		t.Title,
		t.Desc,
		t.Res,
		t.Status,
	)
}
