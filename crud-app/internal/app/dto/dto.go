package dto

type StatusResponse struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Status  string `json:"status"`
}

type EmployeeAllResponse struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	EmployeeID int    `json:"employee_id"`
}
