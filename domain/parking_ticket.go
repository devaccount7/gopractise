package domain

import (
	"time"
)

type ParkingTicket struct {
	Id         int64
	Spot       *ParkingSpot
	EntryGate  *Entry
	CheckedIn  time.Time
	CheckedOut time.Time
	Vehicle    *Vehicle
	Amount     int64
	Status     string // Pending, Paid, Cancelled
	ExitGate   *Exit
	// Should we keep transaction details here or in a separate struct? what about amount  and ticket status?
	// like payment pending, paid, cancelled
	// todo: adding payment amount and status for now will discuss later

}

func NewParkingTicket(spot *ParkingSpot, vehicle *Vehicle, entryGate *Entry) *ParkingTicket {
	return &ParkingTicket{
		Id:        0,
		Spot:      spot,
		Vehicle:   vehicle,
		CheckedIn: time.Now(),
		Status:    "Active",
		EntryGate: entryGate,
	}
}

func (ticket *ParkingTicket) CheckOut(exitGate *Exit) *ParkingTicket {
	ticket.ExitGate = exitGate
	ticket.CheckedOut = time.Now()
	ticket.Status = "Paid"
	return ticket
}

// func (ticket *ParkingTicket) createTicket(spot *ParkingSpot, vehicle *Vehicle) *ParkingTicket {
// 	fmt.Println("Ticket created successfully:", ticket.Id)
// 	return ticket
// }
