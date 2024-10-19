import React, { useState, useEffect, useRef } from 'react';
import axios from 'axios';

const CallComponent = () => {
  const [pcSender, setPcSender] = useState(null);
  const [pcReceiver, setPcReceiver] = useState(null);
  const senderVideoRef = useRef(null);
  const receiverVideoRef = useRef(null);

  // Assuming you use React Router for route parameters
  const meetingId = new URLSearchParams(window.location.search).get('meetingId');
  const peerId = new URLSearchParams(window.location.search).get('peerId');
  const userId = new URLSearchParams(window.location.search).get('userId');



  useEffect(() => {
        const pcSenderConfig = new RTCPeerConnection({
        iceServers: [{ urls: 'stun:stun.l.google.com:19302' }]
        });
        const pcReceiverConfig = new RTCPeerConnection({
        iceServers: [{ urls: 'stun:stun.l.google.com:19302' }]
        });

        setPcSender(pcSenderConfig);
        setPcReceiver(pcReceiverConfig);

        pcSenderConfig.onicecandidate = (event) => {
            if (event.candidate === null) {
                axios.post(`http://localhost:8080/webrtc/sdp/m/${meetingId}/c/${userId}/p/${peerId}/s/true`, {
                Sdp: btoa(JSON.stringify(pcSenderConfig.localDescription))
                }).then((response) => {
                pcSenderConfig.setRemoteDescription(new RTCSessionDescription(JSON.parse(atob(response.data.Sdp))));
                });
            }
        };

        pcReceiverConfig.onicecandidate = (event) => {
        if (event.candidate == null) {            
            axios.post(`http://localhost:8080/webrtc/sdp/m/${meetingId}/c/${userId}/p/${userId}/s/false`, {
            Sdp: btoa(JSON.stringify(pcReceiverConfig.localDescription))
            }).then(response => {

            pcReceiverConfig.setRemoteDescription(new RTCSessionDescription(JSON.parse(atob(response.data.Sdp))));
            });
        }
        };

        pcReceiverConfig.ontrack = (event) => {
        if (receiverVideoRef.current) {
            receiverVideoRef.current.srcObject = event.streams[0];
            receiverVideoRef.current.autoplay = true;
            receiverVideoRef.current.controls = true;
            // receiverVideoRef.current.muted = true;
        }
        };

    }, [meetingId, peerId, userId]);

  const startCall = () => {
    if (pcSender) {
      navigator.mediaDevices.getUserMedia({ video: true, audio: true })
        .then((stream) => {
          if (senderVideoRef.current) {
            senderVideoRef.current.srcObject = stream;
          }

          stream.getTracks().forEach((track) => {
            pcSender.addTrack(track);
          });

          pcSender.createOffer().then((offer) => {
            pcSender.setLocalDescription(offer);
          });
        });

      pcSender.addEventListener('connectionstatechange', () => {
        if (pcSender.connectionState === 'connected') {
          console.log("Hooray! You're connected.");
        }
      });
    }

    if (pcReceiver) {
      pcReceiver.addTransceiver('video', { direction: 'recvonly' });
      pcReceiver.createOffer().then((offer) => {
        pcReceiver.setLocalDescription(offer);
      });
    }
  };

  return (
    <div>
      <button onClick={startCall} className="start-call">
        Start the call!
      </button>
      <div className="container_row">
        <p>Admin</p>
        <video ref={senderVideoRef} autoPlay width="500" height="500" controls muted></video>
        <div className="layer2">
            <p>Member</p>

          <video ref={receiverVideoRef} autoPlay width="500" height="500" controls muted></video>
        </div>
      </div>
    </div>
  );
};

export default CallComponent;
