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
1. cd to `backend .` directory starting from /app directory
2. go mod tidy
3. go run main.go

or
```sh
go run ./app/xsface/backend/cmd/main.go
```

4. It should be running on port `8080` and serving your frontend when you run it.

**Frontend**
1. cd to `frontend .` starting from /app directory
2. npm install `.`
```shell
npm run dev
```
4. It should be running on port `5173`

**Testing**
1. Visit this page in your browser: `http://localhost:5173/`
2. Create a log-in user profile and sign in to get started!