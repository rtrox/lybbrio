import { infiniteQueryOptions } from "@tanstack/react-query";
import { BookWhereInput } from "../gql/graphql";
import { useGraphQLClient } from "../context/GraphQLProvider";
import { graphql } from "../gql";
import { GraphQLClient } from "graphql-request";

export const BookFragment = graphql(/* GraphQL */ `
  fragment BookItem on Book {
    id
    title
    authors {
      name
    }
    covers {
      width
      height
      url
    }
    tags {
      name
    }
    description
  }
`);

interface infiniteBookQueryProps {
  pageSize: number;
  where?: BookWhereInput;
}

const whereablePaginatedBookQuery = graphql(`
  query whereableBooks($first: Int!, $after: Cursor, $where: BookWhereInput) {
    books(first: $first, after: $after, where: $where) {
      totalCount
      edges {
        node {
          ...BookItem
        }
      }
      pageInfo {
        hasNextPage
        startCursor
        endCursor
      }
    }
  }
`);

export const infiniteBookQueryOptions = ({
  pageSize,
  where,
}: infiniteBookQueryProps) => {
  const { graphql: g } = useGraphQLClient();
  return infiniteQueryOptions({
    queryKey: ["books", { first: pageSize, where: where }],
    queryFn: async ({ pageParam }) =>
      g.request(whereablePaginatedBookQuery, {
        first: pageSize,
        after: pageParam,
        before: pageParam,
        where: where,
      }),
    initialPageParam: undefined,
    getNextPageParam: (lastPage) =>
      lastPage.books.pageInfo.endCursor ?? undefined,
  });
};
