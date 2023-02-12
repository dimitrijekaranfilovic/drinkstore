FROM golang:1.18-alpine3.15

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY main.go ./
COPY database/*.go ./database/
COPY handlers/*.go ./handlers/
COPY model/*.go ./model/
COPY repository/*.go ./repository/

RUN go build -o /user-service

EXPOSE 8081

# CMD ["../scripts/wait-for-it.sh", "ntp-database:5432",  "--" ,"/user-service"]
CMD ["/user-service"]