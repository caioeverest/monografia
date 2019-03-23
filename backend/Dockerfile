#Builder
FROM golang:1.11.0-stretch as builder
COPY . /App
WORKDIR /App
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o App

#Monta container compilado
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /App .
CMD ["./App"]