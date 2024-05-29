package tasks

import (
	"todo/internal/models"

	"github.com/gofrs/uuid/v5"
	"github.com/jmoiron/sqlx"
)

type service struct {
	db *sqlx.DB
}

// NewService creates a new home service
func NewService(db *sqlx.DB) models.TaskService {
	return &service{db: db}
}

func (s *service) Create(task *models.Task) error {
	id := uuid.Must(uuid.NewV4())
	task.ID = id

	_, err := s.db.Exec("INSERT INTO tasks (id, title) VALUES (?, ?)", task.ID, task.Title)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Update(task *models.Task) error {
	_, err := s.db.Exec("UPDATE tasks SET title = ?, completed = ? WHERE id = ?", task.Title, task.Completed, task.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Delete(id uuid.UUID) error {
	_, err := s.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Complete(id uuid.UUID) error {
	_, err := s.db.Exec("UPDATE tasks SET completed = 1 WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) List() ([]models.Task, error) {
	var tasks []models.Task
	err := s.db.Select(&tasks, "SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *service) Find(id uuid.UUID) (models.Task, error) {
	var task models.Task
	err := s.db.Get(&task, "SELECT * FROM tasks WHERE id = ?", id)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}
