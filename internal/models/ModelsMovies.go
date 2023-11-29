package models

import "time"

type Movies struct {
	ID          int    `json:"id"`
	Title       string    `json:"title"`
	ReleaseDate time.Time    `json:"release_date"`
	Runtime     int       `json:"runtime"`
	MPAARATING  string    `json:"m_rating"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"-"`
	UdatedAt    time.Time `json:"-"`
}
