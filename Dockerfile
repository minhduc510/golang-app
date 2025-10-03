FROM golang:1.22.2-alpine

WORKDIR /app
COPY . .
EXPOSE 3000

CMD ["go", "run", "."]