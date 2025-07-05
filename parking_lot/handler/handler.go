package handler

import (
	"parkinglot/domain"
	"parkinglot/service"
	"time"

	"github.com/gin-gonic/gin"
)

type DeAllocateParkingSpotRequest struct {
	Id int64 `json:"ticket"`
}

type AllocateParkingSpotResponse struct {
	Id        int64          `json:"id"`
	Spot      string        `json:"spot"`
	CheckedIn time.Time      `json:"checked_in"`
	Status    string         `json:"status"`
	Vehicle   *domain.Vehicle `json:"vehicle"`
}

type ParkingLotHandler struct {
	parkingLotManager *service.ParkingLotManager
}

func NewParkingLotHandler(parkingLotManager *service.ParkingLotManager) *ParkingLotHandler {
	return &ParkingLotHandler{
		parkingLotManager: parkingLotManager,
	}
}

func (h *ParkingLotHandler) AllocateParkingSpot(c *gin.Context) {
	var vehicle domain.Vehicle
	if err := c.ShouldBindJSON(&vehicle); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ticket, err := h.parkingLotManager.AllocateParkingSpot(&vehicle)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"AllocateParkingSpotResponse": AllocateParkingSpotResponse{
			Id:        ticket.Id,
			Spot:      ticket.Spot.Location,
			CheckedIn: ticket.CheckedIn,
			Status:    ticket.Status,
			Vehicle:   ticket.Vehicle,
		},
	})
}

func (h *ParkingLotHandler) DeAllocateParkingSpot(c *gin.Context) {
	var request DeAllocateParkingSpotRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := h.parkingLotManager.DeAllocateParkingSpot(request.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// return result, nil
	c.JSON(200, result)
}
