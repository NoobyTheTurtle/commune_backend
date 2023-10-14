FROM golang:1.21.2-alpine AS builder

ENV GOOS=linux

WORKDIR /build

ADD go.mod /build
ADD go.sum /build

COPY cmd /build/cmd
COPY internal /build/internal
COPY docs /build/docs

RUN go build -o app ./cmd/app

FROM alpine

WORKDIR /build

ADD .env /build

COPY seeds /build/seeds

COPY --from=builder /build/app /build/app

CMD ["./app"]