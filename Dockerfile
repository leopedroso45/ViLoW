FROM golang

WORKDIR /app/VideoLocalManager

RUN go get github.com/gorilla/mux
RUN go get github.com/leopedroso45/VideoLocalManager/db
RUN go get -u github.com/go-sql-driver/mysql

COPY . .

ENTRYPOINT go run model.go feedDB.go main.go
