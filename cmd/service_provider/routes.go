package service_provider

import (
	"net/http"
)

func (s *ServiceProvider) GetRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /classes", s.GetClassController().Create)
	mux.HandleFunc("PUT /classes/{id}", s.GetClassController().Update)
	mux.HandleFunc("GET /classes", s.GetClassController().FindAll)

	mux.HandleFunc("POST /lessons", s.GetLessonController().Create)
	mux.HandleFunc("PUT /lessons{id}", s.GetLessonController().Update)

	mux.HandleFunc("POST /teachers", s.GetTeacherController().Create)
	mux.HandleFunc("PUT /teachers/{id}", s.GetTeacherController().Update)

	return mux
}
