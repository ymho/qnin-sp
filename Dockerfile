FROM golang:1.20-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
COPY *.crt ./
COPY *.pem ./
RUN go build -o app main.go
CMD ["/app/app"]
EXPOSE 80