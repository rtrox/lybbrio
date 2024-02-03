export enum ButtonColor {
  Primary = 1,
  Accent,
}

interface ButtonProps {
  color: ButtonColor;
  children: React.ReactNode;
  onClick?: () => void;
}

export default function Button({ color, children, onClick }: ButtonProps) {
  const onClickFunc = onClick || (() => {});
  const bgColor =
    color === ButtonColor.Primary ? "bg-primary-500" : "bg-accent-500";
  const hoverColor =
    color === ButtonColor.Primary
      ? "hover:bg-primary-400"
      : "hover:bg-accent-400";
  return (
    <button
      onClick={onClickFunc}
      type="submit"
      className={`my-2 ${bgColor} text-light-500 font-semibold rounded-full w-80 p-2.5 mb-4 ${hoverColor} transition duration-300 ease-in-out cursor-pointer`}
    >
      {children}
    </button>
  );
}
