FROM golang:1.16.0-alpine AS build_base

RUN apk add --no-cache git \
                        make \
                        upx

ENV GO111MODULE=on

# Set the Current Working Directory inside the container
WORKDIR /tmp/conductor

# We want to populate the module cache based on the go.{mod,sum} files.
# COPY go.mod .
# COPY go.sum .
COPY . .

RUN make mod
RUN go get -u github.com/go-swagger/go-swagger/cmd/swagger
RUN make swagger

# Unit tests
# RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN make build
RUN make pack

# Start fresh from a smaller image
FROM alpine:3.9
RUN apk add ca-certificates

# Copy binary
COPY --from=build_base /tmp/conductor/bin/conductor /app/conductor

# Copy swagger static files
COPY --from=build_base /tmp/conductor/static /static

# Run the binary program produced by `build phase`
CMD /app/conductor

# This container exposes port 8080 to the outside world
EXPOSE 8080
