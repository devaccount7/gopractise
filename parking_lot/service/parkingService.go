package service

import "parkinglot/domain"

type ParkingService interface{
	AllocateParkingSpot(*domain.Vehicle) (*domain.ParkingSpot,error) // AllocateParkingSpot
	DeAllocateParkingSpot(*domain.ParkingTicket) (*domain.ParkingTicket,error) 
}