import { Outlet } from "react-router-dom";

export default function LayoutComponent() {
    return (
      <div>
       <header>XsQuadrant</header>
       <main>
         <Outlet /> 
        </main>
        <footer>Footer Content</footer>
      </div>
    );
  };
