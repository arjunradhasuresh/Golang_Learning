package model

type EmployeeAllResponse struct {
	Name       string `json:"name"`
	EmployeeID int    `json:"employee_id"`
	Age        int    `json:"age"`
}
