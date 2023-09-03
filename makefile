# ----------------------------------
# Container Name
# ----------------------------------
APP_APPMIXER=app-appmixer

# ----------------------------------
# Docker container operations
# ----------------------------------
build:
	cd .docker && docker compose build --no-cache
up:
	cd .docker && docker compose up -d
ps:
	cd .docker && docker compose ps
restart:
	cd .docker && docker compose restart
start:
	cd .docker && docker compose start
stop:
	cd .docker && docker compose stop
down:
	cd .docker && docker compose down


# ----------------------------------
# Go into container
# ----------------------------------
into-appmixer:
	cd .docker && docker compose exec ${APP_APPMIXER} /bin/sh

# ----------------------------------
# Run grpc on "background"
# ----------------------------------
run-appmixer-back:
	cd .docker && docker compose exec -d ${APP_APPMIXER} /bin/sh -c "go run /go/src/app/services/appmixer/main.go"

# ----------------------------------
# Run grpc on "foregrond"
# ----------------------------------
run-appmixer:
	cd .docker && docker compose exec ${APP_APPMIXER} /bin/sh -c "go run /go/src/app/services/appmixer/main.go"


# proto builder
BUILD_FROM=proto/services
proto-gen:
	protoc --go_out=. \
        --go-grpc_out=. \
        ${BUILD_FROM}/*

# grpc-gateway builder
proto-gw:
	protoc -I ./proto \
		--go_out ./ \
		--go-grpc_out ./ \
		--grpc-gateway_out ./ \
		./proto/gateway/appmixer.proto

# proto-gw:
# 	protoc -I ./src/proto \
# 		--go_out=./ \
# 		--go-grpc_out=. \
# 		--grpc-gateway_out=. \
# 		--validate_out="lang=go:." \
# 		--plugin=protoc-gen-grpc-gateway=${GOBIN}/protoc-gen-grpc-gateway \
# 		proto/gateway/appmixer.proto