FROM golang:1.21.2-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./src /app/src

RUN go build -o app-api src/cmd/app-api/main.go

EXPOSE 1323

CMD [ "/app/app-api" ]
