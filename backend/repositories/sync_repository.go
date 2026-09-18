package repositories

import (
	"database/sql"
	"media-sequencer/models"
)

type SyncRepository struct {
	db *sql.DB
}

func NewSyncRepository(db *sql.DB) *SyncRepository {
	return &SyncRepository{
		db: db,
	}
}

func (r *SyncRepository) CreateSync(session models.SyncSession) error {
	_, err := r.db.Exec(`
		INSERT INTO sync_sessions 
		(media_id, started_at, duration_seconds, ended_at, status, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`,
		session.MediaID,
		session.StartedAt,
		session.DurationSeconds,
		session.EndedAt,
		session.Status,
		session.CreatedAt,
	)
	return err
}

func (r *SyncRepository) GetActiveSync() (*models.SyncSession, error) {
	var session models.SyncSession

	err := r.db.QueryRow(`
		SELECT id, media_id, started_at, duration_seconds, ended_at, status, created_at
		FROM sync_sessions
		WHERE status = 'active'
		ORDER BY created_at DESC
		LIMIT 1
	`).Scan(
		&session.ID,
		&session.MediaID,
		&session.StartedAt,
		&session.DurationSeconds,
		&session.EndedAt,
		&session.Status,
		&session.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *SyncRepository) CompleteSync(id int) error {
	_, err := r.db.Exec(`
		UPDATE sync_sessions
		SET status = 'completed'
		WHERE id = ?
	`, id)
	return err
}
