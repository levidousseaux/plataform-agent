import * as React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider } from "react-router-dom";
import { APP_ROUTER } from "./pages/router";

document.addEventListener("DOMContentLoaded", () => {
  ReactDOM.createRoot(document.getElementById("root")!).render(
    <React.StrictMode>
      <RouterProvider router={APP_ROUTER} />
    </React.StrictMode>
  );
});
