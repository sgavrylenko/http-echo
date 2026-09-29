FROM --platform=$BUILDPLATFORM golang:1.27-alpine AS builder
ARG TARGETOS
ARG TARGETARCH
WORKDIR /app
COPY go.mod ./
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o echo-headers .

FROM alpine:3.24.2
WORKDIR /app
COPY --from=builder /app/echo-headers .

# Порт за замовчуванням всередині контейнера
ENV PORT=8000

CMD ["./echo-headers"]
