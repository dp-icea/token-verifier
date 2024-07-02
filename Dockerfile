FROM golang:1.22-alpine AS build
RUN apk add git bash make
RUN mkdir /app
WORKDIR /app
COPY . .

# Get dependencies - will also be cached if we won't change mod/sum
RUN go mod download

RUN go install ./...

FROM alpine:latest
COPY --from=build /go/bin/token-verifier /usr/bin
COPY .env.example .env
RUN mkdir /certs
COPY certs/auth2.pem /cert.pem
ENTRYPOINT ["/usr/bin/token-verifier"]