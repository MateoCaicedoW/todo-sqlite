package tasks

import (
	"net/http"
	"todo/internal/models"

	"github.com/gofrs/uuid/v5"
	"github.com/leapkit/core/render"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	rw := render.FromCtx(r.Context())
	service := r.Context().Value("taskService").(models.TaskService)
	id := uuid.FromStringOrNil(r.PathValue("id"))
	err := service.Delete(id)
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
