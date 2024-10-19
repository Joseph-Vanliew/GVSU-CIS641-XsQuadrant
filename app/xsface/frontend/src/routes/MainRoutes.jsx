import Homepage from "../pages/HomePage"
import LayoutComponent from "../pages/LayoutComponent"
import CallComponent from "../pages/CallComponent";

const MainRoutes = [
    // {
    //     path: "/",
    //     element: <Homepage/>,
    //     inSearch: "Homepage"
    // },
    {
        path: "/",
        element: <CallComponent/>,
        inSearch: "Call Component"
    }
]

export default MainRoutes;