export const connectViewer = (viewerWebsocketAddr, setViewerCount) => {
    const viewerCountElement = document.getElementById("viewer-count");
    const ws = new WebSocket(viewerWebsocketAddr);

    ws.onclose = () => {
        console.log("Viewer WebSocket closed");
        if (viewerCountElement) viewerCountElement.innerHTML = "0";
        setViewerCount && setViewerCount(0); // Optional React state update
        setTimeout(() => connectViewer(viewerWebsocketAddr, setViewerCount), 1000);
    };

    ws.onmessage = (evt) => {
        const data = evt.data;
        if (!isNaN(data)) return;
        if (viewerCountElement) viewerCountElement.innerHTML = data;
        setViewerCount && setViewerCount(parseInt(data, 10)); // Optional React state update
    };

    ws.onerror = (evt) => {
        console.error("Viewer WebSocket error:", evt.data);
    };

    return ws; // Return the WebSocket instance for cleanup if needed
};