package router

import (
	"crud-app/internal/app/controller"
	"crud-app/internal/app/repository"
	"crud-app/internal/app/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	employeeRepo    = repository.NewEmployeeRepository()
	employeeService = service.NewEmployeeService(employeeRepo)
	employeeCtrl    = controller.NewEmployeeController(employeeService)
)

func EmployeeRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger) // Log all requests

	router.Route("/employees", func(r chi.Router) {
		r.Get("/", employeeCtrl.GetAllEmployees) // Get all employees
		r.Post("/", employeeCtrl.CreateEmployee) // Create a new employee
	})

	return router
}
