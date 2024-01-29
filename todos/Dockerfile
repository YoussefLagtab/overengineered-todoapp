FROM golang:alpine3.19 as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build ./cmd/todos

# ------------------------
FROM alpine:3.19

WORKDIR /app
COPY --from=build /app/todos .

CMD ["./todos"]
