import { FaCircleNotch } from "react-icons/fa";

import styles from "./Spinner.module.scss";

export function Spinner() {
  return (
    <div>
      <FaCircleNotch className={styles.spinner} />
    </div>
  );
}
