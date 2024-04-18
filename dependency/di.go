package dependency

import (
	"example.com/go/pkg/controllers"
	"example.com/go/pkg/database"
	"example.com/go/pkg/service"
	servise "example.com/go/pkg/service"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SettingDepInjection(app *fiber.App, db *gorm.DB, redis *redis.Client) {
	carRepository := database.NewCarRepository(db)
	carService := servise.NewCarService(carRepository)
	controllers.NewCarController(app, carService)

	clientRepository := database.NewClientRepository(db)
	clientService := servise.NewClientService(clientRepository)
	controllers.NewClientContorller(app, clientService)

	driverRepository := database.NewDriverRepository(db)
	driverService := servise.NewDriverService(driverRepository)
	controllers.NewDriverController(app, driverService)

	cargoTypeRepository := database.NewCargoTypeRepository(db)
	cargoTypeService := servise.NewCargoTypeService(cargoTypeRepository)
	controllers.NewCargoTypeController(app, cargoTypeService)

	transportationContractRepository := database.NewTransportationContractRepository(db)
	transportationContractSerivce := service.NewTransportationContractService(transportationContractRepository)
	controllers.NewTransportationContractController(app, transportationContractSerivce)

	waybillRepository := database.NewWaybillRepository(db)
	waybillService := service.NewWaybillService(waybillRepository)
	controllers.NewWaybillController(app, waybillService)

	orderRepository := database.NewOrderRepository(db)
	orderService := servise.NewOrderService(orderRepository, driverRepository, transportationContractRepository, waybillRepository)
	controllers.NewOrderController(app, orderService)

	feedbackRepository := database.NewFeedbackRepository(db)
	feedbackService := service.NewFeedbackService(feedbackRepository)
	controllers.NewFeedbackController(app, feedbackService)

	emailService := service.NewEmailService("smtp.gmail.com", "587", "yawaihv2@gmail.com", "bdkp ntae lrro bswq")
	controllers.NewEmailController(app, emailService)

	passwordRecoveryService := service.NewPasswordRecoveryService(redis, db)
	controllers.NewPasswordRecoveryController(app, passwordRecoveryService)
}
