import { Outlet, Link as RouterLink } from "react-router-dom";
import Link from "@mui/material/Link";

export default function LayoutComponent() {
    return (
        <div>
            <header>
                <Link
                    component={RouterLink}
                    to="/" // -> Landing page, redirects to user home if user has a valid token
                    style={{textDecoration: "none", fontSize: "1.5rem", fontWeight: "bold", color: "inherit",}}>
                    XsQuadrant
                </Link>
            </header>
            <main>
                <Outlet />
            </main>
            <footer></footer>
        </div>
    );
};