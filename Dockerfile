#First stage
FROM golang:alpine as builder

RUN apk update && apk upgrade

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o api


# Second stage
FROM apline

LABEL maintainer="vrbalu00@students.zcu.cz"

COPY --from builder /build /app/
ENV GIN_MODE=release
ENV LOG_LEVEL=ERROR

ENTRYPOINT ["app/api"]