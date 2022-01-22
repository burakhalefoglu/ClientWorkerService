# This is a multi-stage build. First we are going to compile and then
# create a small image for runtime.
FROM golang:1.17.6 as builder

RUN mkdir -p /go/src/github.com/client-worker-service
WORKDIR /go/src/github.com/client-worker-service
RUN useradd -u 10001 app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch

COPY --from=builder /go/src/github.com/client-worker-service/main /main
COPY --from=builder /etc/passwd /etc/passwd
USER app

EXPOSE 8000
CMD ["/main/cmd"]