import React, { useState } from "react";
import { addMediaToWindow } from "../api/api";

function AddMediaForm({ windows }) {
  const [windowId, setWindowId] = useState("");
  const [name, setName] = useState("");
  const [type, setType] = useState("image");
  const [url, setUrl] = useState("");
  const [duration, setDuration] = useState("");
  const [message, setMessage] = useState("");
  const [error, setError] = useState("");

  const handleAdd = async () => {
    if (!windowId || !name || !duration) {
      setError("Please fill all required fields");
      return;
    }

    if (type !== "blank" && !url) {
      setError("URL is required for image and video");
      return;
    }

    try {
      await addMediaToWindow(parseInt(windowId), {
        name,
        type,
        url: type === "blank" ? "" : url,
        duration_seconds: parseInt(duration),
      });

      setMessage("Media added successfully");
      setError("");

      // reset form
      setName("");
      setUrl("");
      setDuration("");
      setType("image");
      setWindowId("");

      setTimeout(() => setMessage(""), 3000);
    } catch (err) {
      setError("Failed to add media");
      setMessage("");
    }
  };

  return (
    <div className="control-card">
      <div className="control-title">+ Add Media to Window</div>

      <div className="form-group">
        <label className="form-label">Window</label>
        <select
          className="form-select"
          value={windowId}
          onChange={(e) => setWindowId(e.target.value)}
        >
          <option value="">-- Select window --</option>
          {windows &&
            windows.map((win) => (
              <option key={win.window.id} value={win.window.id}>
                {win.window.name}
              </option>
            ))}
        </select>
      </div>

      <div className="form-group">
        <label className="form-label">Media Name</label>
        <input
          className="form-input"
          type="text"
          placeholder="e.g. Summer_Promo"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
      </div>

      <div className="form-group">
        <label className="form-label">Type</label>
        <select
          className="form-select"
          value={type}
          onChange={(e) => setType(e.target.value)}
        >
          <option value="image">Image</option>
          <option value="video">Video</option>
          <option value="blank">Blank</option>
        </select>
      </div>

      {type !== "blank" && (
        <div className="form-group">
          <label className="form-label">URL</label>
          <input
            className="form-input"
            type="text"
            placeholder="https://example.com/media.jpg"
            value={url}
            onChange={(e) => setUrl(e.target.value)}
          />
        </div>
      )}

      <div className="form-group">
        <label className="form-label">Duration (seconds)</label>
        <input
          className="form-input"
          type="number"
          min="1"
          placeholder="e.g. 30"
          value={duration}
          onChange={(e) => setDuration(e.target.value)}
        />
      </div>

      <button className="btn btn-add" onClick={handleAdd}>
        Add Media
      </button>

      {message && <div className="success-msg">{message}</div>}
      {error && <div className="error-msg">{error}</div>}
    </div>
  );
}

export default AddMediaForm;