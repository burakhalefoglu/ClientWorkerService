FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /client-worker-service ./cmd

EXPOSE 8000

CMD [ "/client-worker-service" ]