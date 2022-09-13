# syntax=docker/dockerfile:1

FROM golang:1.16-alpine as builder

ADD . /go/src/dealljobs
WORKDIR /go/src/dealljobs

COPY . /go/src/dealljobs

RUN go mod download

RUN go build -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /go/src/dealljobs/app .
COPY --from=builder /go/src/dealljobs/.env .

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD [ "./app" ]