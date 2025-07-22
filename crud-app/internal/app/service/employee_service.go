package service

import (
	"crud-app/internal/app/dto"
	"crud-app/internal/app/model"
	"crud-app/internal/app/repository"
)

type EmployeeService interface {
	GetAllEmployees() (*[]dto.EmployeeAllResponse, error)
	CreateEmployee(dto.EmployeeAllResponse) error
}
type employeeService struct {
	EmployeeRepo repository.EmployeeRepository
}

func NewEmployeeService(employeeRepo repository.EmployeeRepository) EmployeeService {
	return &employeeService{
		EmployeeRepo: employeeRepo,
	}
}

func (s *employeeService) GetAllEmployees() (*[]dto.EmployeeAllResponse, error) {
	response, err := s.EmployeeRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var employee []dto.EmployeeAllResponse
	for _, val := range *response {
		employee = append(employee, dto.EmployeeAllResponse{
			Name:       val.Name,
			EmployeeID: val.EmployeeID,
			Age:        val.Age,
		})
	}
	return &employee, err
}

func (s *employeeService) CreateEmployee(req dto.EmployeeAllResponse) error {
	employee := model.EmployeeAllResponse{
		Name:       req.Name,
		EmployeeID: req.EmployeeID,
		Age:        req.Age,
	}

	err := s.EmployeeRepo.Create(employee)
	if err != nil {
		return err
	}
	return nil
}
