package service

import (
	"drumkit.com/interview/src/gateway"
	"drumkit.com/interview/src/model"
)

type LoadService struct {
	repo *gateway.TurvoAPIGateway
}

func NewLoadService(repo *gateway.TurvoAPIGateway) *LoadService {
	return &LoadService{repo: repo}
}

func (s *LoadService) CreateLoad(load model.CreateLoadRequest) error {
	return s.repo.CreateLoad([]model.CreateLoadRequest{load})
}

func (s *LoadService) RetrieveLoads(start, pageSize string) ([]model.Shipment, error) {
	return s.repo.RetrieveLoads(start, pageSize)
}
