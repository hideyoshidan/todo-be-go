# gRPC Greeting
FROM golang:1.20.0-alpine3.17 AS base
RUN apk update && apk add git


###################################
# gPRC app-appmixer
###################################
FROM base AS app-appmixer
WORKDIR /go/todo.com/app

COPY ../../todo.com /go/todo.com/app
RUN go mod tidy

###################################
# gPRC app-statusmaster
###################################
FROM base AS app-statusmaster
WORKDIR /go/todo.com/app

COPY ../../todo.com /go/todo.com/app
RUN go mod tidy