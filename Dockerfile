FROM golang:latest

WORKDIR /app

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/gorilla/securecookie
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u golang.org/x/oauth2
RUN go get -u cloud.google.com/go/compute/metadata

COPY . .

RUN chmod +x wait-for-it.sh

ENTRYPOINT go run .
