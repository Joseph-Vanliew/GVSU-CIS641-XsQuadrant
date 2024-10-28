import { useRecoilState } from "recoil";


export function authenticateAndConnect(user) {
    if (user == true) {
        return true
    } else {
        return false
    }
}



// xport function toggleVideo() {

//     localStream.getVideoTracks().forEach(track => track.enabled = !track.enabled);
//     isVideoStopped = !isVideoStopped;
//     isVideoStopped ? seVideoState('Start Video') : seVideoState('Stop Video');
// }e