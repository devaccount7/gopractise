package router

import (
	"parkinglot/domain"
	"parkinglot/handler"
	"parkinglot/service"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	paymentService := service.NewPaymentServiceImpl()
	slots := getSlots() // Initialize parking slots
	parkingLotManager := service.NewParkingLotManager(slots, paymentService)
	parkingLotHandler := handler.NewParkingLotHandler(parkingLotManager)

	router.POST("/allocate", parkingLotHandler.AllocateParkingSpot)
	router.POST("/deallocate", parkingLotHandler.DeAllocateParkingSpot)
	return router

}

func getSlots() map[string]*domain.ParkingSpot {
	slots := make(map[string]*domain.ParkingSpot)
	slots["car"] = &domain.ParkingSpot{
		Id:          1,
		Type:        "car",
		IsAvailable: true,
		Location:    "Floor 1, Section A",
	}
	slots["bike"] = &domain.ParkingSpot{
		Id:          2,
		Type:        "bike",
		IsAvailable: true,
		Location:    "Floor 1, Section B",
	}
	slots["truck"] = &domain.ParkingSpot{
		Id:          3,
		Type:        "truck",
		IsAvailable: true,
		Location:    "Floor 1, Section A",
	}

	return slots
}
