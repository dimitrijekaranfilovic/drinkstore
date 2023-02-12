
FROM golang:1.18-alpine3.15

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY main.go ./
COPY handlers/*.go ./handlers/
COPY model/*.go ./model/
COPY util/*.go ./util/

RUN go build -o /comment-service

EXPOSE 8083

CMD [ "/comment-service" ]
