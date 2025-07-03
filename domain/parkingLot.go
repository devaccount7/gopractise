package domain

import "time"

type ParkingLot struct {
	Id    int64
	Spots []ParkingSpot
}

type ParkingSpot struct {
	Id          int64
	IsAvailable bool
	Type        string // Car/Bike
	Location    string //int64 it will be int eventually, but for now string is fine
}

type ParkingTicket struct {
	Id         int64
	Spot       *ParkingSpot
	CheckedIn  time.Time
	CheckedOut time.Time
	Vehicle    *Vehicle
	Amount     int64
	Status     string // Pending, Paid, Cancelled
	// Should we keep transaction details here or in a separate struct? what about amount  and ticket status?
	// like payment pending, paid, cancelled
	// todo: adding payment amount and status for now will discuss later

}

type Vehicle struct {
	Id                  int64
	Type                string
	RegisterationNumber string
}

type Transaction struct {
	Id             int64
	PaymentGateway string
	Amount         int64
	Status         string
}
