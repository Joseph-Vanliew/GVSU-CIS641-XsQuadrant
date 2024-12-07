import React, { useState, useEffect, useRef } from "react";

const Chat = ({ ChatWebsocketAddr }) => {
    const [slideOpen, setSlideOpen] = useState(false);
    const [chatWs, setChatWs] = useState(null);
    const [messages, setMessages] = useState([]);
    const msgRef = useRef(null);
    const logRef = useRef(null);

    // Toggle chat visibility
    const slideToggle = () => {
        setSlideOpen((prev) => !prev);
    };


    const appendLog = (message) => {
        setMessages((prevMessages) => [...prevMessages, message]);

        const logElement = logRef.current;
        if (
            logElement &&
            logElement.scrollTop > logElement.scrollHeight - logElement.clientHeight - 1
        ) {
            logElement.scrollTop = logElement.scrollHeight - logElement.clientHeight;
        }
    };

    // Get current time in HH:mm format
    const currentTime = () => {
        const date = new Date();
        let hour = date.getHours();
        let minute = date.getMinutes();
        if (hour < 10) hour = "0" + hour;
        if (minute < 10) minute = "0" + minute;
        return `${hour}:${minute}`;
    };

    // Handle form submission
    const handleSubmit = (e) => {
        e.preventDefault();
        if (!chatWs || !msgRef.current.value) return;
        chatWs.send(msgRef.current.value);
        msgRef.current.value = "";
    };

    // Connect to WebSocket
    const connectChat = () => {
        const ws = new WebSocket(ChatWebsocketAddr);

        ws.onclose = () => {
            console.log("WebSocket has closed");
            setTimeout(connectChat, 1000); // Reconnect after 1 second
        };

        ws.onmessage = (evt) => {
            const newMessages = evt.data.split("\n");
            if (!slideOpen) {
                document.getElementById("chat-alert").style.display = "block";
            }
            newMessages.forEach((message) => {
                appendLog(`${currentTime()} - ${message}`);
            });
        };

        ws.onerror = (evt) => {
            console.error("WebSocket error:", evt.data);
        };

        ws.onopen = () => {
            console.log("WebSocket connected");
        };

        setChatWs(ws);
    };

    // Initialize WebSocket connection on component mount
    useEffect(() => {
        connectChat();
        return () => chatWs?.close(); // Cleanup WebSocket connection on unmount
    }, []);

    return (
        <div id="chat">
            <article className="message chat">
                <div className="message-header" onClick={slideToggle}>
                    <p>Chat</p>
                    <i id="chat-alert" style={{ display: slideOpen ? "none" : "block" }}></i>
                </div>
                <div id="chat-content" style={{ display: slideOpen ? "block" : "none" }}>
                    <div className="body">
                        <div id="log" ref={logRef}>
                            {messages.map((msg, index) => (
                                <div key={index}>{msg}</div>
                            ))}
                        </div>
                    </div>
                    <form id="form" autoComplete="off" onSubmit={handleSubmit}>
                        <div className="field has-addons">
                            <div className="send">
                                <div className="control-input">
                                    <input
                                        className="input"
                                        id="msg"
                                        type="text"
                                        placeholder="Type message..."
                                        ref={msgRef}
                                    />
                                </div>
                                <div className="control">
                                    <input
                                        id="chat-button"
                                        className="button is-info"
                                        type="submit"
                                        value="Send"
                                        disabled={!chatWs || chatWs.readyState !== WebSocket.OPEN}
                                    />
                                </div>
                            </div>
                        </div>
                    </form>
                </div>
            </article>
        </div>
    );
};

export default Chat;