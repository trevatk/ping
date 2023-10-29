FROM golang:1.21-alpine as builder

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod tidy && go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /usr/bin/client ./cmd/client

FROM gcr.io/distroless/static-debian11

COPY --from=builder /usr/bin/ /app/bin/

ENTRYPOINT [ "/app/bin/client" ]