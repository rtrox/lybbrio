import { Outlet, createRootRouteWithContext } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/router-devtools";
import { QueryClient } from "@tanstack/react-query";

import { AuthContext } from "../context/AuthProvider";
import { GraphQLContext } from "../context/GraphQLProvider";

interface RouterContext {
  auth: AuthContext;
  graphql: GraphQLContext;
  queryClient: QueryClient;
}
export const Route = createRootRouteWithContext<RouterContext>()({
  component: () => (
    <>
      <Outlet />
      <TanStackRouterDevtools />
    </>
  ),
});
