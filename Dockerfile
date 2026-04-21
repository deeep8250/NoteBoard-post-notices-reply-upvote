FROM golang:1.25-alpine as builder


WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o threadpulse ./cmd/api/


FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/threadpulse .

EXPOSE 8080
CMD [ "./threadpulse" ]