FROM golang:1.16.0-alpine3.13
RUN mkdir /src 
ADD . /src
WORKDIR /src
RUN apk add --no-cache ca-certificates && update-ca-certificates
RUN go build -o main .
CMD ["/src/main"]