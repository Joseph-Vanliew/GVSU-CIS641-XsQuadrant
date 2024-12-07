import React, { useEffect, useState } from "react";
import { Navigate, Outlet, useLocation } from "react-router-dom";
import { validateSession } from "./ValidateSession.js";

const ProtectedRoute = () => {
    const [isAuthenticated, setIsAuthenticated] = useState(null);
    const location = useLocation();

    useEffect(() => {
        const checkSession = async () => {
            try {
                const user = await validateSession();
                console.log("validateSession response:", user);
                setIsAuthenticated(!!user); // True if user exists, otherwise false
            } catch (error) {
                console.error("Session validation failed:", error);
                setIsAuthenticated(false); // Mark authentication as invalid
            }
        };

        checkSession();
    }, []);

    if (isAuthenticated === null) {
        return <div>Loading...</div>; // Optional loading state
    }

    if (!isAuthenticated) {
        // Redirect to welcome page ("/") for invalid session
        return <Navigate to="/" replace />;
    }

    return <Outlet />;
};

export default ProtectedRoute;