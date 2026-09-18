import React, { useState } from "react";
import { triggerSync } from "../api/api";

function SyncControl({ windows, syncStatus }) {
  const [mediaId, setMediaId] = useState("");
  const [message, setMessage] = useState("");
  const [error, setError] = useState("");

  const allMedia = [];
  const seenIds = new Set();
  if (windows) {
    windows.forEach((win) => {
      win.playlist.forEach((media) => {
        if (!seenIds.has(media.id)) {
          seenIds.add(media.id);
          allMedia.push(media);
        }
      });
    });
  }

  const handleSync = async () => {
    if (!mediaId) {
      setError("Please select a media item");
      return;
    }

    const selectedMedia = allMedia.find((m) => m.id === parseInt(mediaId));
    if (!selectedMedia) {
      setError("Media not found");
      return;
    }

    // for video use 30s default since we cant get actual duration here
    // for image and blank use DB duration
    let duration = selectedMedia.duration_seconds;
    if (selectedMedia.type === "video") {
      duration = 30; // video will play fully then sync ends
    }

    try {
      await triggerSync(selectedMedia.id, duration);
      setMessage("Sync triggered");
      setError("");
      setTimeout(() => setMessage(""), 3000);
    } catch (err) {
      setError("Failed to trigger sync");
    }
  };

  return (
    <div className="control-card">
      <div className="control-title">⚡ Sync Control</div>

      {syncStatus?.is_active && (
        <div style={{
          background: "#f59e0b22",
          border: "1px solid #f59e0b55",
          borderRadius: "6px",
          padding: "8px 12px",
          fontSize: "12px",
          color: "#f59e0b",
          marginBottom: "14px"
        }}>
          Sync active — {syncStatus.media?.name} — {syncStatus.remaining_seconds}s remaining
        </div>
      )}

      <div className="form-group">
        <label className="form-label">Select Media</label>
        <select
          className="form-select"
          value={mediaId}
          onChange={(e) => setMediaId(e.target.value)}
        >
          <option value="">-- Select media --</option>
          {allMedia.map((media) => (
            <option key={media.id} value={media.id}>
              {media.name} ({media.type})
            </option>
          ))}
        </select>
      </div>

      <button className="btn btn-sync" onClick={handleSync}>
        Trigger Sync
      </button>

      {message && <div className="success-msg">{message}</div>}
      {error && <div className="error-msg">{error}</div>}
    </div>
  );
}

export default SyncControl;