FROM golang:1.24-alpine AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o fizzbuzz-server .

FROM alpine:3.21
COPY --from=build /app/fizzbuzz-server /usr/local/bin/
EXPOSE 8080
CMD ["fizzbuzz-server"]
