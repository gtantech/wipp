package pdm

import (
	"fmt"
	"uuid"
)

type ProjectNotFoundError struct {
	Id      uuid.UUID
	Message string
}

func NewProjectNotFoundError(id uuid.UUID) *ProjectNotFoundError {
	return &ProjectNotFoundError{Id: id, Message: fmt.Sprintf("Project not found for id: %v", id)}
}

func (p *ProjectNotFoundError) Error() string {
	return p.Message
}
