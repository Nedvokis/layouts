FROM golang:1.16.0-alpine3.13
RUN mkdir /src 
ADD . /src
WORKDIR /src
COPY go.mod go.sum ./
RUN go build -o main .
CMD ["/src/main"]