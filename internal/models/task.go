package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type Task struct {
	ID        uuid.UUID `db:"id"`
	Title     string    `db:"title"`
	Completed bool      `db:"completed"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type TaskService interface {
	Create(task *Task) error
	Update(task *Task) error
	Delete(id uuid.UUID) error
	Complete(id uuid.UUID) error
	List() ([]Task, error)
	Find(id uuid.UUID) (Task, error)
}
