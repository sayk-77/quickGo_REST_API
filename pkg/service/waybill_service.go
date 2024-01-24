package service

import (
	"example.com/go/models"
	"example.com/go/pkg/database"
)

type WaybillService struct {
	waybillRepository database.WaybillRepository
}

func NewWaybillService(waybillRepository database.WaybillRepository) *WaybillService {
	return &WaybillService{
		waybillRepository: waybillRepository,
	}
}

func (ws *WaybillService) GetWaybillById(waybillId int) (*models.WayBill, error) {
	return ws.waybillRepository.GetWaybillById(waybillId)
}

func (ws *WaybillService) GetAllWaybill() ([]*models.WayBill, error) {
	return ws.waybillRepository.GetAllWaybill()
}

func (ws *WaybillService) CreateNewWaybill(newWaybill *models.WayBill) (*models.WayBill, error) {
	return ws.waybillRepository.CreateNewWaybill(newWaybill)
}
