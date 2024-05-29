package tasks

import (
	"net/http"
	"todo/internal/models"

	"github.com/leapkit/core/render"
)

func List(w http.ResponseWriter, r *http.Request) {
	rw := render.FromCtx(r.Context())
	service := r.Context().Value("taskService").(models.TaskService)

	tasks, err := service.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Set("tasks", tasks)
	err = rw.Render("tasks/list.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
