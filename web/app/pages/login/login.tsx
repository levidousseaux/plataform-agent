import { Button } from "@/app/components/ui/button";

export default function Login() {
  async function loginWithGoogle() {
    const res = await fetch("/login/google", {
      method: "GET",
    });

    const url = await res.text();
    window.location.href = url;
  }

  return (
    <div>
      <h1>Login</h1>

      <Button onClick={loginWithGoogle}>Login with Google</Button>
    </div>
  );
}
