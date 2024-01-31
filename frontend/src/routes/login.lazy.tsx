import { useState } from "react";
import { createLazyFileRoute } from "@tanstack/react-router";

import { IconType } from "react-icons";
import { FaRegEnvelope, FaRegUser } from "react-icons/fa";
import { MdLockOutline } from "react-icons/md";

export const Route = createLazyFileRoute("/login")({
  component: Login,
});

enum ButtonColor {
  Primary = 1,
  Accent,
}

interface ButtonProps {
  color: ButtonColor;
  children: React.ReactNode;
  onClick?: () => void;
}

function Button({ color, children, onClick }: ButtonProps) {
  const onClickFunc = onClick || (() => {});
  const bgColor =
    color === ButtonColor.Primary ? "bg-primary-500" : "bg-accent-500";
  const hoverColor =
    color === ButtonColor.Primary
      ? "hover:bg-primary-400"
      : "hover:bg-accent-400";
  return (
    <div
      onClick={onClickFunc}
      className={`my-2 ${bgColor} text-light-500 font-semibold rounded-full w-80 p-2.5 mb-4 ${hoverColor} transition duration-300 ease-in-out cursor-pointer`}
    >
      {children}
    </div>
  );
}

interface FormFieldProps {
  name: string;
  placeholder: string;
  type: string;
  Icon: IconType;
}

function FormField({ name, placeholder, type, Icon }: FormFieldProps) {
  return (
    <div className="border-2 bg-white border-light-accent w-80 p-2 flex items-center rounded-full my-2">
      <Icon className="my-1 ml-2 mr-2 text-dark-300" />
      <input
        type={type}
        name={name}
        placeholder={placeholder}
        className="bg-white w-full outline-none text-sm flex-1 text-dark-600 placeholder-dark-300"
      />
    </div>
  );
}

interface FormProps {
  children: React.ReactNode;
  translate: string;
  isActive: boolean;
}
function Form({ children, isActive, translate }: FormProps) {
  const t = isActive ? "" : translate;
  return (
    <div
      className={`absolute flex flex-col justify-center items-center w-full h-full text-center pt-16 ${t} transition duration-300 ease-in-out`}
    >
      {children}
    </div>
  );
}

function Login() {
  const [isSignIn, setIsSignIn] = useState<boolean>(true);
  const h = isSignIn ? "h-[600px]" : "h-[640px]";

  return (
    <div className="flex items-center justify-center min-h-screen py-2 bg-dark-600">
      <div
        className={`relative flex bg-dark-500 text-light-500 rounded-2xl shadow-2xl w-11/12 sm:w-[32rem] ${h} overflow-hidden transition-all duration-300 ease-in-out`}
      >
        <Form translate="-translate-x-[600px]" isActive={isSignIn}>
          {/* Sign In */}
          <h2 className="text-4xl font-bold font-serif text-light-500 mb-5">
            Sign in to Start Reading
          </h2>
          <div className="border-2 w-10 border-light-accent inline-block mb-5" />
          <FormField
            name="email"
            placeholder="Email"
            type="email"
            Icon={FaRegEnvelope}
          />
          <FormField
            name="password"
            placeholder="Password"
            type="password"
            Icon={MdLockOutline}
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
                onClick={() => setIsSignIn(false)}
              >
                Sign up
              </span>
            </p>{" "}
          </div>
        </Form>

        <Form translate="translate-x-[600px]" isActive={!isSignIn}>
          {/* Register */}
          <h2 className="text-4xl font-bold font-serif text-light-500 mb-5">
            Create an Account
          </h2>
          <div className="border-2 w-10 border-light-accent inline-block mb-5" />
          <FormField
            name="username"
            placeholder="Username"
            type="text"
            Icon={FaRegUser}
          />
          <FormField
            name="email"
            placeholder="Email"
            type="email"
            Icon={FaRegEnvelope}
          />
          <FormField
            name="password"
            placeholder="Password"
            type="password"
            Icon={MdLockOutline}
          />
          <Button color={ButtonColor.Primary}>Register</Button>

          <div className="flex items-center justify-between w-80 m-6">
            <div className=" border-2 w-20 h-min border-light-accent inline-block" />
            <p className="">or</p>
            <div className=" border-2 w-20 h-min border-light-accent inline-block" />
          </div>
          <Button color={ButtonColor.Accent}>Register with Keycloak</Button>

          <div className="text-center w-full mt-4 mb-14">
            <p>
              Already Registered?{" "}
              <span
                className="font-semibold text-primary-500 cursor-pointer hover:underline"
                onClick={() => setIsSignIn(true)}
              >
                Sign in
              </span>
            </p>{" "}
          </div>
        </Form>
      </div>
    </div>
  );
}
