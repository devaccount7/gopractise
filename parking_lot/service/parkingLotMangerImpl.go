package service

import (
	"fmt"
	"parkinglot/domain"
	"sync"
	"time"
)

var once sync.Once

type ParkingLotManager struct {
	// todo : can we have a slice of parkingspots instead of map?
	// this will allow us to iterate over the spots easily
	// or we can have a map of slice map[string][]*domain.ParkingSpot
	// where key is the type of vehicle and value is a slice of parking spots
	Slots          map[string]*domain.ParkingSpot
	PaymentService *PaymentServiceImpl
}

var singleInstance *ParkingLotManager
var ticketStore = make(map[int64]*domain.ParkingTicket)
var ticketCounter int64 = 0

func getInstance() *ParkingLotManager {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstance = &ParkingLotManager{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

func NewParkingLotManager(spot map[string]*domain.ParkingSpot, paymentService *PaymentServiceImpl) *ParkingLotManager {
	newInstance := getInstance()
	newInstance.PaymentService = paymentService
	newInstance.Slots = spot
	return newInstance
}

func (p *ParkingLotManager) AllocateParkingSpot(entryGate *domain.Entry, vehicle *domain.Vehicle) (*domain.ParkingTicket, error) {
	for _, spot := range p.Slots {
		if spot.IsAvailable && spot.Type == vehicle.Type {
			spot.IsAvailable = false
			spot.Location = "some location_coordinates_or_floor" // This should be set to a specific location
			ticket := domain.NewParkingTicket(spot, vehicle, entryGate)
			fmt.Println("Spot allocated successfully:", spot.Id)
			return ticket, nil
		}
	}
	return nil, fmt.Errorf("no available spot for vehicle: %v", vehicle)
}

func (p *ParkingLotManager) DeAllocateParkingSpot(ticketId int64) (*domain.ParkingTicket, error) {
	ticket, exists := ticketStore[ticketId]
	if !exists {
		return nil, fmt.Errorf("ticket %d not found", ticketId)
	}
	ticket.CheckedOut = time.Now()
	p.calculateParkingFee(ticket)
	p.PaymentService.ProcessPayment(ticket)
	for _, spot := range p.Slots {
		if spot.Id == ticket.Spot.Id {
			spot.IsAvailable = true
			fmt.Println("Spot deallocated successfully:", spot.Id)
			return ticket, nil
		}
	}
	fmt.Print(ticket.Spot.Id, " not found in parking lot")
	return nil, fmt.Errorf("spot %d not found in parking lot", ticket.Spot.Id)
}

func (p *ParkingLotManager) calculateParkingFee(ticket *domain.ParkingTicket) (*domain.ParkingTicket, error) {
	fmt.Println("Calculating parking fee for ticket:", ticket.Id)
	ParkingTime := ticket.CheckedOut.Sub(ticket.CheckedIn).Seconds()
	if ParkingTime < 0 {
		return nil, fmt.Errorf("invalid time: exit time is before entry time")
	}
	ticket.Amount = int64(ParkingTime * 10)
	return ticket, nil
}
