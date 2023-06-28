FROM golang:1.21.0-alpine AS build_base

ARG CI_PROJECT_NAME

ENV SERVICE_NAME $CI_PROJECT_NAME

LABEL stage=intermediate
LABEL service=$SERVICE_NAME

ENV CGO_ENABLED=1
ENV GO111MODULE=on
RUN apk add --no-cache git  git gcc g++

# Set the Current Working Directory inside the container
WORKDIR /src

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/app ./cmd/api/main.go

# Start fresh from a smaller image
FROM golang:1.21.0-alpine
RUN apk add ca-certificates

RUN addgroup -S gouser && adduser -S -G gouser gouser 

WORKDIR /app

COPY --chown=gouser:gouser --from=build_base /src/out/app /app/service
COPY --chown=gouser:gouser --from=build_base /src/bin /app/bin

RUN chmod +x service
RUN chmod +x bin/start

USER 1000

# This container exposes port 8080 to the outside world
EXPOSE 3000

# Run the binary program produced
# ENTRYPOINT ./service
CMD [ "./bin/start" ]
