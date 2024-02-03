import { FragmentType, useFragment } from "../gql/fragment-masking";
import { graphql } from "../gql/gql";

export const BookFragment = graphql(/* GraphQL */ `
  fragment BookItem on Book {
    id
    title
    authors {
      name
    }
    description
  }
`);

export function Book({ book }: { book: FragmentType<typeof BookFragment> }) {
  const bk = useFragment(BookFragment, book);
  return (
    <div>
      <h3 className="text-xl">{bk.title}</h3>
      <p className="text-sm italic text-gray-500">
        by {bk.authors?.map((a) => <span>{a.name}</span>)}
      </p>
      <p>{bk.description}</p>
    </div>
  );
}
