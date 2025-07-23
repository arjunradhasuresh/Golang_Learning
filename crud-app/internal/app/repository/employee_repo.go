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
	GetEmployees := `SELECT name, employee_id, age FROM "employees"`
	employees, err := database.Query(GetEmployees)
	if err != nil {
		return &response, fmt.Errorf("failed to create emploss: %w", err)
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
	query := `INSERT INTO employees (name, employee_id, age) VALUES ($1, $2, $3)`

	res, err := database.Exec(query, employee.Name, employee.EmployeeID, employee.Age)
	if err != nil {
		log.Printf("Error executing insert query: %v\n", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected: %v\n", err)
		return err
	}

	if rowsAffected > 0 {
		log.Println("Employee successfully inserted.")
	} else {
		log.Println("Insert operation did not affect any rows.")
	}

	return nil
}
