interface FormProps {
  children: React.ReactNode;
  inactiveTranslate: string;
  isActive: boolean;
  onSubmit?: (evt: React.FormEvent<HTMLFormElement>) => void;
}
export default function Form({
  children,
  isActive,
  inactiveTranslate,
  onSubmit,
}: FormProps) {
  const t = isActive ? "" : inactiveTranslate;
  const handleSubmit = onSubmit || (() => {});
  return (
    <form
      className={`absolute flex flex-col justify-center items-center w-full h-full text-center pt-16 ${t} transition duration-300 ease-in-out`}
      onSubmit={handleSubmit}
    >
      {children}
    </form>
  );
}
