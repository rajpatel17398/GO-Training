#  base image that is golang letest version
FROM golang:1.17

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 9000

RUN go build
CMD ["./rest"]
