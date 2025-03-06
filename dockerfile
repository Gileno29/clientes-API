FROM golang:1.23.7-alpine3.21 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

#DESABILITA OS COPILADORES DO C QUE NÃO ESTÁ PRESENTE NA IMAGEM FINAL
RUN CGO_ENABLED=0 GOOS=linux go build -o apiclientes

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app .

EXPOSE 8080

CMD ["./apiclientes"]