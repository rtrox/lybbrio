# Calibre Library

To access calibre, we are directly integrating with the `metadata.db` database calibre generates, using a repository pattern in [`internal/calibre`](internal/calibre/). These models are then exposed via lybbr.io's internal API server for use in the front-end, and in 3rd party integrations.
