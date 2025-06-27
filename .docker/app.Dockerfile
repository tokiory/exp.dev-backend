FROM golang:1.24-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN apk add --no-cache make

RUN make

EXPOSE 8080

CMD ["./bin/main"]
