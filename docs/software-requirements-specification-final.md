# **Overview**
### This is a Software Requirements Specifications document for [XsQuadrant](https://github.com/Joseph-Vanliew/GVSU-CIS641-XsQuadrant).
### Project Website [here](https://joseph-vanliew.github.io/GVSU-CIS641-XsQuadrant/)

## **Software Requirements**
#### This section contains Functional and Non-Functional Requirements for the application.

## 1. MeetingController

###              Functional requirements
1.1 The system shall allow users to create a new meeting by specifying required details (e.g., title, description, time).  
1.2 The system shall provide the functionality to update existing meeting details, such as time or description.  
1.3 Users shall be able to delete a meeting using a unique identifier.  
1.4 The system shall generate and return a unique meeting link upon meeting creation.  
1.5 The system shall list all scheduled meetings for a specific database connection.

###              Non-Functional Requirements
1.1 The system shall ensure meeting data is stored securely with proper encryption.  
1.2 The application shall handle up to 10,000 concurrent meeting requests without performance degradation.  
1.3 Meeting operations (create, update, delete) shall respond within 100ms under normal conditions.  
1.4 The system shall log all meeting-related operations for audit purposes.  
1.5 Meeting creation and updates shall follow data validation rules to ensure consistency.



## 2. UserController

###          Functional requirements
2.1 The system shall allow users to sign up with their details using the SignUp method.  
2.2 The system shall validate user sessions through cookies to ensure authentication.  
2.3 Users shall be able to log in by providing valid credentials (username and password).  
2.4 The application shall allow administrators to manage user accounts.  
2.5 A new UserController instance shall be instantiated when required for operations.

###              Non-Functional Requirements
2.1 The application shall secure user authentication with industry-standard hashing and encryption techniques.  
2.2 The system shall handle up to 1,000 simultaneous logins without slowing down.  
2.3 User session validation shall be completed within 50ms.  
2.4 The application shall comply with GDPR and other relevant data privacy regulations.  
2.5 Error messages shall be clear and user-friendly during login and sign-up processes.



## 3. Client

###          Functional requirements
3.1 The system shall support WebSocket communication for sending and receiving messages.  
3.2 The application shall enable clients to connect to the central Hub.  
3.3 The system shall ensure clients can send messages using a channel (Send).  
3.4 Each client shall have a dedicated process (readPump) to read incoming messages.  
3.5 The application shall allow peer-to-peer connections between clients through PeerConn.

###              Non-Functional Requirements
3.1 The WebSocket connection shall maintain a maximum latency of 100ms for real-time communication.  
3.2 The system shall support up to 5,000 simultaneous WebSocket connections.  
3.3 The application shall handle disconnections gracefully without data loss.  
3.4 The system shall minimize memory usage for client processes to enhance scalability.  
3.5 Communication over WebSocket shall be encrypted using TLS 1.3.


## 4. Room

###          Functional requirements
4.1 The system shall allow the creation of new meeting rooms with unique identifiers (UUID).  
4.2 The application shall retrieve existing rooms or create a new one if it doesn’t exist.  
4.3 Each room shall maintain a list of participants (Peers).  
4.4 The system shall provide functionality to associate a Hub instance with a room.  
4.5 The system shall manage and organize peer connections within a room.

###              Non-Functional Requirements

4.1 The application shall allow a maximum of 500 peers in a single room without performance degradation.  
4.2 Room creation and retrieval operations shall complete within 200ms.  
4.3 The system shall maintain consistency in room data, even during high-concurrency operations.  
4.4 Room operations shall be scalable to support dynamic peer addition/removal.  
4.5 The system shall have a fault-tolerant mechanism to recover room data in case of failures.

## 5. Peer

###          Functional requirements
5.1 The application shall allow peers to add remote tracks (AddTrack) for media sharing.  
5.2 The system shall enable peers to remove local tracks when no longer needed.  
5.3 The system shall maintain a lock mechanism (ListLock) to manage concurrent modifications.  
5.4 Peers shall have the ability to track and manage multiple connections.  
5.5 The application shall handle media stream tracking for peers.

###              Non-Functional Requirements
5.1 The system shall support multiple concurrent media tracks per peer with minimal latency.  
5.2 Peer connections shall be encrypted end-to-end for secure media sharing.  
5.3 The application shall handle media addition/removal seamlessly without interrupting ongoing sessions.  
5.4 Peer media quality shall adjust dynamically based on network conditions.  
5.5 The system shall minimize memory leaks when managing peer connections and tracks.


## 6. Hub

###          Functional requirements
6.1 The system shall facilitate communication between multiple clients connected to the hub.  
6.2 The hub shall broadcast messages to all connected clients.  
6.3 The system shall allow registration of new clients to the hub.  
6.4 A new instance of the hub shall be created (NewHub) when needed.  
6.5 The hub shall manage client sessions and their states efficiently.

###              Non-Functional Requirements
6.1 The hub shall support up to 10,000 connected clients with consistent performance.  
6.2 Message broadcasting latency within the hub shall be under 100ms.  
6.3 The hub shall provide high availability with an uptime of 99.9%.  
6.4 The application shall ensure thread-safe operations for managing hub clients.  
6.5 The hub shall handle spikes in traffic dynamically, maintaining smooth operation.

## 7. Meeting Room

###          Functional requirements
7.1 Any user shall be able to create a meeting room and become that meeting room's Host.  
7.2 The users shall be able to share their screen with others during a meeting provided they have been given permission by that meeting's admin.  
7.3 The Host shall be able to end the meeting at any time for all users in the meeting.  
7.4 The user shall be able to join a meeting by following the meeting url they posses.  
7.5 Users shall be able to apply virtual backgrounds or blur their video background during the meeting.  
7.6 The application shall allow the host to record the conference session and save the recording for future use.  
7.7 Users shall be able to mute and unmute their audio and video during the conference.  
7.8 The application shall display a list of all users currently in the meeting.

###              Non-Functional Requirements
7.1 The system shall support a large number of participants in a single conference without compromising performance.  
7.2 The application shall ensure end-to-end encryption of communication to protect users' privacy.  
7.3 The system shall maintain video quality with minimal latency even under fluctuating network conditions.  
7.4 The application shall work seamlessly across modern web browsers (e.g., Chrome, Firefox, Safari) without requiring additional plugins or installations.  
7.5 The system shall ensure the video and audio streams are smooth, maintaining a consistent frame rate and minimal delay under normal network conditions.  
7.6 The application shall have a user-friendly interface with intuitive controls for users to quickly learn and operate the basic functionalities (e.g., mute/unmute, chat).  
7.7 The application shall be optimized to consume minimal CPU, memory, and bandwidth resources on users' devices to ensure smooth performance, even on lower-end hardware.  
7.8 The system shall maintain an uptime of 99.9% and have mechanisms in place to handle failures gracefully without disrupting ongoing conferences.    
7.9 The system shall maintain video and audio latency below 100ms for users with high-speed internet connections to ensure smooth communication.  
7.10 The system shall support a minimum of 500 users without significant performance degradation.  
7.11 The system shall encrypt all video, audio, and text communication using AES-256 to ensure security and privacy of user data.  

## Change Management Plan

### Description

This section describes the successful implementation of the video chatting application that will rely on a three-phase change management plan to ensure efficient usage,\
integration, and issue resolutions.

---

### Training Plan

When a user first creates an account, an interactive training module will appear, walking them through basic functionalities and features.\
If the video application is implemented into an organization, a few live training sessions will be created in collaboration with the IT team that employees can attend if they choose.\
A small website will also be created that has a searchable help center for self-paced learning and a detailed PDF user manual.\
Feedback from each of these areas will be gathered post-training through surveys. As a result, training materials will be adjusted to address common challenges that users present.

---

### Integration Plan

The objective for this plan is to fit the video application into users’ existing software ecosystems.\
As the application evolves, this could turn into a hub for someone that uses multiple common productivity tools such as Slack or Google Workspace.\
Support for calendar synchronization to streamline meeting scheduling will also be introduced through API compatibility.\
Customizable configuration will allow IT administrators to configure settings for single sign-on via OAuth2 or JWT.\
Throughout this process, cross-platform testing will ensure compatibility with Windows, macOS, and Android devices as well as on major browsers.\
Technical documentation for IT teams to ensure integration and a dedicated support channel will be set up to address potential issues that arise.

---

### Issue Resolution Process

The biggest key performance indicator will be user satisfaction, tracked through identification, tracking, and resolution process.\
Implementing real-time monitoring will help measure any latency, connection, or performance issues.\
Third-party analytics tools such as PowerBI will be used to keep a live look at the raw data.\
For incident reporting, a ticket system will be designed into the platform so users can report any bugs.\
This system will capture the date and time they experienced it automatically, while categorizing and prioritizing the tickets.\
The categories will include critical, moderate, and minimal with a timeframe of 4 hours, 24 hours, and 48 hours respectively.\
To integrate continuous improvement, monthly meetings will include root cause analyses of unresolved bugs and updating any current documentation.\
Regular release updates will also be a part of sustaining a user base and high satisfaction.

## Traceability Document

### Purpose

This document aims to provide a clear traceability matrix linking software requirements to their corresponding design artifacts.

---

### Use Case Diagram Traceability

| Artifact ID |          Artifact Name           |       Requirement ID       |
|:-----------:|:--------------------------------:|:--------------------------:|
|     UC1     |      Meeting Room Use Case       | FR7.1, FR7.2, FR7.4, FR7.8 |
|     UC2     | Creating a Meeting Room Use Case | FR1.1, FR1.4, FR7.6, FR7.7 |
|     UC3     |      Start Meeting Use Case      |    FR7.1, FR7.4, FR7.8     |
|     UC4     |   Manage Permissions Use Case    |        FR7.2, FR7.6        |
|     UC5     |      Join Meeting Use Case       |           FR7.4            |
|     TBD     |               TBD                |    FR7.3, FR7.5, NFR7.7    |

---

### Class Diagram Traceability

|   Artifact Name   |                  Requirement ID                   |
|:-----------------:|:-------------------------------------------------:|
| MeetingController | FR1.1, FR1.2, FR1.3, FR1.4, FR1.5, NFR1.3, NFR1.5 |
|  UserController   |    FR2.1, FR2.2, FR2.3, FR2.5, NFR2.1, NFR2.3     |
|   MeetingModel    |            FR1.1, FR1.4, FR1.5, NFR1.5            |
|      Client       |    FR3.1, FR3.2, FR3.3, FR3.5, NFR3.1, NFR3.5     |
|        Hub        |        FR6.1, FR6.2, FR6.3, FR6.5, NFR6.4         |
|       Room        |    FR4.1, FR4.2, FR4.3, FR4.4, NFR4.2, NFR4.5     |
|       Peer        |        FR5.1, FR5.2, FR5.5, NFR5.2, NFR5.4        |

---

### Activity Diagram Traceability

| Artifact ID |        Artifact Name         |       Requirement ID        |
|:-----------:|:----------------------------:|:---------------------------:|
|     AD1     |    Start Meeting Activity    |     FR7.1, FR7.4, FR7.7     |
|     AD2     |    Join Meeting Activity     |        FR7.4, FR7.8         |
|     AD3     | Create Meeting Room Activity | FR1.1, FR1.4, FR1.5, NFR1.5 |
|     AD4     | Manage Permissions Activity  |    FR7.2, FR7.6, NFR7.1     |
|     AD5     |    Leave Meeting Activity    |            FR7.3            |

---

### Links to Artifacts

Below are links to the referenced artifacts for further details:

* [Use Case Diagrams](../artifacts/Window%20Diagram%20XsQuadrant.drawio.pdf)
* [Class Diagrams](../artifacts/Class%20Diagrams.pdf)
* [Activity Diagrams](../artifacts/XsQuadrant_Activity_Diagrams.drawio.pdf)