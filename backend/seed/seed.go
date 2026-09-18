package seed

import (
	"database/sql"
	"log"
	"time"
)

func SeedData(db *sql.DB) {

	// check if already seeded
	var count int
	db.QueryRow(`SELECT COUNT(*) FROM windows`).Scan(&count)
	if count > 0 {
		log.Println("Seed data already exists — skipping")
		return
	}

	now := time.Now()

	// ── SEED WINDOWS ──
	windows := []string{
		"Window 1",
		"Window 2",
		"Window 3",
		"Window 4",
	}

	windowIDs := []int{}
	for _, name := range windows {
		result, err := db.Exec(`
			INSERT INTO windows (name, created_at, updated_at)
			VALUES (?, ?, ?)
		`, name, now, now)
		if err != nil {
			log.Fatal("Failed to seed window:", err)
		}
		id, _ := result.LastInsertId()
		windowIDs = append(windowIDs, int(id))
		log.Println("Seeded window:", name)
	}

	// ── SEED MEDIA ──
	type MediaSeed struct {
		Name     string
		Type     string
		URL      string
		Duration int
	}

	medias := []MediaSeed{
		{
			"Brand_Campaign",
			"image",
			"https://picsum.photos/seed/1/1920/1080",
			2700,
		},
		{
			"Product_Demo",
			"video",
			"https://www.w3schools.com/html/mov_bbb.mp4",
			4800,
		},
		{
			"Blank",
			"blank",
			"",
			1800,
		},
		{
			"Promo_Summer",
			"video",
			"https://www.w3schools.com/html/movie.mp4",
			8700,
		},
		{
			"Hero_Banner",
			"image",
			"https://picsum.photos/seed/2/1920/1080",
			2400,
		},
		{
			"Company_Story",
			"video",
			"https://www.w3schools.com/html/mov_bbb.mp4",
			5400,
		},
		{
			"Sale_Poster",
			"image",
			"https://picsum.photos/seed/3/1920/1080",
			3600,
		},
		{
			"New_Collection",
			"video",
			"https://www.w3schools.com/html/movie.mp4",
			7200,
		},
		{
			"Blank2",
			"blank",
			"",
			1800,
		},
		{
			"Event_Promo",
			"image",
			"https://picsum.photos/seed/4/1920/1080",
			3000,
		},
	}

	mediaIDs := []int{}
	for _, m := range medias {
		result, err := db.Exec(`
			INSERT INTO media (name, type, url, duration_seconds, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?, ?)
		`, m.Name, m.Type, m.URL, m.Duration, now, now)
		if err != nil {
			log.Fatal("Failed to seed media:", err)
		}
		id, _ := result.LastInsertId()
		mediaIDs = append(mediaIDs, int(id))
		log.Println("Seeded media:", m.Name)
	}

	// ── SEED PLAYLIST ITEMS ──
	type PlaylistSeed struct {
		WindowIndex int
		MediaIndex  int
		Position    int
	}

	playlists := []PlaylistSeed{
		// Window 1 — total 5h
		{0, 0, 1}, // Brand_Campaign   45min
		{0, 1, 2}, // Product_Demo     1h20min
		{0, 2, 3}, // Blank            30min
		{0, 3, 4}, // Promo_Summer     2h25min

		// Window 2 — total ~5h
		{1, 4, 1}, // Hero_Banner      40min
		{1, 5, 2}, // Company_Story    1h30min
		{1, 6, 3}, // Sale_Poster      1h
		{1, 7, 4}, // New_Collection   2h

		// Window 3 — total ~5h
		{2, 7, 1}, // New_Collection   2h
		{2, 8, 2}, // Blank2           30min
		{2, 9, 3}, // Event_Promo      50min
		{2, 5, 4}, // Company_Story    1h30min

		// Window 4 — total ~5h
		{3, 0, 1}, // Brand_Campaign   45min
		{3, 5, 2}, // Company_Story    1h30min
		{3, 7, 3}, // New_Collection   2h
		{3, 2, 4}, // Blank            30min
		{3, 4, 5}, // Hero_Banner      40min
	}

	for _, p := range playlists {
		_, err := db.Exec(`
			INSERT INTO playlist_items (window_id, media_id, position, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?)
		`,
			windowIDs[p.WindowIndex],
			mediaIDs[p.MediaIndex],
			p.Position,
			now,
			now,
		)
		if err != nil {
			log.Fatal("Failed to seed playlist:", err)
		}
	}

	log.Println("Seed completed successfully ✅")
}
