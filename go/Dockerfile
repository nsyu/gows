FROM golang:1.14

RUN apt-get update
RUN apt-get install vim -y
RUN go get -u github.com/gocolly/colly

WORKDIR /go/src/app

COPY . .

CMD ["go", "run", "main.go"]
# RUN go install -v ./...
# CMD ["app"]