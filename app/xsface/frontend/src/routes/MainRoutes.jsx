import SignIn from "../pages/SignIn"; // Adjust the path if necessary
import SignUp from "../pages/SignUp";
import CallComponent from "../pages/CallComponent";
import { Navigate } from "react-router-dom";

const MainRoutes = [
    {
        path: "/",
        element: <SignIn />, // Default route pointing to SignIn
        inSearch: "Sign In"
    },
    {
        path: "/signup",
        element: <SignUp />, // Route for SignUp
        inSearch: "Sign Up"
    },
    {
        path: "/call",
        element: <CallComponent />,
        inSearch: "Call Component"
    },
    {
        path: "*",
        element: <Navigate to="/" replace />, // Redirect any unknown paths to SignIn
    }
];

export default MainRoutes;