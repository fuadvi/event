package model

import "time"

type EventCreate struct {
	Thumbnail   string    `json:"thumbnail" validate:"required"`
	Title       string    `json:"title" validate:"required,max=100"`
	StartDate   time.Time `json:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" validate:"required"`
	Description string    `json:"description" validate:"required"`
	MaxJoin     uint8     `json:"max_join" validate:"required,number"`
	Price       float64   `json:"price" validate:"required,numeric"`
	IsPublish   bool      `json:"is_publish" validate:"required,boolean"`
	Location    bool      `json:"location" validate:"required,boolean"`
}

type EventUpdate struct {
	Thumbnail   string    `json:"thumbnail" validate:"required"`
	Title       string    `json:"title" validate:"required,max=100"`
	StartDate   time.Time `json:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" validate:"required"`
	Description string    `json:"description" validate:"required"`
	MaxJoin     uint8     `json:"max_join" validate:"required,number"`
	Price       float64   `json:"price" validate:"required,numeric"`
	IsPublish   bool      `json:"is_publish" validate:"required,boolean"`
	Location    bool      `json:"location" validate:"boolean"`
}

type EventResponse struct {
	ID          int       `json:"id,omitempty"`
	Thumbnail   string    `json:"thumbnail,omitempty"`
	Title       string    `json:"title,omitempty"`
	StartDate   time.Time `json:"start_date,omitempty"`
	EndDate     time.Time `json:"end_date,omitempty"`
	Description string    `json:"description,omitempty"`
	MaxJoin     uint8     `json:"max_join,omitempty"`
	Price       float64   `json:"price,omitempty"`
	IsPublish   bool      `json:"is_publish,omitempty"`
	Location    bool      `json:"location"`
}
