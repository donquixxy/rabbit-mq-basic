# builder image
FROM golang:1.21rc4-alpine as builder
WORKDIR /build
COPY . .
RUN sed -i s/https/http/ /etc/apk/repositories
RUN apk add curl
RUN apk add git && CGO_ENABLED=0 GOOS=linux go build -o go-microservices .
# RUN go build -o go-microservices .

# generate clean, final image for end users
FROM alpine
RUN apk update && apk add ca-certificates && apk add tzdata && apk add git && apk add curl
COPY --from=builder /build .
ENV TZ="Asia/Makassar"
EXPOSE 8213

CMD ./go-microservices