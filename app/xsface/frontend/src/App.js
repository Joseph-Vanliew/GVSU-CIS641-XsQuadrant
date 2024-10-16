import logo from './logo.svg';
import './App.css';
import "./pages/LayoutComponent"
import MainRoutes from './routes/MainRoutes';
import {createBrowserRouter, RouterProvider, Outlet} from "react-router-dom";
import LayoutComponent from './pages/LayoutComponent';

const router = createBrowserRouter([
  {
     element: <LayoutComponent />,
     children: MainRoutes
  },
]);

function App() {

  return <RouterProvider router={router} />
}

export default App;
  