package tasks

import (
	"database/sql"
	"todo/internal/models"

	"github.com/gofrs/uuid/v5"
)

type service struct {
	db *sql.DB
}

// NewService creates a new home service
func NewService(db *sql.DB) models.TaskService {
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
	rows, err := s.db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Completed, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *service) Find(id uuid.UUID) (models.Task, error) {
	var task models.Task
	row := s.db.QueryRow("SELECT * FROM tasks WHERE id = ?", id)

	err := row.Scan(&task.ID, &task.Title, &task.Completed, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}
