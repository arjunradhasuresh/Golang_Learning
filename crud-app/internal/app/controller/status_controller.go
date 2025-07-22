package controller

import (
	"crud-app/internal/app/service"
	"fmt"
	"net/http"
)

type StatusController interface {
	Status(w http.ResponseWriter, r *http.Request)
}

type statusControllerImpl struct {
	StatusService service.StatusService
}

func NewStatusController(service service.StatusService) StatusController {
	return &statusControllerImpl{
		StatusService: service,
	}
}

func (s *statusControllerImpl) Status(w http.ResponseWriter, r *http.Request) {
	response := s.StatusService.GetStatus()
	fmt.Println(response)

}
