FROM golang:1.17

WORKDIR /dockerapp

COPY .. .

RUN go mod vendor

RUN go build -o /main

EXPOSE 80

CMD ["/main"]