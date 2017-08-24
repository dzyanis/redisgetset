FROM golang
ADD . /go/src/github.com/dzyanis/redisgetset
RUN go get github.com/go-redis/redis
RUN go install github.com/dzyanis/redisgetset
ENTRYPOINT /go/bin/redisgetset
EXPOSE 8080