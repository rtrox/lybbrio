import { useState } from "react";
import Form from "./Form";
import FormField from "./FormField";
import Button, { ButtonColor } from "./Button";
import { FaRegUser, FaRegEnvelope } from "react-icons/fa";
import { MdLockOutline } from "react-icons/md";

export default function RegisterForm({
  inactiveTranslate,
  isActive,
  setSignIn,
}: {
  inactiveTranslate: string;
  isActive: boolean;
  setSignIn: () => void;
}) {
  const [username, setUsername] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  return (
    <Form inactiveTranslate={inactiveTranslate} isActive={isActive}>
      <h2 className="text-4xl font-bold font-serif text-light-500 mb-5">
        Create an Account
      </h2>
      <div className="border-2 w-10 border-light-accent inline-block mb-5" />
      <FormField
        name="username"
        placeholder="Username"
        type="text"
        state={username}
        setState={(value: string) => setUsername(value)}
        Icon={FaRegUser}
      />
      <FormField
        name="email"
        placeholder="Email"
        type="email"
        state={email}
        setState={(value: string) => setEmail(value)}
        Icon={FaRegEnvelope}
      />
      <FormField
        name="password"
        placeholder="Password"
        type="password"
        state={password}
        setState={(value: string) => setPassword(value)}
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
            onClick={() => setSignIn()}
          >
            Sign in
          </span>
        </p>{" "}
      </div>
    </Form>
  );
}
