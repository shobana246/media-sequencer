package services

import (
	"media-sequencer/models"
	"media-sequencer/repositories"
	"time"
)

type WindowService struct {
	windowRepo   *repositories.WindowRepository
	playlistRepo *repositories.PlaylistRepository
	syncRepo     *repositories.SyncRepository
	mediaRepo    *repositories.MediaRepository
}

func NewWindowService(
	windowRepo *repositories.WindowRepository,
	playlistRepo *repositories.PlaylistRepository,
	syncRepo *repositories.SyncRepository,
	mediaRepo *repositories.MediaRepository,
) *WindowService {
	return &WindowService{
		windowRepo:   windowRepo,
		playlistRepo: playlistRepo,
		syncRepo:     syncRepo,
		mediaRepo:    mediaRepo,
	}
}

type WindowState struct {
	Window        models.Window  `json:"window"`
	Playlist      []models.Media `json:"playlist"`
	CurrentMedia  *models.Media  `json:"current_media"`
	CyclePosition int            `json:"cycle_position_seconds"`
	IsSynced      bool           `json:"is_synced"`
}

func (s *WindowService) GetAllWindows() ([]WindowState, error) {
	syncedMedia, err := s.getActiveSyncedMedia()
	if err != nil {
		return nil, err
	}

	windows, err := s.windowRepo.GetAllWindows()
	if err != nil {
		return nil, err
	}

	var result []WindowState

	for _, window := range windows {
		playlist, err := s.playlistRepo.GetPlaylistByWindowID(window.ID)
		if err != nil {
			return nil, err
		}

		if syncedMedia != nil {
			result = append(result, WindowState{
				Window:        window,
				Playlist:      playlist,
				CurrentMedia:  syncedMedia,
				CyclePosition: 0,
				IsSynced:      true,
			})
			continue
		}

		currentMedia, cyclePosition := calculateCurrentMedia(playlist)

		result = append(result, WindowState{
			Window:        window,
			Playlist:      playlist,
			CurrentMedia:  currentMedia,
			CyclePosition: cyclePosition,
			IsSynced:      false,
		})
	}

	return result, nil
}

func (s *WindowService) getActiveSyncedMedia() (*models.Media, error) {
	session, err := s.syncRepo.GetActiveSync()
	if err != nil || session == nil {
		return nil, nil
	}

	if time.Now().After(session.EndedAt) {
		s.syncRepo.CompleteSync(session.ID)
		return nil, nil
	}

	media, err := s.mediaRepo.GetMediaByID(session.MediaID)
	if err != nil {
		return nil, err
	}

	return media, nil
}

func calculateCurrentMedia(playlist []models.Media) (*models.Media, int) {
	if len(playlist) == 0 {
		return nil, 0
	}

	fiveHours := 5 * 60 * 60

	totalDuration := 0
	for _, media := range playlist {
		totalDuration += media.DurationSeconds
	}

	now := time.Now()
	secondsToday := now.Hour()*3600 + now.Minute()*60 + now.Second()
	cyclePosition := secondsToday % fiveHours
	positionInPlaylist := cyclePosition % totalDuration

	elapsed := 0
	for i := range playlist {
		elapsed += playlist[i].DurationSeconds
		if positionInPlaylist < elapsed {
			return &playlist[i], cyclePosition
		}
	}

	return &playlist[0], cyclePosition
}
