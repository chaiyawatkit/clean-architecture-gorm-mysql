FROM golang:1.21 AS builder
WORKDIR /srv/go-app
COPY . .
RUN go build -o clean-architecture-gorm-mysql


FROM golang:1.21
WORKDIR /srv/go-app
COPY --from=builder /srv/go-app/config.json .
COPY --from=builder /srv/go-app/clean-architecture-gorm-mysql .

CMD ["./clean-architecture-gorm-mysql"]