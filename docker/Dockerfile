FROM golang:1.14.1 as build-img
ENV GO111MODULE=off
RUN apk update && apk add git && mkdir /home/anglebroking
COPY ${PWD}/ /go/src/github.com/andanhm/anglebroking
WORKDIR /go/src/github.com/andanhm/anglebroking
RUN git rev-parse HEAD > /home/anglebroking/version
RUN CGO_ENABLED=0 go build -a -ldflags "-s -w" -o /home/anglebroking/anglebroking cmd/*.go

FROM registry.bookmyshow.org/alpine:latest

RUN apk --no-cache add ca-certificates
COPY --from=build-img /home/anglebroking/** /home/
WORKDIR /home/
EXPOSE 80
CMD ["sh","-c","./anglebroking"]
