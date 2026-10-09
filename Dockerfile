# Stage 1: Build aplikasi Go
FROM golang:1.27.1-alpine AS builder

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux \
    go build -trimpath -ldflags="-s -w" \
    -o /out/belajar-cicd .

# Stage 2: Runtime minimal
FROM scratch

COPY --from=builder /out/belajar-cicd /belajar-cicd

USER 65532:65532

EXPOSE 8080

ENTRYPOINT ["/belajar-cicd"]
