import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_layout/layout-a")({
  component: About,
});

function About() {
  return <h3>Hello from Layout A!</h3>;
}
