# ----------------------------------
# Container Name
# ----------------------------------
APP_APPMIXER=app-appmixer
APP_STATUSMASTER=app-statusmaster

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


run-all-back:
	./start-grpc.sh

# ----------------------------------
# Go into container
# ----------------------------------
into-appmixer:
	cd .docker && docker compose exec ${APP_APPMIXER} /bin/sh
into-statusmaster:
	cd .docker && docker compose exec ${APP_STATUSMASTER} /bin/sh


# ----------------------------------
# Database Operation
# ----------------------------------
db-migrate:
	cd .docker && docker compose exec -d ${APP_APPMIXER} /bin/sh -c "go run /go/todo.com/app/cmd/main.go -d migrate"

db-drop:
	cd .docker && docker compose exec -d ${APP_APPMIXER} /bin/sh -c "go run /go/todo.com/app/cmd/main.go -d drop"
# ----------------------------------
# Run grpc on "background"
# ----------------------------------
run-appmixer-back:
	cd .docker && docker compose exec -d ${APP_APPMIXER} /bin/sh -c "go run /go/todo.com/app/services/appmixer/main.go"
run-statusmaster-back:
	cd .docker && docker compose exec -d ${APP_STATUSMASTER} /bin/sh -c "go run /go/todo.com/app/services/statusmaster/main.go"
# ----------------------------------
# Run grpc on "foregrond"
# ----------------------------------
run-appmixer:
	cd .docker && docker compose exec ${APP_APPMIXER} /bin/sh -c "go run /go/todo.com/app/services/appmixer/main.go"
run-statusmaster:
	cd .docker && docker compose exec ${APP_STATUSMASTER} /bin/sh -c "go run /go/todo.com/app/services/statusmaster/main.go"

# proto builder
proto-gen:
	protoc -I ./proto \
		--go_out=. \
        --go-grpc_out=. \
        proto/services/*

# grpc-gateway builder
proto-gw:
	protoc -I ./proto \
		--go_out ./ \
		--go-grpc_out ./ \
		--grpc-gateway_out ./ \
		./proto/gateway/appmixer.proto

# proto-gw:
# 	protoc -I ./todo.com/proto \
# 		--go_out=./ \
# 		--go-grpc_out=. \
# 		--grpc-gateway_out=. \
# 		--validate_out="lang=go:." \
# 		--plugin=protoc-gen-grpc-gateway=${GOBIN}/protoc-gen-grpc-gateway \
# 		proto/gateway/appmixer.proto