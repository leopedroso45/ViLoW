FROM golang:latest

WORKDIR /app

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/gorilla/securecookie
RUN go get -u github.com/go-sql-driver/mysql

COPY . .

RUN chmod +x wait-for-it.sh

ENTRYPOINT go run db_connection.go ip.go UpdateData.go Vote.go video.go user.go Comment.go VideoPage.go handler.go session.go feedDB.go main.go
