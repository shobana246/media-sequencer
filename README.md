# Multi-Window Media Sequencer with Sync Playback

A full-stack media sequencing application that manages multiple display windows, each with its own media playlist and continuous playback cycle.

The application consists of a React frontend, a Golang backend, and a MySQL database for persistent storage.

## Tech Stack

### Frontend

* React
* JavaScript
* HTML/CSS

### Backend

* Golang
* Gin
* REST APIs

### Database

* MySQL

## Project Structure

```text
media-sequencer/
├── backend/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── config/
│   │   └── database.go
│   ├── handlers/
│   │   ├── media_handler.go
│   │   ├── sync_handler.go
│   │   └── window_handler.go
│   ├── models/
│   │   ├── media.go
│   │   ├── playlist_item.go
│   │   ├── sync_session.go
│   │   └── window.go
│   ├── repositories/
│   │   ├── media_repository.go
│   │   ├── playlist_repository.go
│   │   ├── sync_repository.go
│   │   └── window_repository.go
│   ├── routes/
│   │   └── routes.go
│   ├── seed/
│   │   └── seed.go
│   ├── services/
│   │   ├── media_service.go
│   │   ├── sync_service.go
│   │   └── window_service.go
│   ├── .env.example
│   ├── go.mod
│   └── go.sum
│
├── frontend/
│   ├── public/
│   ├── src/
│   │   ├── api/
│   │   ├── components/
│   │   ├── App.jsx
│   │   ├── index.css
│   │   └── index.js
│   ├── package.json
│   └── package-lock.json
│
└── README.md
```

## Architecture

The backend follows a layered structure:

```text
HTTP Request
     ↓
Handler
     ↓
Service
     ↓
Repository
     ↓
MySQL
```

### Handlers

Receive HTTP requests, validate request data, call the appropriate service, and return HTTP responses.

### Services

Contain the application's business logic, including playlist updates, media validation, playback calculation, and synchronization.

### Repositories

Handle database operations and keep SQL/database access separate from the business logic.

### Models

Represent the application's database and API data structures.

## Core Features

### Multiple Media Windows

Each display window has its own configured playlist.

The backend stores the relationship between windows and playlist items in the database.

### Continuous 5-Hour Cycle

Each window operates within a 5-hour playback cycle.

The configured playlist is continuously looped within that cycle.

If the total playlist duration is shorter than five hours, the playlist repeats rather than leaving the remaining time blank.

### Supported Media Types

The application supports:

* Image
* Video
* Blank

Blank is treated as a configured playlist item. It is not automatically inserted into the playback cycle.

### Dynamic Media Addition

Media can be added to an existing window through the backend API.

The new media is:

1. Validated
2. Stored in the media table
3. Assigned the next playlist position
4. Added to the selected window's playlist

### Synchronized Playback

A synchronization request can temporarily synchronize all windows to the same media.

The backend creates a synchronization session containing:

* Media ID
* Start time
* Duration
* End time
* Status

While synchronization is active, the frontend uses the synchronization status to display the synchronized media.

After the synchronization duration ends, the synchronization session is completed and the windows continue their normal configured playback.

## API Documentation

Base URL:

```text
http://localhost:8080
```

### Get All Windows

```http
GET /api/windows
```

Returns all configured windows along with their playlists, current media, and cycle position.

### Add Media to a Window

```http
POST /api/windows/:id/media
```

Example request:

```json
{
  "name": "M13",
  "type": "image",
  "url": "https://example.com/image.jpg",
  "duration_seconds": 30
}
```

### Trigger Synchronization

```http
POST /api/sync
```

Example request:

```json
{
  "media_id": 13,
  "duration_seconds": 30
}
```

### Get Synchronization Status

```http
GET /api/sync/status
```

Returns whether a synchronization session is currently active and, when active, the synchronized media and remaining duration.

## Database

The application uses MySQL for persistent storage.

The main entities are:

* Windows
* Media
* Playlist Items
* Sync Sessions

The database connection is configured using environment variables.

### Environment Variables

Create a `.env` file inside `backend/`:

```env
DB_USER=root
DB_PASSWORD=your_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=media_sequencer
```

A `.env.example` file is included in the repository as a configuration template.

The actual `.env` file is intentionally excluded from Git.

## Running the Backend

From the `backend` directory:

```bash
go mod download
go run cmd/server/main.go
```

The backend runs on:

```text
http://localhost:8080
```

## Running the Frontend

From the `frontend` directory:

```bash
npm install
npm start
```

The React application runs on:

```text
http://localhost:3000
```

The frontend API configuration is provided through the environment configuration.

## Seed Data

The backend contains seed data for example windows, media, playlists, and related assignment data.

This provides initial data for testing the media sequencing and synchronization behavior.

## Playback Flow

The normal playback flow is:

```text
Window
   ↓
Load Playlist
   ↓
Calculate Position in 5-Hour Cycle
   ↓
Determine Current Media
   ↓
Display Media
   ↓
Continue Playlist Loop
```

When synchronization is triggered:

```text
User Triggers Sync
        ↓
Backend Creates Sync Session
        ↓
Frontend Checks Sync Status
        ↓
All Windows Display Sync Media
        ↓
Sync Duration Ends
        ↓
Normal Window Playback Continues
```

## Assumptions

* Each window has its own configured playlist.
* The playback cycle is treated as a 5-hour cycle.
* Playlist items repeat when the configured playlist duration is shorter than the cycle.
* Blank playback occurs only when a blank media item is explicitly configured.
* Synchronization temporarily overrides normal playback for the synchronization duration.
* After synchronization ends, each window resumes its configured playlist.
* Media and playlist configuration are persisted in MySQL.
* The frontend and backend are deployed separately when using production deployment.

## Deployment

The application consists of two separately deployable parts:

* React frontend
* Golang backend

For production deployment, the frontend must point to the deployed backend API URL instead of the local development URL.

The backend requires access to the configured MySQL database and the following environment variables:

```text
DB_USER
DB_PASSWORD
DB_HOST
DB_PORT
DB_NAME
```

Production deployment URLs can be added here after deployment.

## Assignment Notes

The implementation prioritizes the backend requirements, including persistent storage, playlist management, playback calculation, dynamic media addition, and synchronization.

The frontend provides the display windows and playback interface required to interact with the backend.
