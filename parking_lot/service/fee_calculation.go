package service

import "parkinglot/domain"

type FeeCalculation interface {
	CalculateFee(ticket *domain.ParkingTicket) (int64, error)
}

type NormalFeeCalculationStrategy struct{}

type PassHolderFeeCalculationStrategy struct{}
