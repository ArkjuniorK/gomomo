# syntax=docker/dockerfile:1

# build stage
FROM golang:1.20 AS Builder
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app .

# run stage
FROM alpine:latest
RUN apk upgrade -U \
  && apk add ca-certificates \
  && rm -rf /var/cache/*
WORKDIR /app/
COPY --from=Builder /go/src/app ./
EXPOSE 3000
ENTRYPOINT [ "./app" ]