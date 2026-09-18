package models

import "time"

type PlaylistItem struct {
	ID        int       `json:"id"`
	WindowID  int       `json:"window_id"`
	MediaID   int       `json:"media_id"`
	Position  int       `json:"position"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
