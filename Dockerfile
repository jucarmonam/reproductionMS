FROM golang:1.17

WORKDIR /application

COPY go.mod ./
COPY go.sum ./
COPY .env ./

RUN go mod download

COPY main.go ./
COPY ./app/ ./app
COPY ./database/ ./database

RUN go build -o application

CMD ./application
