package service

import (
	"drumkit.com/interview/src/gateway"
	"drumkit.com/interview/src/model"
)

type LoadService struct {
	gw *gateway.TurvoAPIGateway
}

func NewLoadService(gw *gateway.TurvoAPIGateway) *LoadService {
	return &LoadService{gw: gw}
}

func (s *LoadService) CreateLoad(load model.CreateLoadRequest) error {
	pickUp, err := s.gw.RetrieveLocations(load.Pickup.Name)
	if err != nil {
		return err
	}
	PickUpId := pickUp[0].ID

	Delivery, err := s.gw.RetrieveLocations(load.Consignee.Name)
	if err != nil {
		return err
	}

	DeliveryId := Delivery[0].ID
	return s.gw.CreateLoad([]model.CreateLoadRequest{load}, PickUpId, DeliveryId)
}

func (s *LoadService) RetrieveLoads(start, pageSize string) ([]model.Shipment, error) {
	return s.gw.RetrieveLoads(start, pageSize)
}
