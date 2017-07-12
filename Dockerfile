FROM golang:1.8

WORKDIR /go/src/app
COPY main.go main.go 

RUN go-wrapper download
RUN go-wrapper install 

EXPOSE 5000

CMD ["go-wrapper", "run", "app"]
