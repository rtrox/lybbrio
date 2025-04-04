/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
    "\n  fragment BookItem on Book {\n    id\n    title\n    authors {\n      name\n    }\n    covers {\n      width\n      height\n      url\n    }\n    tags {\n      name\n    }\n    description\n  }\n": types.BookItemFragmentDoc,
    "\n  query whereableBooks($first: Int!, $after: Cursor, $where: BookWhereInput) {\n    books(first: $first, after: $after, where: $where) {\n      totalCount\n      edges {\n        node {\n          ...BookItem\n        }\n      }\n      pageInfo {\n        hasNextPage\n        startCursor\n        endCursor\n      }\n    }\n  }\n": types.WhereableBooksDocument,
};

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = graphql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function graphql(source: string): unknown;

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  fragment BookItem on Book {\n    id\n    title\n    authors {\n      name\n    }\n    covers {\n      width\n      height\n      url\n    }\n    tags {\n      name\n    }\n    description\n  }\n"): (typeof documents)["\n  fragment BookItem on Book {\n    id\n    title\n    authors {\n      name\n    }\n    covers {\n      width\n      height\n      url\n    }\n    tags {\n      name\n    }\n    description\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query whereableBooks($first: Int!, $after: Cursor, $where: BookWhereInput) {\n    books(first: $first, after: $after, where: $where) {\n      totalCount\n      edges {\n        node {\n          ...BookItem\n        }\n      }\n      pageInfo {\n        hasNextPage\n        startCursor\n        endCursor\n      }\n    }\n  }\n"): (typeof documents)["\n  query whereableBooks($first: Int!, $after: Cursor, $where: BookWhereInput) {\n    books(first: $first, after: $after, where: $where) {\n      totalCount\n      edges {\n        node {\n          ...BookItem\n        }\n      }\n      pageInfo {\n        hasNextPage\n        startCursor\n        endCursor\n      }\n    }\n  }\n"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;