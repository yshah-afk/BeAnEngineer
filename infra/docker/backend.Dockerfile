# ---- Build Stage ----
FROM golang:1.25-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /server ./cmd/server

# ---- Run Stage ----
FROM gcr.io/distroless/static-debian12

COPY --from=build /server /server

EXPOSE 8080

ENTRYPOINT ["/server"]
