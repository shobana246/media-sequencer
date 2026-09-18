package models

import "time"

type SyncSession struct {
	ID              int       `json:"id"`
	MediaID         int       `json:"media_id"`
	StartedAt       time.Time `json:"started_at"`
	DurationSeconds int       `json:"duration_seconds"`
	EndedAt         time.Time `json:"ended_at"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
}
