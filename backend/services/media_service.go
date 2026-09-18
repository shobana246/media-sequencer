package services

import (
	"errors"
	"media-sequencer/models"
	"media-sequencer/repositories"
	"time"
)

type MediaService struct {
	mediaRepo    *repositories.MediaRepository
	windowRepo   *repositories.WindowRepository
	playlistRepo *repositories.PlaylistRepository
}

func NewMediaService(
	mediaRepo *repositories.MediaRepository,
	windowRepo *repositories.WindowRepository,
	playlistRepo *repositories.PlaylistRepository,
) *MediaService {
	return &MediaService{
		mediaRepo:    mediaRepo,
		windowRepo:   windowRepo,
		playlistRepo: playlistRepo,
	}
}

func (s *MediaService) AddMediaToWindow(windowID int, name string, mediaType string, url string, durationSeconds int) error {

	// Step 1 — check window exists
	_, err := s.windowRepo.GetWindowByID(windowID)
	if err != nil {
		return errors.New("window not found")
	}

	// Step 2 — validate media type
	if mediaType != "image" && mediaType != "video" && mediaType != "blank" {
		return errors.New("invalid media type — must be image, video or blank")
	}

	// Step 3 — validate duration
	if durationSeconds <= 0 {
		return errors.New("duration must be greater than 0")
	}

	// Step 4 — create media
	media := models.Media{
		Name:            name,
		Type:            mediaType,
		URL:             url,
		DurationSeconds: durationSeconds,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	// Step 5 — save media to DB and get its ID
	mediaID, err := s.mediaRepo.CreateMedia(media)
	if err != nil {
		return err
	}

	// Step 6 — get next position in playlist
	nextPosition, err := s.playlistRepo.GetNextPosition(windowID)
	if err != nil {
		return err
	}

	// Step 7 — add to playlist
	playlistItem := models.PlaylistItem{
		WindowID:  windowID,
		MediaID:   mediaID,
		Position:  nextPosition,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.playlistRepo.AddToPlaylist(playlistItem)
	if err != nil {
		return err
	}

	return nil
}
