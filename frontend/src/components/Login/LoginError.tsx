export default function LoginError({
  message,
  visible,
  onClick,
}: {
  message: string;
  visible: boolean;
  onClick?: () => void;
}) {
  const opacity = visible ? "opacity-100" : "opacity-0";
  const onClickFunc = onClick || (() => {});
  return (
    <div
      className={`text-sm italic text-danger-300 w-80 h-0 mb-3 -mt-3 transition-opacity ease-in duration-150 ${opacity} cursor-pointer`}
      onClick={onClickFunc}
    >
      {message}
    </div>
  );
}
