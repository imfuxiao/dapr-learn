FROM golang:1.16-alpine as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o pubsub

FROM alpine:3.9.4
COPY --from=builder ["/app/pubsub", "/bin"]
CMD [ "/bin/pubsub" ]