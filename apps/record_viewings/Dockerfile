# build stage
FROM golang:1.14 AS build-stage

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o ./record_viewings ./cmd/main

RUN ls -a

# run stage
FROM golang:1.14.2

WORKDIR /app

COPY --from=build-stage /app/record_viewings /app/record_viewings

CMD ["./record_viewings"]