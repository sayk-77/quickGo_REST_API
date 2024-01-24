package dependency

import (
	"example.com/go/pkg/controllers"
	"example.com/go/pkg/database"
	"example.com/go/pkg/service"
	servise "example.com/go/pkg/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SettingDepInjection(app *fiber.App, db *gorm.DB) {
	carRepository := database.NewCarRepository(db)
	carService := servise.NewCarService(carRepository)
	controllers.NewCarController(app, carService)

	clientRepository := database.NewClientRepository(db)
	clientService := servise.NewClientService(clientRepository)
	controllers.NewClientContorller(app, clientService)

	driverRepository := database.NewDriverRepository(db)
	driverService := servise.NewDriverService(driverRepository)
	controllers.NewDriverController(app, driverService)

	orderRepository := database.NewOrderRepository(db)
	orderService := servise.NewOrderService(orderRepository)
	controllers.NewOrderController(app, orderService)

	cargoTypeRepository := database.NewCargoTypeRepository(db)
	cargoTypeService := servise.NewCargoTypeService(cargoTypeRepository)
	controllers.NewCargoTypeController(app, cargoTypeService)

	transportationContractRepository := database.NewTransportationContractRepository(db)
	transportationContractSerivce := service.NewTransportationContractService(transportationContractRepository)
	controllers.NewTransportationContractController(app, transportationContractSerivce)

	waybillRepository := database.NewWaybillRepository(db)
	waybillService := service.NewWaybillService(waybillRepository)
	controllers.NewWaybillController(app, waybillService)
}
