package repositories

import (
	"database/sql"
	"media-sequencer/models"
)

type PlaylistRepository struct {
	db *sql.DB
}

func NewPlaylistRepository(db *sql.DB) *PlaylistRepository {
	return &PlaylistRepository{
		db: db,
	}
}

func (r *PlaylistRepository) GetPlaylistByWindowID(windowID int) ([]models.Media, error) {
	rows, err := r.db.Query(`
		SELECT 
			m.id,
			m.name,
			m.type,
			m.url,
			m.duration_seconds,
			m.created_at,
			m.updated_at
		FROM playlist_items pi
		JOIN media m ON pi.media_id = m.id
		WHERE pi.window_id = ?
		ORDER BY pi.position ASC
	`, windowID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var playlist []models.Media

	for rows.Next() {
		var media models.Media
		err := rows.Scan(
			&media.ID,
			&media.Name,
			&media.Type,
			&media.URL,
			&media.DurationSeconds,
			&media.CreatedAt,
			&media.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		playlist = append(playlist, media)
	}

	return playlist, nil
}
func (r *PlaylistRepository) GetNextPosition(windowID int) (int, error) {
	var maxPosition int

	err := r.db.QueryRow(`
		SELECT COALESCE(MAX(position), 0)
		FROM playlist_items
		WHERE window_id = ?
	`, windowID).Scan(&maxPosition)

	if err != nil {
		return 0, err
	}

	// next position is max + 1
	return maxPosition + 1, nil
}

func (r *PlaylistRepository) AddToPlaylist(item models.PlaylistItem) error {
	_, err := r.db.Exec(`
		INSERT INTO playlist_items (window_id, media_id, position, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`,
		item.WindowID,
		item.MediaID,
		item.Position,
		item.CreatedAt,
		item.UpdatedAt,
	)
	return err
}
