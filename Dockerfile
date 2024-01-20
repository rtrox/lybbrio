FROM golang:1.20-alpine AS backend_build
WORKDIR /tmp/lybbrio

ARG VERSION="dev"
ARG REVISION=""
ARG BUILDTIME=""

RUN apk --update add \
    ca-certificates \
    gcc \
    g++

# Copy go.mod and go.sum first to leverage Docker layer cache
COPY go.mod ./
COPY go.sum ./
# Cache dependency packages
RUN go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go get

# Copy only source code directories to minimize cache misses
COPY . .

# CGO is required by the sqlite3 driver
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
        CGO_ENABLED=1 go build \
        -ldflags="-linkmode external -extldflags '-static' -s -w -X main.version=${VERSION} -X main.revision=${REVISION} -X main.buildTime=${BUILDTIME}" \
        -o /tmp/lybbrio/out/lybbrio \
        ./cmd/lybbrio/main.go

FROM scratch
COPY --from=backend_build /tmp/lybbrio/out/lybbrio /bin/lybbrio
COPY --from=backend_build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 8080/tcp
ENTRYPOINT ["/bin/lybbrio"]