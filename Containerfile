# syntax=docker/dockerfile:1

FROM golang:1.16 AS build
WORKDIR /personalf-services-currencies

# databaseInfo package
WORKDIR /go/src/git.local.osmonfam.net
ADD https://github.com/eosorio/personalf-services.git .
WORKDIR /go/src/git.local.osmonfam.net/personalf-services/databaseInfo
RUN go mod init git.local.osmonfam.net/personalf-services/databaseInfo
RUN go mod edit -replace git.local.osmonfam.net/databaseInfo=/go/src/git.local.osmonfam.net/personalf-services/databaseInfo
RUN go get github.com/lib/pq

# handlers package
WORKDIR /go/src/git.local.osmonfam.net
COPY handlers personalf-services-currencies/handlers/
WORKDIR /go/src/git.local.osmonfam.net/personalf-services-currencies/handlers
RUN go mod init git.local.osmonfam.net/personalf-services-currencies/handlers
RUN go mod edit -replace git.local.osmonfan.net/personalf-services-currencies/currencies=../currencies
RUN go mod edit -replace git.local.osmonfam.net/personalf-services/databaseInfo=../../personalf-services/databaseInfo
RUN go mod tidy

# currencies package
WORKDIR /go/src/git.local.osmonfam.net
COPY currencies personalf-services-currencies/currencies/
WORKDIR /go/src/git.local.osmonfam.net/personalf-services-currencies/currencies
RUN go mod init git.local.osmonfam.net/personalf-services-currencies/currencies
RUN go mod edit -replace git.local.osmonfam.net/personalf-services/databaseInfo=../../personalf-services/databaseInfo
RUN go mod tidy

# main package
WORKDIR /go/src/git.local.osmonfam.net
COPY currencies.go personalf-services-currencies/
WORKDIR /go/src/git.local.osmonfam.net/personalf-services-currencies
RUN go mod init git.local.osmonfam.net/personalf-services-currencies
RUN go mod edit -replace git.local.osmonfam.net/personalf-services-currencies/handlers=./handlers
RUN go mod edit -replace git.local.osmonfan.net/personalf-services-currencies/currencies=./currencies
RUN go mod edit -replace git.local.osmonfam.net/personalf-services/databaseInfo=../personalf-services/databaseInfo
RUN go mod tidy

RUN go build -o /personalf-services-currencies/currenciesService -ldflags "-linkmode external -extldflags -static"
EXPOSE 11114

FROM scratch
LABEL maintainer="eduardo.osorio.it@gmail.com"
EXPOSE 11114
WORKDIR /
COPY --from=build /personalf-services-currencies/currenciesService /personalf-services-currencies/currenciesService
CMD ["/personalf-services-currencies/currenciesService"]
