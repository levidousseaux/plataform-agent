import { createBrowserRouter } from "react-router-dom";
import Home from "./home/home";
import Login from "./login/login";
import { HomeIcon } from "lucide-react";
import DashboardLayout from "../components/layouts/dashboard-layout";

export const DASHBOARD_ROUTES = [
  {
    href: "/dashboard",
    label: "Dashboard",
    component: <Home />,
    icon: HomeIcon,
    disabled: false,
  },
];

export const APP_ROUTER = createBrowserRouter([
  { path: "/login", element: <Login /> },
  {
    path: "/",
    element: <DashboardLayout />,
    children: [
      { path: "/", element: <Home /> },
      ...DASHBOARD_ROUTES.map((x) => ({
        path: x.href,
        element: x.component,
      })),
    ],
  },
]);
