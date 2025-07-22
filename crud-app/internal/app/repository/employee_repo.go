package repository

import (
	"crud-app/internal/app/model"

	"crud-app/pkg/db"
	"log"

	"fmt"
)

type EmployeeRepository interface {
	GetAll() (*[]model.EmployeeAllResponse, error)
	Create(model.EmployeeAllResponse) error
}

type employeeRepository struct {
}

func NewEmployeeRepository() EmployeeRepository {
	return &employeeRepository{}
}

func (r *employeeRepository) GetAll() (*[]model.EmployeeAllResponse, error) {
	database := db.GetDBConnection()
	var response []model.EmployeeAllResponse
	GetEmployees := `select * from "employee"`
	employees, err := database.Query(GetEmployees)
	if err != nil {
		return &response, fmt.Errorf("failed to create employee: %w", err)
	}
	for employees.Next() {
		var employee model.EmployeeAllResponse
		err := employees.Scan(&employee.Name, &employee.EmployeeID, &employee.Age)
		if err != nil {
			log.Println("error in scanning data")
		}
		response = append(response, employee)
	}
	return &response, err
}

func (r *employeeRepository) Create(employee model.EmployeeAllResponse) error {
	database := db.GetDBConnection()
	CreateEmployee := `INSERT INTO employee (name,employee_id,age) VALUE ($1,$2,$3)`
	res, err := database.Exec(CreateEmployee, employee.Name, employee.Age, employee.EmployeeID)
	if err != nil {
		log.Println(err)
	}

	val, err := res.RowsAffected()
	if err != nil {
		log.Println("error in creating data")
	}

	if val > 0 {
		log.Println("successfully inseared")
	} else {
		log.Println("failed o insert data")
	}
	return err

}
