import { createBrowserRouter, RouterProvider } from "react-router-dom";
import LayoutComponent from './components/layout/LayoutComponent.jsx';
import { createTheme, ThemeProvider } from "@mui/material/styles";
import CssBaseline from "@mui/material/CssBaseline";
import { getDesignTokens } from "./shared-theme/themePrimitives";

// Import components directly
import LandingPage from './components/pages/LandingPage';
import SignIn from './components/pages/SignIn';
import SignUp from './components/pages/SignUp';
import UserHome from './components/pages/UserHome';
import Dashboard from './components/pages/Dashboard';

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