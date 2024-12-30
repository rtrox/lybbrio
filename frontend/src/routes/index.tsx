import { createFileRoute } from "@tanstack/react-router";
import { Link } from "@tanstack/react-router";
import { useAuth } from "../context/AuthProvider";
import { useNavigate } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
  component: Index,
});

function Index() {
  const { loggedIn, logout } = useAuth();
  const navigate = useNavigate({ from: "/" });
  const handleLogout = () => {
    logout().then(() => {
      navigate({ to: "/login", search: { redirect: "/" } });
    });
  };
  return (
    <div>
      <h3 className="text-3xl font-bold underline">Welcome Home!</h3>
      <ul>
        <li>
          <Link to="/login" search={{ redirect: "/" }}>
            Login
          </Link>
        </li>
        <li>
          <Link to="/about">About</Link>
        </li>
        <li>
          <Link to="/books">Books</Link>
        </li>
        {loggedIn && (
          <li>
            <div className="cursor-pointer" onClick={handleLogout}>
              Logout
            </div>
          </li>
        )}
      </ul>
    </div>
  );
}
