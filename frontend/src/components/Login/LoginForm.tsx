import React, { useState } from "react";
import { useNavigate } from "@tanstack/react-router";
import { UnauthorizedError, useAuth } from "../../context/AuthProvider";
import Form from "./Form";
import FormField from "./FormField";
import Button, { ButtonColor } from "./Button";
import LoginError from "./LoginError";
import { FaRegUser } from "react-icons/fa";
import { MdLockOutline } from "react-icons/md";

export default function LoginForm({
  inactiveTranslate,
  setRegister,
  isActive,
  redirect,
}: {
  inactiveTranslate: string;
  setRegister: () => void;
  isActive: boolean;
  redirect: string;
}) {
  const [username, setUsername] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const { login } = useAuth();
  const navigate = useNavigate({ from: "/login" });
  const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
  const [error, setError] = useState<string>("");

  const handleLogin = (evt: React.FormEvent<HTMLFormElement>) => {
    evt.preventDefault();
    setIsSubmitting(true);

    login(username, password)
      .then(() => {
        navigate({ to: redirect });
      })
      .catch((error) => {
        console.log(error);
        if (error == UnauthorizedError) {
          setError("Invalid username or password");
        } else {
          setError("An error occurred");
        }
        setUsername("");
        setPassword("");
        setIsSubmitting(false);
      });
  };
  const handleSetRegister = () => {
    setError("");
    setRegister();
  };
  return (
    <Form
      inactiveTranslate={inactiveTranslate}
      isActive={isActive}
      onSubmit={handleLogin}
    >
      {/* Sign In */}
      <h2 className="text-4xl font-bold font-serif text-light-500 mb-5">
        Sign in to Start Reading
      </h2>
      <div className="border-2 w-10 border-light-accent inline-block mb-5" />
      <LoginError
        message={error}
        visible={error != ""}
        onClick={() => setError("")}
      />
      <FormField
        name="name"
        placeholder="Username"
        type="text"
        Icon={FaRegUser}
        isSubmitting={isSubmitting}
        state={username}
        setState={(value: string) => setUsername(value)}
      />
      <FormField
        name="password"
        placeholder="Password"
        type="password"
        Icon={MdLockOutline}
        isSubmitting={isSubmitting}
        state={password}
        setState={(value: string) => setPassword(value)}
      />
      <Button color={ButtonColor.Primary}>Sign In</Button>

      <div className="flex items w-80 justify-between mb-8">
        <label className="flex items-center text-xs">
          <input type="checkbox" name="remember" className="mr-1" />
          Remember me
        </label>
        <a href="#" className="text-xs hover:underline">
          Forgot Password?
        </a>
      </div>

      <div className="flex items-center justify-between w-80 mb-6">
        <div className=" border-2 w-20 h-min border-light-accent inline-block" />
        <p className="">or</p>
        <div className=" border-2 w-20 h-min border-light-accent inline-block" />
      </div>

      <Button color={ButtonColor.Accent}>Sign in with Keycloak</Button>

      <div className="text-center w-full mt-4 mb-14">
        <p>
          Don't have an account?{" "}
          <span
            className="font-semibold text-primary-500 cursor-pointer hover:underline"
            onClick={handleSetRegister}
          >
            Sign up
          </span>
        </p>{" "}
      </div>
    </Form>
  );
}
