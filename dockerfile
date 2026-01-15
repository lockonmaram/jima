# ---------- Build stage ----------
FROM golang:1.25.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o jima

# ---------- Runtime stage ----------
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/jima /app/jima
COPY --from=builder /app/view /app/view

EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/app/jima"]
