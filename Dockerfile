FROM golang:latest AS builder
ARG APP_NAME=stress-test
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server ./cmd/${APP_NAME}

FROM scratch
COPY --from=builder /app/server .
CMD ["./server"]