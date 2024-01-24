package database

import (
	"example.com/go/models"
	"gorm.io/gorm"
)

type WaybillRepository interface {
	GetWaybillById(waybillID int) (*models.WayBill, error)
	GetAllWaybill() ([]*models.WayBill, error)
	CreateNewWaybill(newWaybill *models.WayBill) (*models.WayBill, error)
}

type WaybillRepositoryImpl struct {
	db *gorm.DB
}

func NewWaybillRepository(db *gorm.DB) *WaybillRepositoryImpl {
	return &WaybillRepositoryImpl{
		db: db,
	}
}

func (wr *WaybillRepositoryImpl) GetWaybillById(waybillID int) (*models.WayBill, error) {
	var waybill models.WayBill
	if err := wr.db.Preload("Driver").Preload("Car").Preload("TransportationContract").First(&waybill, waybillID).Error; err != nil {
		return nil, err
	}

	return &waybill, nil
}

func (wr *WaybillRepositoryImpl) GetAllWaybill() ([]*models.WayBill, error) {
	var waybillRecord []*models.WayBill
	if err := wr.db.Preload("Driver").Preload("Car").Preload("TransportationContract").Find(&waybillRecord).Error; err != nil {
		return nil, err
	}

	return waybillRecord, nil
}

func (wr *WaybillRepositoryImpl) CreateNewWaybill(newWaybill *models.WayBill) (*models.WayBill, error) {
	if err := wr.db.Create(newWaybill).Error; err != nil {
		return nil, err
	}

	return wr.GetWaybillById(int(newWaybill.ID))
}
