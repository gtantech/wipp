package pdm

import (
	"time"
	"uuid"

	"github.com/gtantech/pdm/activity"
)

type task struct {
	id uuid.UUID
	activity.Data
}

func NewTask(id uuid.UUID, duration time.Duration) *task {
	return &task{id: id, Data: activity.NewData(duration)}
}
