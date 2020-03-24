FROM golang

WORKDIR /app/app

RUN go get github.com/gorilla/mux
RUN go get -u github.com/go-sql-driver/mysql

COPY . .

ENTRYPOINT go run model.go feedDB.go main.go
