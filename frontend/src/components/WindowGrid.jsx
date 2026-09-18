import React from "react";
import MediaWindow from "./MediaWindow";

function WindowGrid({ windows, onSyncTriggered }) {
  if (!windows || windows.length === 0) {
    return (
      <div style={{ color: "#444", textAlign: "center", padding: "80px 0" }}>
        Connecting...
      </div>
    );
  }

  return (
    <div className="window-grid">
      {windows.map((window) => (
        <MediaWindow
          key={window.window.id}
          window={window}
          onSyncTriggered={onSyncTriggered}
        />
      ))}
    </div>
  );
}

export default WindowGrid;