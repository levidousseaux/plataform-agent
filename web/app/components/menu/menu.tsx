import { cn } from "@/app/lib/utils";
import Navbar from "./navbar";
import Sidebar from "./sidebar";
import { useState } from "react";

export default function Menu({ children }: { children: React.ReactNode }) {
  const [toggle, setToggle] = useState(
    localStorage.getItem("toggle") === "true"
  );

  function onToggle(toggle: boolean) {
    setToggle(toggle);
    localStorage.setItem("toggle", toggle ? "true" : "false");
  }

  return (
    <div>
      <div className="hidden h-full md:flex md:flex-col md:fixed md:inset-y-0 z-80">
        <Sidebar
          toggle={toggle}
          onRouteChange={() => {}}
          onToggle={() => onToggle(!toggle)}
        />
      </div>
      <main className={cn("pb-10", toggle ? "md:pl-20" : "md:pl-72")}>
        <Navbar />
        <div className="p-8">{children}</div>
      </main>
    </div>
  );
}
