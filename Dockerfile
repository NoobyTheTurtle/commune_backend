FROM golang:1.21.2-alpine AS builder

ENV GOOS=linux

WORKDIR /build

ADD go.mod /build
ADD go.sum /build

COPY cmd /build/cmd
COPY internal /build/internal

RUN go build -o api_server ./cmd/api_server

FROM alpine

WORKDIR /build

ADD .env /build

COPY seeds /build/seeds

COPY --from=builder /build/api_server /build/api_server

CMD ["./api_server"]