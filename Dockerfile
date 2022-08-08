# syntax=docker/dockerfile:1

FROM golang:1.16

WORKDIR /personalf-services-currencies/

COPY currencies currencies/
COPY handlers handlers/
COPY currencies.go .

RUN go mod init github.com/eosorio/personalf-services-currencies
RUN go mod tidy
RUN go get github.com/eosorio/personalf-services
RUN go get github.com/lib/pq
RUN go build -o /personalf-currencies -ldflags "-linkmode external -extldflags -static"

FROM scratch
LABEL maintainer="eduardo.osorio.it@gmail.com"
EXPOSE 11114
WORKDIR /
COPY --from=0 /personalf-currencies /personalf-currencies
CMD ["/personalf-currencies"]


