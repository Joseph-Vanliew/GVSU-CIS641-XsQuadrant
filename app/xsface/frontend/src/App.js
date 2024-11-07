import './App.css';
import "./pages/LayoutComponent"
import MainRoutes from './routes/MainRoutes';
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import LayoutComponent from './pages/LayoutComponent';
import {createTheme, ThemeProvider} from "@mui/material/styles";
import CssBaseline from "@mui/material/CssBaseline";
import {getDesignTokens} from "./shared-theme/themePrimitives";

const router = createBrowserRouter([
  {
     element: <LayoutComponent />,
     children: MainRoutes
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
  