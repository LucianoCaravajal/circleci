FROM golang:1.15.3-alpine3.12 as builder

ENV GOOS=linux
ENV CGO_ENABLED=0
ENV PORT 80

WORKDIR /build
COPY app.go .

# Install dependencies
RUN apk add --update git

RUN go get -v -u github.com/gorilla/mux && \
    go build -o go-test-app .

FROM scratch

COPY --from=builder /build/go-test-app /go-test-app

CMD ["/go-test-app"]
