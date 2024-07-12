FROM golang:1.22.0 AS builder
WORKDIR /user/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o payment ./cmd/main.go

FROM scratch
COPY --from=builder /usr/src/app/payment ./payment
CMD ["./payment"]