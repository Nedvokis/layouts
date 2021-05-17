FROM golang:1.16.0-alpine3.13
RUN mkdir /src 
ADD . /src
WORKDIR /src
RUN apk --no-cache add ca-certificates
COPY . .
RUN go build -o main .
CMD ["/src/main"]