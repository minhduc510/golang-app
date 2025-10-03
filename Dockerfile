FROM golang:1.24-alpine

WORKDIR /app
COPY . .
EXPOSE 3000

RUN go mod tidy

CMD ["go", "run", "."]