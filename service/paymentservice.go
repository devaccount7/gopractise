package service

import "parkinglot/domain"

type PaymentService interface {
	// ProcessPayment processes the payment for a parking ticket.
	ProcessPayment(*domain.ParkingTicket) (*domain.ParkingTicket, error)
}
