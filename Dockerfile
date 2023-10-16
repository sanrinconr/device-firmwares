FROM golang:1.21.3-alpine3.18 AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main dc-nearshore/cmd/main.go

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp  /build/main .
RUN cp  /build/.env .

# Command to run when starting the container
CMD ["/dist/main"]

# Build a small image
FROM alpine

RUN apk add --no-cache tzdata
ENV TZ America/Bogota
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY --from=builder /dist/ /

# Command to run
ENTRYPOINT ["/main"]
