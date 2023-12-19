package entity

import "time"

type Event struct {
	ID          int
	Thumbnail   string
	Title       string
	StartDate   time.Time
	EndDate     time.Time
	Description string
	MaxJoin     uint8
	Price       float64
	IsPublish   bool
	Location    bool
}
