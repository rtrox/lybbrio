import { createLazyFileRoute } from "@tanstack/react-router";

export const Route = createLazyFileRoute("/")({
  component: Index,
});

function Index() {
  return (
      <h3 className="text-3xl font-bold underline">Welcome Home!</h3>
  );
}
