
# syntax=docker/dockerfile:1

FROM golang:1.18-alpine3.15

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY main.go ./
COPY database/*.go ./database/
COPY handlers/*.go ./handlers/
COPY images/* ./images/
COPY model/*.go ./model/
COPY repository/*.go ./repository/
COPY util/*.go ./util/

RUN go build -o /drink-service

EXPOSE 8082

CMD [ "/drink-service" ]
