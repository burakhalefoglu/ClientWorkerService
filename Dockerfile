FROM golang:1.18.0-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/client-worker-service .
RUN rm -rf /app/internal
RUN rm -rf /app/pkg
RUN rm -rf /app/main.go
RUN rm -rf /app/go.mod
EXPOSE 8000

CMD [ "/app/client-worker-service" ]