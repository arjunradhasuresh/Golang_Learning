package router

import (
	"crud-app/internal/app/controller"
	"crud-app/internal/app/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	statusService = service.NewStatusService()
	statusCtrl    = controller.NewStatusController(statusService)
)

func InitRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger) // Log all requests
	r.Route("/", func(r chi.Router) {
		r.Get("/status", statusCtrl.Status) // Status endpoint
		r.Group(func(r chi.Router) {
			r.Mount("/employee", EmployeeRouter())
		})
	})
	return r
}
