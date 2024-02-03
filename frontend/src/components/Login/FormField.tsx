import { IconType } from "react-icons";

interface FormFieldProps {
  name: string;
  placeholder: string;
  type: string;
  Icon: IconType;
  isSubmitting?: boolean;
  state: string;
  setState?: (value: string) => void;
}

export default function FormField({
  name,
  placeholder,
  type,
  Icon,
  state,
  setState,
  isSubmitting,
}: FormFieldProps) {
  const set = setState || (() => {});
  const disabled = isSubmitting || false;
  return (
    <div className="border-2 bg-white border-light-accent w-80 p-2 flex items-center rounded-full my-2">
      <Icon className="my-1 ml-2 mr-2 text-dark-300" />
      <input
        type={type}
        name={name}
        disabled={disabled}
        placeholder={placeholder}
        className="bg-white w-full outline-none text-sm flex-1 text-dark-600 placeholder-dark-300"
        value={state}
        onChange={(e) => set(e.target.value)}
      />
    </div>
  );
}
