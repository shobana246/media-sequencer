import React, { useState, useEffect } from "react";
import { getWindows, getSyncStatus } from "./api/api";
import WindowGrid from "./components/WindowGrid";
import SyncControl from "./components/SyncControl";
import AddMediaForm from "./components/AddMediaForm";
import "./index.css";

function App() {
  const [windows, setWindows] = useState([]);
  const [syncStatus, setSyncStatus] = useState(null);

  const fetchData = async () => {
    try {
      const [windowsData, syncData] = await Promise.all([
        getWindows(),
        getSyncStatus(),
      ]);
      setWindows(windowsData || []);
      setSyncStatus(syncData);
    } catch (err) {
      console.error("Failed to fetch:", err);
    }
  };

  useEffect(() => {
    fetchData();
    const interval = setInterval(fetchData, 3000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="app">
      <div className="app-header">
        <div className="app-title">Media Sequencer</div>
        {syncStatus?.is_active && (
          <div className="sync-badge">
            ⚡ SYNC — {syncStatus.media?.name} — {syncStatus.remaining_seconds}s
          </div>
        )}
      </div>

      <WindowGrid windows={windows} onSyncTriggered={fetchData} />

      <div className="controls">
        <SyncControl windows={windows} syncStatus={syncStatus} />
        <AddMediaForm windows={windows} />
      </div>
    </div>
  );
}

export default App;