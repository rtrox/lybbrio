import { useEffect } from "react";
import {
  useNavigate,
  createFileRoute,
  redirect,
  Link,
} from "@tanstack/react-router";
import { useAuth } from "../context/AuthProvider";

export const Route = createFileRoute("/about")({
  beforeLoad: ({ context, location }) => {
    if (!context.auth.loggedIn) {
      console.log("Not authenticated, redirecting to login");
      console.log(context.auth.user);
      console.log(context.auth);
      throw redirect({
        to: "/login",
        search: {
          redirect: location.href,
        },
      });
    }
  },
  component: About,
});

function About() {
  const { user, logout } = useAuth();
  const navigate = useNavigate({ from: "/about" });
  const handleLogout = () => {
    logout().then(() => {
      navigate({ to: "/login", search: { redirect: "" } });
    });
  };

  useEffect(() => {
    if (!user) {
      navigate({ to: "/login", search: { redirect: "/about" } });
    }
  }, [user, navigate]);
  return (
    <div>
      <h3>Hello from About!</h3>
      <Link to="/">Home</Link>
      <Link to="/books">Books</Link>
      <div className="cursor-pointer" onClick={handleLogout}>
        Logout
      </div>
    </div>
  );
}
