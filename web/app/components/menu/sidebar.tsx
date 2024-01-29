import { cn } from "@/app/lib/utils";
import { DASHBOARD_ROUTES } from "@/app/pages/router";
import { ChevronsLeft, ChevronsRight } from "lucide-react";
import { Link, useLocation } from "react-router-dom";

export interface SidebarProps {
  toggle: boolean;
  onToggle: () => void;
  onRouteChange: () => void;
}

export default function Sidebar({
  toggle,
  onToggle,
  onRouteChange,
}: SidebarProps) {
  const { pathname } = useLocation();

  return (
    <div className="space-y-0 py-4 flex flex-col h-full bg-gray-200/30 text-gray-700 dark:bg-transparent dark:text-white border-r">
      <div
        className={cn(
          "flex flex-col h-full px-3",
          toggle ? "md:w-20" : "md:w-72"
        )}
      >
        <Link
          to="/dashboard"
          className={cn(
            "flex items-center mt-3 mb-8",
            toggle ? "justify-center" : "pl-3"
          )}
        >
          <h1 className="text-2xl font-extrabold bg-clip-text text-transparent bg-gradient-to-r from-blue-600 to-cyan-900">
            {toggle ? "TA" : "Todo App"}
          </h1>
        </Link>

        <div className="space-y-1">
          {DASHBOARD_ROUTES.map((route) => (
            <Link
              key={route.href}
              to={route.href}
              onClick={onRouteChange}
              className={cn(
                "text-sm group flex p-3 w-full font-medium cursor-pointer hover:bg-gray-300/30 rounded-lg transition",
                pathname.startsWith(route.href)
                  ? "bg-gray-300/30 text-blue-900 dark:text-blue-200"
                  : "",
                route.disabled ? "pointer-events-none text-zinc-400" : "",
                toggle ? "justify-center" : "justify-start"
              )}
            >
              <div
                className={cn(
                  "flex items-center flex-1",
                  toggle ? "justify-center" : ""
                )}
              >
                <route.icon className="h-5 w-5" />
                {!toggle && <span className="ml-3">{route.label}</span>}
              </div>
            </Link>
          ))}
        </div>

        <div
          className={cn(
            "hidden mt-auto md:flex cursor-pointer",
            toggle ? "mx-auto" : "ml-auto mr-3"
          )}
          onClick={onToggle}
        >
          {toggle ? <ChevronsRight /> : <ChevronsLeft />}
        </div>
      </div>
    </div>
  );
}
