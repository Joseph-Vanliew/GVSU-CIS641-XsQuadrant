Candidate: ICE candidates represents various ways a device can communicate with another peer, including different IP addresses, ports, and protocols available on the device.
ICE Candidate Exchange
When two peers want to establish a connection using WebRTC:

Each peer gathers its ICE candidates, including host, server-reflexive, and relay candidates.
These candidates are exchanged between peers using a signaling server (through WebSocket, HTTP, etc.).
Peers try the different candidates (a process called ICE connectivity checks) to determine the best path for communication.
Once a successful connection is found, the peers use that candidate to establish the peer-to-peer connection.
