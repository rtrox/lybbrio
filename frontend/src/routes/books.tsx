import { Suspense } from "react";
import { createFileRoute, redirect } from "@tanstack/react-router";

import { InfiniteBookList } from "../components/InfiniteBookList/InfiniteBookList";

export const Route = createFileRoute("/books")({
  component: Books,
  beforeLoad: ({ context, location }) => {
    if (!context.auth.loggedIn) {
      throw redirect({
        to: "/login",
        search: {
          redirect: location.href,
        },
      });
    }
  },
});

function Books() {
  return (
    <Suspense fallback={<h1> Loading ...</h1>}>
      <div className="p-8">
        <InfiniteBookList />
      </div>
    </Suspense>
  );
}
