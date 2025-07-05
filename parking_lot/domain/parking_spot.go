package domain

import "sync"

type ParkingSpot struct {
	Id                int64
	IsAvailable       bool
	Type              string // Car/Bike
	Location          string //int64 it will be int eventually, but for now string is fine
	DistanceToEntries map[int64]int64
	DistanceToExits   map[int64]int64

	lock sync.RWMutex
}
