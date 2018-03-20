package models

import (
	"time"

	"github.com/smartcontractkit/chainlink/utils"
	"go.uber.org/multierr"
	null "gopkg.in/guregu/null.v3"
)

type AssignmentSpec struct {
	Assignment Assignment `json:"assignment"`
	Schedule   Schedule   `json:"schedule"`
}

type Assignment struct {
	Subtasks []Subtask `json:"subtasks"`
}

type Subtask struct {
	Type   string `json:"adapterType"`
	Params JSON   `json:"adapterParams"`
}

type Schedule struct {
	EndAt null.Time `json:"endAt"`
	RunAt []Time    `json:"runAt"`
}

func (s AssignmentSpec) ConvertToJobSpec() (JobSpec, error) {
	var merr error
	tasks := []TaskSpec{}
	for _, st := range s.Assignment.Subtasks {
		params, err := st.Params.Add("type", st.Type)
		multierr.Append(merr, err)
		tasks = append(tasks, TaskSpec{
			Type:   st.Type,
			Params: params,
		})
	}
	initiators := []Initiator{{Type: "web"}}
	for _, r := range s.Schedule.RunAt {
		initiators = append(initiators, Initiator{
			Type: "runAt",
			Time: r,
		})
	}

	j := JobSpec{
		ID:         utils.NewBytes32ID(),
		CreatedAt:  Time{Time: time.Now()},
		Tasks:      tasks,
		EndAt:      s.Schedule.EndAt,
		Initiators: initiators,
	}

	return j, merr
}
