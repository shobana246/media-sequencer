package services

import (
	"errors"
	"media-sequencer/models"
	"media-sequencer/repositories"
	"time"
)

type SyncService struct {
	syncRepo  *repositories.SyncRepository
	mediaRepo *repositories.MediaRepository
}

func NewSyncService(
	syncRepo *repositories.SyncRepository,
	mediaRepo *repositories.MediaRepository,
) *SyncService {
	return &SyncService{
		syncRepo:  syncRepo,
		mediaRepo: mediaRepo,
	}
}

type SyncStatus struct {
	IsActive         bool          `json:"is_active"`
	Media            *models.Media `json:"media"`
	RemainingSeconds int           `json:"remaining_seconds"`
	StartedAt        *time.Time    `json:"started_at"`
}

func (s *SyncService) TriggerSync(mediaID int, durationSeconds int) error {

	// Step 1 — validate duration
	if durationSeconds <= 0 {
		return errors.New("duration must be greater than 0")
	}

	// Step 2 — check media exists
	media, err := s.mediaRepo.GetMediaByID(mediaID)
	if err != nil {
		return errors.New("media not found")
	}
	if media == nil {
		return errors.New("media not found")
	}

	// Step 3 — check if sync already active
	// if yes complete it first then create new one
	existing, err := s.syncRepo.GetActiveSync()
	if err == nil && existing != nil {
		err = s.syncRepo.CompleteSync(existing.ID)
		if err != nil {
			return err
		}
	}

	// Step 4 — create new sync session
	now := time.Now()
	session := models.SyncSession{
		MediaID:         mediaID,
		StartedAt:       now,
		DurationSeconds: durationSeconds,
		EndedAt:         now.Add(time.Duration(durationSeconds) * time.Second),
		Status:          "active",
		CreatedAt:       now,
	}

	err = s.syncRepo.CreateSync(session)
	if err != nil {
		return err
	}

	return nil
}

func (s *SyncService) GetSyncStatus() (SyncStatus, error) {

	// Step 1 — get active sync from DB
	session, err := s.syncRepo.GetActiveSync()
	if err != nil || session == nil {
		// no active sync — return inactive status
		return SyncStatus{
			IsActive: false,
		}, nil
	}

	// Step 2 — check if sync duration has ended
	now := time.Now()
	if now.After(session.EndedAt) {
		// sync expired — mark as completed in DB
		s.syncRepo.CompleteSync(session.ID)
		return SyncStatus{
			IsActive: false,
		}, nil
	}

	// Step 3 — get media details
	media, err := s.mediaRepo.GetMediaByID(session.MediaID)
	if err != nil {
		return SyncStatus{IsActive: false}, err
	}

	// Step 4 — calculate remaining seconds
	remaining := int(session.EndedAt.Sub(now).Seconds())

	return SyncStatus{
		IsActive:         true,
		Media:            media,
		RemainingSeconds: remaining,
		StartedAt:        &session.StartedAt,
	}, nil
}
