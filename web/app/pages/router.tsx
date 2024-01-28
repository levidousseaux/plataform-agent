import { createBrowserRouter } from "react-router-dom";
import Home from "./home/home";

export const APP_ROUTER = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
  },
  {
    path: "/2",
    element: <div>home page 2</div>,
  },
]);
