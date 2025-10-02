# ---- build stage ----
FROM golang:1.24.2 AS build
WORKDIR /app

# No deps to download; just copy sources and build
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/scicalc ./cmd/server

# ---- runtime stage ----
FROM gcr.io/distroless/static:nonroot
COPY --from=build /bin/scicalc /usr/local/bin/scicalc
USER nonroot:nonroot
ENTRYPOINT ["/usr/local/bin/scicalc"]
# Optionally: CMD ["--help"]