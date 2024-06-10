FROM golang:1.22.2

WORKDIR /golang

COPY go.mod go.sum /golang/

RUN go mod download

COPY . /golang/

RUN go build -o main .

CMD ["./main"]