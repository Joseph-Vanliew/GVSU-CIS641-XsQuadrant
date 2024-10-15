# **Overview**
### This is a Software Requirements Specifications document for [XsQuadrant](https://github.com/Joseph-Vanliew/GVSU-CIS641-XsQuadrant)


## **Functional Requirements**

### 1. Meeting Room
    1.1 Any user shall be able to create a meeting room and become that meeting room's Host.
    1.2 The users shall be able to share their screen with others during a meeting provided they have been given permission by that meeting's admin.
    1.3 The Host shall be able to end the meeting at any time for all users in the meeting.
    1.4 The user shall be able to join a meeting by following the meeting url they posses.
    1.5 Users shall be able to apply virtual backgrounds or blur their video background during the meeting.
    1.6 The application shall allow the host to record the conference session and save the recording for future use.
    1.7 Users shall be able to mute and unmute their audio and video during the conference.
    1.8 The application shall display a list of all users currently in the meeting.
### 2. Log in
    2.1 The Application shall require a user to authenticate with OAuth when creating an account.
    
## **Non-Functional Requirements**

### 1. Meeting Room 
    1.1 The system shall support a large number of participants in a single conference without compromising performance.
    1.2 The application shall ensure end-to-end encryption of communication to protect users' privacy.
    1.3 The system shall maintain video quality with minimal latency even under fluctuating network conditions.
    1.4 The application shall work seamlessly across modern web browsers (e.g., Chrome, Firefox, Safari) without requiring additional plugins or installations.
    1.5 The system shall ensure the video and audio streams are smooth, maintaining a consistent frame rate and minimal delay under normal network conditions.
    1.6 The application shall have a user-friendly interface with intuitive controls for users to quickly learn and operate the basic functionalities (e.g., mute/unmute, chat).
    1.7 The application shall be optimized to consume minimal CPU, memory, and bandwidth resources on users' devices to ensure smooth performance, even on lower-end hardware.
    1.8 The system shall maintain an uptime of 99.9% and have mechanisms in place to handle failures gracefully without disrupting ongoing conferences.
    1.9 The system shall maintain video and audio latency below 100ms for users with high-speed internet connections to ensure smooth communication.
    1.10 The system shall support a minimum of 500 users without significant performance degradation.
    1.11 The system shall encrypt all video, audio, and text communication using AES-256 to ensure security and privacy of user data.

