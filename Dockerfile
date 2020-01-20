FROM golang:1.12 AS builder

ARG GOPROXY
ENV GORPOXY ${GOPROXY}

ADD . /builder

WORKDIR /builder

RUN go build main.go

FROM golang:1.12

COPY --from=builder /builder/main /app/mysql-api

WORKDIR /app

CMD ["./mysql-api"]

EXPOSE 9020
