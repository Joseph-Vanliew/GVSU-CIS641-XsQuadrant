import React, { useState, useEffect, useRef } from 'react';
import axios from 'axios';
import { useRecoilValue } from 'recoil';
import { videoStateAtom } from '../components/state';

const AdminCaller = (props) => {
  const [pcSender, setPcSender] = useState(null);
  const senderVideoRef = useRef(null);
  const meetingID = props.meetingID;
  const peerID = props.peerID;
  const videoState = useRecoilValue(videoStateAtom);

  useEffect(() => {
    const pcSenderConfig = new RTCPeerConnection({
      iceServers: [{ urls: 'stun:stun.l.google.com:19302' }]
    });
    
    setPcSender(pcSenderConfig);

    pcSenderConfig.onicecandidate = (event) => {
      if (event.candidate === null) {
        axios.post(`http://localhost:8080/webrtc/meeting/${meetingID}/peer/${peerID}/isAdmin/true`, {
          Sdp: btoa(JSON.stringify(pcSenderConfig.localDescription))
        }).then((response) => {
          pcSenderConfig.setRemoteDescription(new RTCSessionDescription(JSON.parse(atob(response.data.Sdp))));
        });
      }
    };

    return () => {
      pcSenderConfig.close(); // Clean up the peer connection on component unmount
    };
  }, [meetingID, peerID]);

  const toggleVideo = () => {
    const stream = senderVideoRef.current.srcObject;
    if (stream) {
      stream.getTracks().forEach((track) => track.stop());
      senderVideoRef.current.srcObject = null;
    }
  };

  const startCall = () => {
    if (pcSender) {
      navigator.mediaDevices.getUserMedia({ video: true, audio: true })
        .then((stream) => {
          console.log("*********admin stream", stream);
          senderVideoRef.current.srcObject = stream;

          stream.getTracks().forEach((track) => {
            console.log("*********admin track", track);
            pcSender.addTrack(track, stream);
          });

          return pcSender.createOffer();
        })
        .then((offer) => {
          pcSender.setLocalDescription(offer);
        })
        .catch((error) => {
          console.error("Error in starting the call:", error);
        });

      pcSender.addEventListener('connectionstatechange', () => {
        if (pcSender.connectionState === 'connected') {
          console.log("Hooray! You're connected.");
        }
      });
    }
  };

  return (
    <div>
      <button onClick={startCall} className="start-call">
        Start the call - Admin!
      </button>
      <button onClick={toggleVideo} className="start-call">
        {videoState}
      </button>
      <div className="container_row">
        <p>Admin</p>
        <video
          style={{ transform: 'rotateY(180deg)' }}
          ref={senderVideoRef}
          autoPlay
          width="500"
          height="500"
          controls
        ></video>
      </div>
    </div>
  );
};

export default AdminCaller;
