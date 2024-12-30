import { Fragment, useEffect } from "react";
import { useSuspenseInfiniteQuery } from "@tanstack/react-query";
import { useInView } from "react-intersection-observer";

import { infiniteBookQueryOptions } from "../../queries/Book";
import { Book } from "../Book/Book";
import { Spinner } from "../Spinner";

import styles from "./InfiniteBookList.module.scss";

export function InfiniteBookList() {
  const { data, isFetching, isLoading, fetchNextPage } =
    useSuspenseInfiniteQuery(infiniteBookQueryOptions({ pageSize: 50 }));
  const [ref, inView] = useInView();

  useEffect(() => {
    if (inView) {
      fetchNextPage();
    }
  }, [inView, fetchNextPage]);

  return (
    <div>
      {isLoading ? (
        <div>Loading...</div>
      ) : (
        <div className={styles.listContainer}>
          {data?.pages.map((page, i) => (
            <Fragment key={i}>
              {page.books.edges?.map((edge, j) => {
                if (edge?.node)
                  return <Book book={edge?.node} key={`film-${j}`} />;
              })}
            </Fragment>
          ))}
          <div ref={ref} className={styles.loaderDiv} />
          {isFetching && <Spinner />}
        </div>
      )}
    </div>
  );
}
