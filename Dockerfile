# Download dependencies
FROM golang:1.13 AS modules

ADD go.mod go.sum /m/
RUN cd /m && go mod download

# Make a builder
FROM golang:1.13 AS builder

# add a non-privileged user
RUN useradd -u 10001 myapp

COPY --from=modules /go/pkg/mod /go/pkg/mod

RUN mkdir -p /go-xke-jan-2020
ADD . /go-xke-jan-2020
WORKDIR /go-xke-jan-2020
ENV PROJECT github.com/dicaormu/go-xke-jan-2020

# Build the binary with go build
RUN CGO_ENABLED=0 go build -o bin/go-xke ${PROJECT}/cmd/go-xke

# Final stage: Run the binary
FROM scratch

ENV PORT 8080

# certificates to interact with other services
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# don't forget /etc/passwd from previous stage
COPY --from=builder /etc/passwd /etc/passwd
USER myapp

# and finally the binary
COPY --from=builder /go-xke-jan-2020/bin/go-xke /go-xke
EXPOSE $PORT

CMD ["/go-xke"]