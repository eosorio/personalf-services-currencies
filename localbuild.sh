#!/bin/bash

if [ -z "$BIN_DIR" ]; then BIN_DIR="/tmp"; fi

# Handlers module
cd handlers
go mod init github.com/eosorio/personalf-services-currencies/handlers
go mod edit -replace github.com/eosorio/personalf-services-currencies/currencies=../currencies
go mod edit -replace github.com/eosorio/personalf-services=../../personalf-services
go mod edit -replace github.com/eosorio/personalf-services/databaseInfo=../../personalf-services/databaseInfo
go mod tidy
cd ..

# Currencies module
cd currencies
go mod init github.com/eosorio/personalf-services-currencies/currencies
go mod edit -replace github.com/eosorio/personalf-services-currencies/handlers=../handlers
go mod edit -replace github.com/eosorio/personalf-services=../../personalf-services
go mod edit -replace github.com/eosorio/personalf-services/databaseInfo=../../personalf-services/databaseInfo
go mod tidy
cd ..

# Main module
go mod init main
go mod edit -replace github.com/eosorio/personalf-services-currencies=./
go mod edit -replace github.com/eosorio/personalf-services-currencies/handlers=./handlers
go mod edit -replace github.com/eosorio/personalf-services-currencies/currencies=./currencies
go mod edit -replace github.com/eosorio/personalf-services/databaseInfo=../personalf-services/databaseInfo

go build -o $BIN_DIR/currenciesService -ldflags "-linkmode external -extldflags -static"