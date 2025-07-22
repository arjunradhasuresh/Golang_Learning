package service

import "crud-app/internal/app/dto"

type StatusService interface {
	GetStatus() *dto.StatusResponse
}

type statusServiceImpl struct{}

func NewStatusService() StatusService {
	return &statusServiceImpl{}
}

func (s *statusServiceImpl) GetStatus() *dto.StatusResponse {
	return &dto.StatusResponse{
		Name:    "CrudApp",
		Version: "1.0.0",
		Status:  "Running",
	}
}
