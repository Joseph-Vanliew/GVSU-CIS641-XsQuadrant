import React, { useState, useEffect, useRef } from 'react';
import axios from 'axios';

const MemberCaller = (props) => {
  const [pcReceiver, setPcReceiver] = useState(null);
  const receiverVideoRef = useRef(null);

  const meetingID = props.meetingID;
  const peerID = props.peerID;
  const mediaStream = useRef(new MediaStream());

  console.log("**********meetingID:", meetingID);
  console.log("**********peerID:", peerID);

  useEffect(() => {
    const pcReceiverConfig = new RTCPeerConnection({
      iceServers: [{ urls: 'stun:stun.l.google.com:19302' }]
    });

    setPcReceiver(pcReceiverConfig);

    pcReceiverConfig.onicecandidate = (event) => {
      if (event.candidate == null) {
        axios.post(`http://localhost:8080/webrtc/meeting/${meetingID}/peer/${peerID}/isAdmin/false`, {
          Sdp: btoa(JSON.stringify(pcReceiverConfig.localDescription))
        }).then(response => {
          console.log("Setting remote description");
          pcReceiverConfig.setRemoteDescription(new RTCSessionDescription(JSON.parse(atob(response.data.Sdp))));
        });
      }
    };

    pcReceiverConfig.ontrack = (event) => {
      console.log("Inside ontrack");

      if (event.track) {
        mediaStream.current.addTrack(event.track);

        console.log("event.streams[0]:", event.streams[0]);
        
        if (receiverVideoRef.current) {
          console.log("mediaStream.current", mediaStream.current);
          receiverVideoRef.current.srcObject = mediaStream.current;
          // receiverVideoRef.current.play(); 

        }
      }
    };

    return () => {
      pcReceiverConfig.close(); // Cleanup on component unmount
    };
  }, [meetingID, peerID]);

  const startCall = () => {
    if (pcReceiver) {
      // Add transceivers to specify expected media directions
      pcReceiver.addTransceiver('video', { direction: 'sendrecv' });
      pcReceiver.addTransceiver('audio', { direction: 'sendrecv' });

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
        <div className="layer2">
          <p>Member</p>
          <video
            ref={receiverVideoRef}
            style={{ transform: 'rotateY(180deg)' }}
            autoPlay
            width="500"
            height="500"
            controls
          ></video>
        </div>
      </div>
    </div>
  );
};

export default MemberCaller;
