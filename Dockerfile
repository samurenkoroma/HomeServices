# syntax=docker/dockerfile:1

FROM golang:1.25.5-alpine as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# RUN GOARCH=arm64 go build -o application cmd/main.go
RUN go build -o application cmd/main.go
RUN go build -o scaner cmd/scaner/main.go
RUN go build -o migrate migrations/auto.go

FROM alpine

COPY --from=build /app/application /app/application
COPY --from=build /app/scaner /app/scaner
COPY --from=build /app/migrate /app/migrate

CMD ["/app/application"]