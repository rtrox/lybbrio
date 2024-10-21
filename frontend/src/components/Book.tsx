import { FragmentType, useFragment } from "../gql/fragment-masking";
import { graphql } from "../gql/gql";
import { Img } from "./Image";

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
    description
  }
`);

export function Book({ book }: { book: FragmentType<typeof BookFragment> }) {
  const bk = useFragment(BookFragment, book);
  const cover = bk.covers?.[0];
  const url = cover?.url || "";
  return (
    <div className="flex flex-col w-[260px]">
      <div className="flex-none">
        <Img className="pr-10" width={260} url={url} alt={bk.title} />
      </div>
      <div>
        <h3 className="text-md overflow-hidden">{bk.title}</h3>
        <p className="text-sm italic text-gray-500">
          by {bk.authors?.map((a) => <span>{a.name}</span>)}
        </p>
      </div>
    </div>
  );
}
