---
layout: default
title: "GVSU CIS641 XsQuadrant"
---

# <ins>Project Overview</ins>

We are developing a video chatting application aimed at enhancing real-time communication for personal and professional use. Our platform will provide smooth, high-quality video and audio interaction together with essential features like instant messaging, screen sharing, and user identification, in response to the growing trend of distant work and virtual meetings.

The objective is to create an interface that is easy to use while prioritizing security, dependability, and performance. Our goal is to make sure users have a seamless online application experience.

A scalable architecture will be used in the construction of our program, guaranteeing low latency and real-time support for numerous users. Delivering a solution that can benefit professionals and casual users alike in a variety of contexts, such as online meetings, virtual gatherings, or interacting with loved ones, is the ultimate objective.

### Technologies

- <img src="{{ '/assets/favicons/reactjs.svg' | relative_url }}" alt="React" style="width: 20px; height: 20px;"> Frontend: React.js for building a responsive and dynamic user interface.
- <img src="{{ '/assets/favicons/golang.svg' | relative_url }}" alt="Golang" style="width: 20px; height: 20px;"> Backend: Golang to handle server-side logic, ensuring performance and scalability.
- <img src="{{ '/assets/favicons/webrtc.svg' | relative_url }}" alt="WebRTC" style="width: 20px; height: 20px;"> WebRTC: For real-time communication, facilitating peer-to-peer video and audio connections.
- <img src="{{ '/assets/favicons/power-socket-us.svg' | relative_url }}" alt="WebSockets" style="width: 20px; height: 20px;"> Socket.io or WebSockets: For real-time messaging and maintaining open communication between users.
- <img src="{{ '/assets/favicons/postgresql.svg' | relative_url }}" alt="PostgreSQL" style="width: 20px; height: 20px;"> PostgreSQL: As the database solution for user data, chat history, and video session metadata.
- <img src="{{ '/assets/favicons/google-cloud.svg' | relative_url }}" alt="Google Cloud" style="width: 20px; height: 20px;"> Cloud services: AWS or Google Cloud for hosting, load balancing, and scaling the infrastructure.
- <img src="{{ '/assets/favicons/oauth.svg' | relative_url }}" alt="OAuth" style="width: 20px; height: 20px;"> Authentication: OAuth2 or JWT for secure login and user verification.
- <img src="{{ '/assets/favicons/github-mark.svg' | relative_url }}" alt="GitHub" style="width: 20px; height: 20px;"> CI/CD tools: Github actions for continuous integration and deployment.

### Approach

An incremental approach to development, starting with the core functionality of the video chat and expanding to additional features in subsequent iterations. The backend server will be configured in Golang initially, and API endpoints for user, session, and real-time communication management will be constructed. The frontend interface will be created in React concurrently, with an initial emphasis on a simple user experience for initiating video conversations.

After WebRTC-based video chatting is operational, we may incorporate the real-time messaging system and make sure it is in sync with video sessions.

We'll keep an eye out for performance problems to guarantee low latency and seamless video streaming.


### Estimated Timeline

- Week 1-2: Set up project structure and develop basic backend API with user management.
- Week 3-4: Build the frontend interface for video chat and integrate WebRTC for video and audio streaming.
- Week 5-6: Implement real-time messaging and sync with video sessions.
- Week 7-8: Add user authentication and session management, including testing for security.
- Week 9-10: Develop group chat and screen-sharing features, followed by thorough testing.
- Week 11-12: Deploy the application, conduct user testing, and optimize for performance and scaling.

### Anticipated Challenges to Tackle

- Latency issues: Ensuring low-latency video and audio, especially for users with slower internet connections, may require optimization of streaming protocols and adaptive video quality.
- Scalability: As the number of users increases, managing server load and ensuring smooth performance will be challenging. We may need to implement load balancing and efficient resource allocation on cloud infrastructure.
- Cross-browser compatibility: WebRTC can behave differently across browsers, so ensuring consistent performance across platforms will require extensive testing.
- Security: Safeguarding user data and communication with encryption, along with preventing unauthorized access, will be critical.





