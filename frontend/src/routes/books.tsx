import { createFileRoute, redirect } from "@tanstack/react-router";
import { useGraphQLClient } from "../context/GraphQLProvider";
import { graphql } from "../gql/gql";
import { useQuery } from "@tanstack/react-query";
import { Book } from "../components/Book";

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

const bookQuery = graphql(`
  query allBooks($first: Int!) {
    books(first: $first) {
      edges {
        node {
          ...BookItem
        }
      }
    }
  }
`);

function Books() {
  const { graphql: g } = useGraphQLClient();
  const { data } = useQuery({
    queryKey: ["allBooks", { first: 10 }],
    queryFn: async () => g.request(bookQuery, { first: 10 }),
  });

  return (
    <div>
      {data && (
        <ul>
          {data.books?.edges?.map(
            (e, i) => e?.node && <Book book={e?.node} key={`film-${i}`} />
          )}
        </ul>
      )}
    </div>
  );
}
