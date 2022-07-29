FROM golang:alpine as builder

WORKDIR /workspace

RUN apk --update add --no-cache ca-certificates && update-ca-certificates

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY cmd/ cmd/
COPY pkg/ pkg/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o xs3 cmd/*.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /workspace/xs3 .

# Run as non root user
USER 1001:1001

ENTRYPOINT ["/xs3"]