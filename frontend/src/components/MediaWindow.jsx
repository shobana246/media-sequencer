import React, { useRef, useEffect } from "react";
import { triggerSync } from "../api/api";

function MediaWindow({ window, onSyncTriggered }) {
  const currentMedia = window.current_media;
  const videoRef = useRef(null);

  useEffect(() => {
    if (videoRef.current) {
      videoRef.current.load();
      videoRef.current.play().catch(() => {});
    }
  }, [currentMedia?.url]);

  const handleClick = async () => {
    if (!currentMedia) return;

    let duration = currentMedia.duration_seconds;

    // for video — use actual video duration
    if (currentMedia.type === "video" && videoRef.current) {
      const videoDuration = videoRef.current.duration;
      if (videoDuration && !isNaN(videoDuration)) {
        duration = Math.ceil(videoDuration);
      }
    }

    try {
      await triggerSync(currentMedia.id, duration);
      if (onSyncTriggered) onSyncTriggered();
    } catch (err) {
      console.error("Sync failed:", err);
    }
  };

  const renderMedia = () => {
    if (!currentMedia) {
      return <div className="blank-screen" />;
    }

    if (currentMedia.type === "image") {
      return (
        <img
          src={currentMedia.url}
          alt=""
          style={{ width: "100%", height: "100%", objectFit: "cover" }}
        />
      );
    }

    if (currentMedia.type === "video") {
      return (
        <video
          ref={videoRef}
          autoPlay
          muted
          playsInline
          style={{ width: "100%", height: "100%", objectFit: "cover" }}
          onEnded={(e) => {
            // loop video but track when it ends for sync
            e.target.play();
          }}
        >
          <source src={currentMedia.url} />
        </video>
      );
    }

    if (currentMedia.type === "blank") {
      return <div className="blank-screen" />;
    }
  };

  return (
    <div
      className={`media-window ${window.is_synced ? "synced" : ""}`}
      onClick={handleClick}
      style={{ cursor: "pointer" }}
    >
      <div className="window-screen">{renderMedia()}</div>
    </div>
  );
}

export default MediaWindow;