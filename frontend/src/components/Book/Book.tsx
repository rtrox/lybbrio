import { useState, useLayoutEffect, useRef, useEffect } from "react";
import { FragmentType, useFragment } from "../../gql/fragment-masking";
import { useDebounce } from "../../hooks/useDebounce";
import { BookFragment } from "../../queries/Book";

import { FaBook, FaCircleCheck } from "react-icons/fa6";
import { BsThreeDots } from "react-icons/bs";

import { Img } from "../Image";

import styles from "./Book.module.scss";

export function Book({ book }: { book: FragmentType<typeof BookFragment> }) {
  const bk = useFragment(BookFragment, book);
  const container = useRef<HTMLDivElement>(null);
  const [size, setSize] = useState({ width: 0, height: 0 });
  const handleResize = useDebounce(() => {
    setSize({
      height: container.current?.offsetHeight || 0,
      width: container.current?.offsetWidth || 0,
    });
  }, 100);
  useEffect(() => {
    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  });
  useLayoutEffect(() => {
    setSize({
      height: container.current?.offsetHeight || 0,
      width: container.current?.offsetWidth || 0,
    });
  }, [container.current]);
  const cover = bk.covers?.[0];
  const authors = bk.authors?.map((a) => a.name).join(", ");
  const url = cover?.url || "";
  return (
    <div className={styles.container} ref={container}>
      <Img url={url} alt={bk.title} />
      <div
        className={styles.info}
        style={{ height: size.height, width: size.width }}
      >
        <div className={styles.menu}>
          <BsThreeDots />
        </div>
        <h1 className="text-center">{bk.title}</h1>
        <p className={styles.byline}>by {authors}</p>
        <p className={styles.tags}>
          {bk.tags?.slice(0, 5).map((tag, i) => (
            <span key={i} className={styles.tag}>
              {tag.name}
            </span>
          ))}
        </p>
      </div>
      <div className={styles.bottomBadges}>
        {/* <FaBook
          onClick={() => console.log('Clicked "Want to Read" on ' + bk.title)}
        />
        <FaCircleCheck
          onClick={() => console.log('Clicked "Read" on ' + bk.title)}
        /> */}
      </div>
    </div>
  );
}
