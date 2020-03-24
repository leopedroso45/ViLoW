FROM golang

WORKDIR /app

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/go-sql-driver/mysql

COPY . .

ENTRYPOINT go run db_connection.go video.go feedDB.go main.go
