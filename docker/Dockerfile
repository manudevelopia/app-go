FROM golang:1.21.2-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./src /app/src

RUN go build -o api src/cmd/api/main.go

EXPOSE 1323

CMD [ "/app/api" ]
