package model

type EmployeeAllResponse struct {
	Name       string `json:"name"`
	EmployeeID string `json:"employee_id"`
	Age        int    `json:"age"`
}
