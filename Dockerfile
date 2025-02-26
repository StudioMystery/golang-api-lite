FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN go build -o golang-api-lite
EXPOSE 8090
CMD ["./golang-api-lite"]
