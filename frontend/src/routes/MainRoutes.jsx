import React from "react";
import UserHome from "../components/pages/UserHome"; // Import the new User Home page
import SignIn from "../components/pages/SignIn";
import SignUp from "../components/pages/SignUp";
import LandingPage from "../components/pages/LandingPage";
import Dashboard from "../components/pages/Dashboard";
import ProtectedRoute from "./ProtectedRoute";

const MainRoutes = [
    {
        path: "/", // Landing Page
        element: <LandingPage />,
    },
    {
        path: "/signin", // Sign In
        element: <SignIn />,
    },
    {
        path: "/signup", // Sign Up
        element: <SignUp />,
    },
    {
        path: "/home", // Protected User Home
        element: (
                <UserHome />
        ),
    },
    {
        path: "/room", // Protected room dash
        element: (
                <Dashboard />
        ),
    },
    // {
    //     path: "/meetings", // Meetings Page
    //     element: (
    //         <ProtectedRoute>
    //             <Meetings />
    //         </ProtectedRoute>
    //     ),
    // },
    // {
    //     path: "/schedule-meeting", // Schedule Meeting Page
    //     element: (
    //         <ProtectedRoute>
    //             <ScheduleMeeting />
    //         </ProtectedRoute>
    //     ),
    // },
];

export default MainRoutes;