package controller

import (
	"crud-app/internal/app/dto"
	"crud-app/internal/app/service"
	"crud-app/pkg/api"
	"encoding/json"
	"net/http"
)

type EmployeeController interface {
	GetAllEmployees(w http.ResponseWriter, r *http.Request)
	CreateEmployee(w http.ResponseWriter, r *http.Request)
}

type employeeController struct {
	EmployeeService service.EmployeeService
}

func NewEmployeeController(employeeService service.EmployeeService) EmployeeController {
	return &employeeController{
		EmployeeService: employeeService,
	}
}

func (ec *employeeController) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := ec.EmployeeService.GetAllEmployees()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	api.Success(w, "success", employees)

}

func (ec *employeeController) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee dto.EmployeeAllResponse
	// Decode the JSON body into the employee struct
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = ec.EmployeeService.CreateEmployee(employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	api.Success(w, "SUCCESS", nil)
}
