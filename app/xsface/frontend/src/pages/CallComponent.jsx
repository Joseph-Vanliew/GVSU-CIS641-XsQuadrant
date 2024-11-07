import React, {} from 'react';
import { useSearchParams } from 'react-router-dom';
import MemberCaller from './MemberCaller';
import AdminCaller from './AdminCaller';

const CallComponent = () => {
  const [searchParams] = useSearchParams()
 
  const meetingID = searchParams.get('meetingID');
  const peerID = searchParams.get('peerID');
  const isAdmin = searchParams.get('isAdmin');

  return (
    
        isAdmin === "true" ? <AdminCaller meetingID={meetingID} peerID={peerID}/> : <MemberCaller meetingID={meetingID} peerID={peerID}/>
    
  );
};

export default CallComponent;
