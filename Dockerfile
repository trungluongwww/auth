# Stage 1: build application
FROM golang:1.23 as builder

WORKDIR /app

# copy dependencies files first for better caching
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# copy src
COPY . .
# remove unnecessary files
RUN rm mysql/ddl/*
# build
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /app/app ./cmd

# Stage 2: create a minimal final image
FROM alpine:latest

WORKDIR /

# copy app
COPY --from=builder /app/app /app
RUN chmod +x /app

# copy DDL files
COPY mysql/ddl/ mysql/ddl/

CMD ["/app"]