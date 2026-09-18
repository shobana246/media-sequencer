package repositories

import (
	"database/sql"
	"media-sequencer/models"
)

type WindowRepository struct {
	db *sql.DB
}

func NewWindowRepository(db *sql.DB) *WindowRepository {
	return &WindowRepository{
		db: db,
	}
}
func (r *WindowRepository) GetAllWindows() ([]models.Window, error) {
	rows, err := r.db.Query(`
    SELECT id, name, COALESCE(url, '') AS url, created_at, updated_at FROM windows
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var windows []models.Window

	for rows.Next() {
		var window models.Window

		err := rows.Scan(
			&window.ID,
			&window.Name,
			&window.CreatedAt,
			&window.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		windows = append(windows, window)
	}

	return windows, nil
}

func (r *WindowRepository) GetWindowByID(id int) (*models.Window, error) {
	var window models.Window

	err := r.db.QueryRow(`
		SSELECT id, name, COALESCE(url, '') AS url, created_at, updated_at FROM windows WHERE id = ?
	`, id).Scan(
		&window.ID,
		&window.Name,
		&window.CreatedAt,
		&window.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &window, nil
}
