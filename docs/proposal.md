Team name:
# XsQuadrant

Team members:
- Joseph Vanliew
- Brenden Kenoval
- Joachim Kuleafenu

# Introduction

(In 2-4 paragraphs, describe your project concept)
We are developing a video chatting application aimed at enhancing real-time communication for personal and professional use.Our platform will provide smooth, high-quality video and audio interaction together with essential features like instant messaging, screen sharing, and user identification, in response to the growing trend of distant work and virtual meetings.

The objective is to create an interface that is easy to use while prioritizing security, dependability, and performance. Our goal is to make sure users have a seamless online application experience.

A scalable architecture will be used in the construction of our program, guaranteeing low latency and real-time support for numerous users. Delivering a solution that can benefit professionals and casual users alike in a variety of contexts, such as online meetings, virtual gatherings, or interacting with loved ones, is the ultimate objective.

# Anticipated Technologies

- Frontend: React.js for building a responsive and dynamic user interface.
- Backend: Golang to handle server-side logic, ensuring performance and scalability.
- WebRTC: For real-time communication, facilitating peer-to-peer video and audio connections.
- Socket.io or WebSockets: For real-time messaging and maintaining open communication between users.
- PostgreSQL: As the database solution for user data, chat history, and video session metadata.
- Cloud services: AWS or Google Cloud for hosting, load balancing, and scaling the infrastructure.
- Authentication: OAuth2 or JWT for secure login and user verification.
- CI/CD tools: Github actions for continuous integration and deployment.

# Method/Approach

We plan to take an incremental approach to development, starting with the core functionality of the video chat and expanding to additional features in subsequent iterations. The backend server will be configured in Golang initially, and API endpoints for user, session, and real-time communication management will be constructed. The frontend interface will be created in React concurrently, with an initial emphasis on a simple user experience for initiating video conversations.

After WebRTC-based video chatting is operational, we may incorporate the real-time messaging system and make sure it is in sync with video sessions.

We'll keep an eye out for performance problems to guarantee low latency and seamless video streaming.


# Estimated Timeline

- Week 1-2: Set up project structure and develop basic backend API with user management.
- Week 3-4: Build the frontend interface for video chat and integrate WebRTC for video and audio streaming.
- Week 5-6: Implement real-time messaging and sync with video sessions.
- Week 7-8: Add user authentication and session management, including testing for security.
- Week 9-10: Develop group chat and screen-sharing features, followed by thorough testing.
- Week 11-12: Deploy the application, conduct user testing, and optimize for performance and scaling.

# Anticipated Problems

- Latency issues: Ensuring low-latency video and audio, especially for users with slower internet connections, may require optimization of streaming protocols and adaptive video quality.
- Scalability: As the number of users increases, managing server load and ensuring smooth performance will be challenging. We may need to implement load balancing and efficient resource allocation on cloud infrastructure.
- Cross-browser compatibility: WebRTC can behave differently across browsers, so ensuring consistent performance across platforms will require extensive testing.
- Security: Safeguarding user data and communication with encryption, along with preventing unauthorized access, will be critical.



