# Developing on Lybbrio

## Calibre Library

To access calibre, we are directly integrating with the `metadata.db` database calibre generates, using a repository pattern in [`internal/calibre`](internal/calibre/). These models are then exposed via lybbr.io's internal API server for use in the front-end, and in 3rd party integrations.

## Regenerating Swagger

```bash
swag init -g ./cmd/lybbrio/main.go --pd
```

## Gotchas

### Request ID should NOT be considered secure

Request IDs are meant to promote traceability, by creating an ID that can be used to correlate individual requests through the multiple distributed systems and multiple pieces of the codebase. As such, always assume the request id was set by a malicious actor, and do not use it for anything other than the tracing function for which it was intended.
