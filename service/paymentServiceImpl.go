package service

import (
	"fmt"
	"parkinglot/domain"
)

type PaymentServiceImpl struct {
	
}

func NewPaymentServiceImpl() *PaymentServiceImpl {
	return &PaymentServiceImpl{}
}

func (p *PaymentServiceImpl) ProcessPayment(ticket *domain.ParkingTicket) (*domain.ParkingTicket, error) {
	fmt.Println("Processing payment for ticket:", ticket.Id)
	fmt.Print("Payment amount:", ticket.Amount, "\n")
	// some external payment gateway logic here or will redirect user to payment page like razorpay or stripe
	ticket.Status = "Paid"
	fmt.Println("Payment processed successfully for ticket:", ticket.Id)
	return ticket, nil

}
