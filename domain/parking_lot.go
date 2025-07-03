package domain

import "sync"

type ParkingLot struct {
	Id      int64
	levels  []ParkingLevel
	entries []Entry
	exits   []Exit

	lock sync.RWMutex
}
