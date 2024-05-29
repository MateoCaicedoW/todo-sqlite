package tasks

import (
	"net/http"
	"todo/internal/models"

	"github.com/gofrs/uuid/v5"
	"github.com/leapkit/core/render"
)

func Complete(w http.ResponseWriter, r *http.Request) {
	service := r.Context().Value("taskService").(models.TaskService)
	id := uuid.FromStringOrNil(r.PathValue("id"))
	task, err := service.Find(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	task.Completed = true
	complete := r.URL.Query().Get("complete")
	if complete == "false" {
		task.Completed = false
	}

	err = service.Update(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rw := render.FromCtx(r.Context())
	rw.Set("task", task)
	err = rw.RenderClean("tasks/task.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
