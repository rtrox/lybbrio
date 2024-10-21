import "./index.css";
import { StrictMode } from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider, createRouter } from "@tanstack/react-router";

import { useAuth, AuthProvider } from "./context/AuthProvider";
import { useGraphQLClient, GraphQLProvider } from "./context/GraphQLProvider";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
// Import the generated route tree
import { routeTree } from "./routeTree.gen.ts";

const queryClient = new QueryClient();

// Create a new router instance
const router = createRouter({
  routeTree,
  context: {
    auth: undefined!,
    graphql: undefined!,
    queryClient: queryClient,
  },
});

// Register the router instance for type safety
declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

function InnerApp() {
  const auth = useAuth();
  const graphql = useGraphQLClient();
  return <RouterProvider router={router} context={{ auth, graphql }} />;
}

// Render the app
const rootElement = document.getElementById("app")!;
if (!rootElement.innerHTML) {
  const root = ReactDOM.createRoot(rootElement);
  root.render(
    <StrictMode>
      <AuthProvider>
        <GraphQLProvider>
          <QueryClientProvider client={queryClient}>
            <InnerApp />
          </QueryClientProvider>
        </GraphQLProvider>
      </AuthProvider>
    </StrictMode>
  );
}
