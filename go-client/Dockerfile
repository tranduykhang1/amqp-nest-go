FROM golang:1.16

WORKDIR /app/go-client

COPY go.mod .
COPY go.sum .

RUN go mod tidy
RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/go-client .


EXPOSE 8080

CMD ["./out/go-client"]