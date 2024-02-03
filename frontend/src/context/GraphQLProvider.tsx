import React from "react";
import { createContext } from "react";
import { GraphQLClient } from "graphql-request";
import { useAuth } from "./AuthProvider";

export interface GraphQLContext {
  graphql: GraphQLClient;
}

const GraphQLContext = createContext<GraphQLContext | null>(null);

export function GraphQLProvider({ children }: { children: React.ReactNode }) {
  const { fetcher } = useAuth();
  const graphql = new GraphQLClient("/graphql", { fetch: fetcher });
  return (
    <GraphQLContext.Provider value={{ graphql }}>
      {children}
    </GraphQLContext.Provider>
  );
}

export function useGraphQLClient() {
  const context = React.useContext(GraphQLContext);
  if (!context) {
    throw new Error("useGraphQL must be used within a GraphQLProvider");
  }
  return context;
}
