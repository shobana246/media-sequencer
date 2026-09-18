const BASE_URL = process.env.REACT_APP_API_URL;

export const getWindows = async () => {
  const response = await fetch(`${BASE_URL}/api/windows`);
  const data = await response.json();
  return data.data;
};

export const triggerSync = async (mediaId, durationSeconds) => {
  const response = await fetch(`${BASE_URL}/api/sync`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      media_id: mediaId,
      duration_seconds: durationSeconds,
    }),
  });
  return response.json();
};

export const getSyncStatus = async () => {
  const response = await fetch(`${BASE_URL}/api/sync/status`);
  const data = await response.json();
  return data.data;
};

export const addMediaToWindow = async (windowId, media) => {
  const response = await fetch(`${BASE_URL}/api/windows/${windowId}/media`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(media),
  });
  return response.json();
};