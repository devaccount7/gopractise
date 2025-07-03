package domain

import "sync"

type ParkingLevel struct {
	Id     int64
	Spots  []ParkingSpot
	isFull bool

	lock sync.RWMutex
}
