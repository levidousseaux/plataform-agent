import { useState } from "react";
import { Menu } from "lucide-react";
import { Sheet, SheetContent, SheetTrigger } from "@/app/components/ui/sheet";
import Sidebar from "./sidebar";

export default function MobileSidebar() {
  const [open, setOpen] = useState(false);

  function closeSidebar() {
    setOpen(false);
  }

  return (
    <Sheet open={open} onOpenChange={() => setOpen(!open)}>
      <SheetTrigger>
        <Menu className="md:hidden" />
      </SheetTrigger>
      <SheetContent side="left" className="p-0" closebuttoncolor="text-white">
        <Sidebar
          onRouteChange={closeSidebar}
          toggle={false}
          onToggle={() => {}}
        />
      </SheetContent>
    </Sheet>
  );
}
