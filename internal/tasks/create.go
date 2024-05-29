package tasks

import (
	"net/http"
	"todo/internal/models"

	"github.com/leapkit/core/render"
)

func Create(w http.ResponseWriter, r *http.Request) {
	rw := render.FromCtx(r.Context())
	service := r.Context().Value("taskService").(models.TaskService)
	title := r.FormValue("title")
	task := models.Task{
		Title: title,
	}
	err := service.Create(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tasks, err := service.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Set("tasks", tasks)
	err = rw.RenderClean("tasks/table.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
