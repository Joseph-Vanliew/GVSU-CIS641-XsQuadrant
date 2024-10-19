# XsQuadrant

A video conferencing web application that is set up similarly to popular video conferencing
platforms of today. The application will allow users to intialize a room which other users may join 
with the appropriate link access created by the room's initiator and admin and chat face to face
with video and audio streaming to all users within the room. 

## Team Members and Roles

* [Joseph Vanliew](https://github.com/Joseph-Vanliew/641-HW2-Vanliew) (Backend Architecture, Frontend Support)
* [Joachim Kuleafenu](https://github.com/kuleafenu/641-HW2-Kuleafenu) (Backend Developer, Frontend Support)
* [Brenden Koneval](https://github.com/konevalb/641-HW2-Koneval) (Frontend Developer, Backend Support)

## Prerequisites

## Run Instructions
**Backend**
1. cd to `backend .` directory
2. go mod tidy
3. go run main.go
4. It should be running on port `8080`

**Frontend**
1. cd to `frontend .`
2. npm install `.`
3. npm run start
4. It should be running on port `3000`

**Testing**
1. Visit this page in your browser: `http://localhost:3000/?meetingId=07927fc8-af0a-11ea-b338-064f26a5f90a&userId=alice&peerId=bob`
2. Click on `Start the call!` button