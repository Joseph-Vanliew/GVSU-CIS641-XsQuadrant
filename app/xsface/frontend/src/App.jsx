import { createBrowserRouter, RouterProvider } from "react-router-dom";
import LayoutComponent from './components/layout/LayoutComponent.jsx';
import { createTheme, ThemeProvider } from "@mui/material/styles";
import CssBaseline from "@mui/material/CssBaseline";
import { getDesignTokens } from "./shared-theme/themePrimitives.js";

// Import components directly
import LandingPage from './components/pages/LandingPage.jsx';
import SignIn from './components/pages/SignIn.jsx';
import SignUp from './components/pages/SignUp.jsx';
import UserHome from './components/pages/UserHome.jsx';
import Dashboard from './components/pages/Dashboard.jsx';

// Define routes directly
const router = createBrowserRouter([
    {
        element: <LayoutComponent />,
        children: [
            { path: "/", element: <LandingPage /> },
            { path: "/signin", element: <SignIn /> },
            { path: "/signup", element: <SignUp /> },
            { path: "/home", element: <UserHome /> }, // Include UserHome directly
            { path: "/room", element: <Dashboard /> }, // Include Dashboard directly
        ],
    },
]);

const theme = createTheme(getDesignTokens('dark'));

function App() {
    return (
        <ThemeProvider theme={theme}>
            <CssBaseline />
            <RouterProvider router={router} />
        </ThemeProvider>
    );
}

export default App;