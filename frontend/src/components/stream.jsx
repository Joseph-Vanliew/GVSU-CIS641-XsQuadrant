import React, { useEffect } from "react";
import Chat from "./Chat";
import { connectViewer } from "./viewer"; // Import the viewer utility

const Stream = ({ noStream, streamWebsocketAddr, chatWebsocketAddr, viewerWebsocketAddr }) => {
    useEffect(() => {
        if (!noStream) {
            connectStream();
            connectViewer(viewerWebsocketAddr); // Use the viewer.js function
        }
    }, [noStream, streamWebsocketAddr, viewerWebsocketAddr]);

    const connectStream = () => {
        const peersElement = document.getElementById("peers");
        const chatElement = document.getElementById("chat");
        const noPermElement = document.getElementById("noperm");

        if (peersElement) peersElement.style.display = "block";
        if (chatElement) chatElement.style.display = "flex";
        if (noPermElement) noPermElement.style.display = "none";

        let pc = new RTCPeerConnection({
            iceServers: [
                { urls: "stun:turn.videochat:3478" },
                { urls: "turn:turn.videochat:3478", username: "akhil", credential: "akhil" },
            ],
        });

        pc.ontrack = (event) => {
            if (event.track.kind === "audio") return;

            const col = document.createElement("div");
            col.className = "column is-6 peer";
            const el = document.createElement(event.track.kind);
            el.srcObject = event.streams[0];
            el.setAttribute("controls", "true");
            el.setAttribute("autoplay", "true");
            el.setAttribute("playsinline", "true");

            col.appendChild(el);
            document.getElementById("noonestream").style.display = "none";
            document.getElementById("nocon").style.display = "none";
            document.getElementById("videos").appendChild(col);

            event.streams[0].onremovetrack = () => {
                if (el.parentNode) {
                    el.parentNode.remove();
                }
                if (document.getElementById("videos").childElementCount <= 2) {
                    document.getElementById("noonestream").style.display = "flex";
                }
            };
        };

        const ws = new WebSocket(streamWebsocketAddr);
        pc.onicecandidate = (e) => {
            if (e.candidate) {
                ws.send(JSON.stringify({ event: "candidate", data: JSON.stringify(e.candidate) }));
            }
        };

        ws.onclose = () => {
            console.log("WebSocket closed");
            pc.close();
            pc = null;
            const videos = document.getElementById("videos");
            while (videos.childElementCount > 2) {
                videos.lastChild.remove();
            }
            document.getElementById("noonestream").style.display = "none";
            document.getElementById("nocon").style.display = "flex";
            setTimeout(connectStream, 1000); // Retry connection
        };

        ws.onmessage = (evt) => {
            const msg = JSON.parse(evt.data);
            if (!msg) {
                console.log("Failed to parse message");
                return;
            }

            switch (msg.event) {
                case "offer":
                    const offer = JSON.parse(msg.data);
                    if (!offer) return console.log("Failed to parse offer");
                    pc.setRemoteDescription(offer);
                    pc.createAnswer().then((answer) => {
                        pc.setLocalDescription(answer);
                        ws.send(JSON.stringify({ event: "answer", data: JSON.stringify(answer) }));
                    });
                    break;

                case "candidate":
                    const candidate = JSON.parse(msg.data);
                    if (!candidate) return console.log("Failed to parse candidate");
                    pc.addIceCandidate(candidate);
                    break;

                default:
                    console.log("Unknown event:", msg.event);
            }
        };

        ws.onerror = (evt) => {
            console.error("WebSocket error:", evt.data);
        };
    };

    return (
        <>
            {noStream ? (
                <div id="nostream" className="columns">
                    <div className="column notif">
                        <article className="notification is-danger is-link">
                            There is no stream for the given Stream Link. <br />
                            Please join another stream room.
                        </article>
                    </div>
                </div>
            ) : (
                <>
                    <Chat ChatWebsocketAddr={chatWebsocketAddr} />

                    <div className="viewer">
                        <p className="icon-users" id="viewer-count"></p>
                    </div>

                    <div id="peers">
                        <div className="columns is-multiline" id="videos">
                            <div id="noonestream" className="column notif">
                                <article className="notification is-primary is-light">
                                    Hey! <br />
                                    No streamer in the room. <br />
                                    Please wait for the streamer.
                                </article>
                            </div>
                            <div id="nocon" className="column notif">
                                <article className="notification is-danger">
                                    Connection is closed! <br />
                                    Please refresh the page.
                                </article>
                            </div>
                        </div>
                    </div>
                </>
            )}
        </>
    );
};

export default Stream;