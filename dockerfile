FROM golang:1.24-alpine

RUN go version 

ENV GOPATH=/

COPY ./ ./ 

RUN go mod download
RUN go build -o url ./cmd/main.go


CMD [ "./url" ]