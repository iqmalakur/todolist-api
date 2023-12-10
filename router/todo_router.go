package router

import (
	"fmt"
	"net/http"
	"todolist/controller"
)

var todoController = controller.NewTodoController()

func TodoRouter(w http.ResponseWriter, r *http.Request) {
	todoId := r.URL.Path[len("/todos/"):]

	fmt.Println(r.Method, r.URL.Path)

	switch {
	case todoId == "" && r.Method == "GET":
		todoController.Index(w, r)
	case todoId == "" && r.Method == "POST":
		todoController.Create(w, r)
	case r.Method == "GET":
		todoController.Show(w, r)
	case r.Method == "PUT":
		todoController.Update(w, r)
	case r.Method == "DELETE":
		todoController.Delete(w, r)
	}
}
