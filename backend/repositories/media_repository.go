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
	`, media.Name, media.Type, media.URL, media.DurationSeconds, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *MediaRepository) GetMediaByID(id int) (*models.Media, error) {
	var media models.Media
	err := r.db.QueryRow(`
		SELECT id, name, type, COALESCE(url, '') AS url, duration_seconds, created_at, updated_at 
		FROM media WHERE id = ?
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

func (r *MediaRepository) GetAllMedia() ([]models.Media, error) {
	rows, err := r.db.Query(`
		SELECT id, name, type, COALESCE(url, '') AS url, duration_seconds, created_at, updated_at 
		FROM media
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	medias := make([]models.Media, 0)
	for rows.Next() {
		var m models.Media
		err := rows.Scan(
			&m.ID,
			&m.Name,
			&m.Type,
			&m.URL,
			&m.DurationSeconds,
			&m.CreatedAt,
			&m.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		medias = append(medias, m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return medias, nil
}
