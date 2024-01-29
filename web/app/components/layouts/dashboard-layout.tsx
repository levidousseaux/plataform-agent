import { Outlet } from "react-router-dom";
import Menu from "../menu/menu";

export default function DashboardLayout() {
  return (
    <div className="h-full relative">
      <Menu>
        <Outlet />
      </Menu>
    </div>
  );
}
