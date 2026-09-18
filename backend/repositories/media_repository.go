package repositories

import (
	"database/sql"
	"media-sequencer/models"
	"time"
)

type MediaRepository struct {
	db *sql.DB
}

func NewMediaRepository(db *sql.DB) *MediaRepository {
	return &MediaRepository{
		db: db,
	}
}

func (r *MediaRepository) CreateMedia(media models.Media) (int, error) {
	result, err := r.db.Exec(`
		INSERT INTO media (name, type, url, duration_seconds, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`,
		media.Name,
		media.Type,
		media.URL,
		media.DurationSeconds,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return 0, err
	}

	// get the inserted media ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *MediaRepository) GetMediaByID(id int) (*models.Media, error) {
	var media models.Media

	err := r.db.QueryRow(`
		SELECT id, name, type, url, duration_seconds, created_at, updated_at
		FROM media
		WHERE id = ?
	`, id).Scan(
		&media.ID,
		&media.Name,
		&media.Type,
		&media.URL,
		&media.DurationSeconds,
		&media.CreatedAt,
		&media.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &media, nil
}
