import React, { useEffect } from "react";
import Swal from "sweetalert2";
import Chat from "./Chat"; //
import { connectViewer } from "./viewer";

const Room = ({ streamLink, roomWebsocketAddr, chatWebsocketAddr, viewerWebsocketAddr }) => {
    useEffect(() => {

        const copyToClipboard = (text) => {
            if (navigator.clipboard) {
                navigator.clipboard
                    .writeText(text)
                    .then(() => {
                        Swal.fire({
                            position: "top-end",
                            text: "Copied",
                            showConfirmButton: false,
                            timer: 1000,
                            width: "150px",
                        });
                    })
                    .catch((err) => console.error("Copy failed", err));
            }
        };

        // Stream WebRTC connection
        const connectStream = (stream) => {
            document.getElementById("peers").style.display = "block";
            document.getElementById("chat").style.display = "flex";
            document.getElementById("noperm").style.display = "none";

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

                document.getElementById("noone").style.display = "none";
                document.getElementById("nocon").style.display = "none";
                document.getElementById("videos").appendChild(col);

                // Remove track logic
                event.streams[0].onremovetrack = () => {
                    if (el.parentNode) {
                        el.parentNode.remove();
                    }
                    if (document.getElementById("videos").childElementCount <= 3) {
                        document.getElementById("noone").style.display = "grid";
                        document.getElementById("noonein").style.display = "grid";
                    }
                };
            };

            stream.getTracks().forEach((track) => pc.addTrack(track, stream));

            const ws = new WebSocket(roomWebsocketAddr);
            pc.onicecandidate = (e) => {
                if (e.candidate) {
                    ws.send(JSON.stringify({ event: "candidate", data: JSON.stringify(e.candidate) }));
                }
            };

            ws.onmessage = (evt) => {
                const msg = JSON.parse(evt.data);
                if (!msg) return console.log("Failed to parse message");

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

            ws.onclose = () => {
                pc.close();
                pc = null;
                const videos = document.getElementById("videos");
                while (videos.childElementCount > 3) {
                    videos.lastChild.remove();
                }
                document.getElementById("noone").style.display = "grid";
                document.getElementById("nocon").style.display = "flex";
                setTimeout(() => connectStream(stream), 1000); // Retry connection
            };

            ws.onerror = (evt) => {
                console.error("WebSocket error:", evt.data);
            };
        };

        // Establish user media and connections
        navigator.mediaDevices
            .getUserMedia({
                video: { width: { max: 1280 }, height: { max: 720 }, aspectRatio: 4 / 3, frameRate: 30 },
                audio: { sampleSize: 16, channelCount: 2, echoCancellation: true },
            })
            .then((stream) => {
                document.getElementById("localVideo").srcObject = stream;
                connectStream(stream);
            })
            .catch((err) => console.error("Failed to access media devices:", err));

        // Connect viewer WebSocket
        connectViewer(viewerWebsocketAddr);
    }, [roomWebsocketAddr, viewerWebsocketAddr]);

    return (
        <div>
            {/* Viewer Count */}
            <div className="viewer">
                <p className="icon-users" id="viewer-count"></p>
            </div>

            {/* Permission Notification */}
            <div id="noperm" className="columns">
                <div className="column notif">
                    <article className="notification is-link">
                        Camera and microphone permissions are needed to join the room. <br />
                        Otherwise, you can join the{" "}
                        <a href={streamLink}>
                            <strong>stream</strong>
                        </a>{" "}
                        as viewer.
                    </article>
                </div>
            </div>

            {/* Chat Component */}
            <Chat ChatWebsocketAddr={chatWebsocketAddr} />

            {/* Peer Videos */}
            <div id="peers">
                <div className="columns is-multiline" id="videos">
                    <div className="column is-6 peer">
                        <video id="localVideo" className="video-area mirror" autoPlay></video>
                    </div>
                    <div id="noone" className="column is-6 notif">
                        <article id="noonein" className="notification is-primary is-light">
                            <button className="delete"></button>
                            No other streamer is in the room. <br />
                            Share your room link to invite your friends.
                        </article>
                    </div>
                    <div id="nocon" className="column is-6 notif">
                        <article className="notification is-danger">
                            Connection is closed!<br />
                            Please refresh the page.
                        </article>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Room;