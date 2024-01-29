import MobileSidebar from "./mobile-sidebar";

export default function Navbar() {
  return (
    <div className="flex items-center p-4 border-b h-[72px]">
      <MobileSidebar />
      <div className="flex w-full justify-end"></div>
    </div>
  );
}
