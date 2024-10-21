import { useState, useEffect } from "react";
import {
  createFileRoute,
  useNavigate,
  getRouteApi,
} from "@tanstack/react-router";

import { useAuth } from "../context/AuthProvider";
import { z } from "zod";

import LoginForm from "../components/Login/LoginForm";
import RegisterForm from "../components/Login/RegisterForm";

const routeApi = getRouteApi("/login");

export const Route = createFileRoute("/login")({
  validateSearch: z.object({
    redirect: z.string().catch("/"),
  }),
  component: Login,
});

function Login() {
  const [isSignIn, setIsSignIn] = useState<boolean>(true);
  const h = isSignIn ? "h-[600px]" : "h-[640px]";
  const { loggedIn, initialized } = useAuth();
  const navigate = useNavigate({ from: "/login" });
  const search = routeApi.useSearch();

  useEffect(() => {
    if (loggedIn) {
      navigate({ to: search.redirect });
    }
  }, [loggedIn]);

  return (
    <>
      {initialized && !loggedIn ? (
        <div className="flex items-center justify-center min-h-screen py-2 bg-dark-600">
          <div
            className={`relative flex bg-dark-500 text-light-500 rounded-2xl shadow-2xl w-11/12 sm:w-[32rem] ${h} overflow-hidden transition-all duration-300 ease-in-out`}
          >
            <LoginForm
              inactiveTranslate="-translate-x-[600px]"
              isActive={isSignIn}
              setRegister={() => setIsSignIn(false)}
              redirect={search.redirect}
            />

            <RegisterForm
              inactiveTranslate="translate-x-[600px]"
              isActive={!isSignIn}
              setSignIn={() => setIsSignIn(true)}
            />
          </div>
        </div>
      ) : (
        <div></div>
      )}
    </>
  );
}
