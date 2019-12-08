FROM golang:1.13-alpine

WORKDIR ~/src/pithy-go
ADD go.mod go.mod
ADD api api
RUN go build -o ./bin/api ./api
EXPOSE 8000
ENTRYPOINT ./bin/api
