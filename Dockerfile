FROM golang

WORKDIR /app/VideoLocalManager

RUN go get github.com/gorilla/mux
RUN go get github.com/leopedroso45/VideoLocalManager/db

COPY . .

ENTRYPOINT go run model.go feedDB.go main.go
