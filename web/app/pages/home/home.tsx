import { Header } from "@/app/components/header";
import { Button } from "@/app/components/ui/button";

export default function Home() {
  return (
    <div>
      <Header />
      <p className="text-2xl font-black">Home 2</p>
      <Button>Clique aqui</Button>
    </div>
  );
}
