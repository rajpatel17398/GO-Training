#  base image that is golang letest version
FROM golang:1.17

# label has the detail of the maintainer
LABEL maintainer="Raj Patel" 

# current working directory inside the container
WORKDIR /app

# copy the module and sum file for the dependencies
COPY go.mod .
COPY go.sum .


# download go mod file
RUN go get github.com/dgrijalva/jwt-go
RUN go mod download

# we need to copy the files
COPY . .

# we need to specify the port number
EXPOSE 8080

# now we habe to build the application
RUN go build

# once container is initialize then the code that we want to run
CMD ["./jwt"]
