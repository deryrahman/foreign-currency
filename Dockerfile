FROM golang:1-stretch
ADD . /go/src/github.com/deryrahman/foreign-currency
WORKDIR /go/src/github.com/deryrahman/foreign-currency
RUN ["go", "get", "./..."]
RUN ["go", "build", "-o", "main"]
EXPOSE 8000
CMD ["./main"]