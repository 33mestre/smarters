FROM golang:1.18-alpine

WORKDIR /app

# Instalar gcc e musl-dev
RUN apk add --no-cache gcc musl-dev

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /smarters

EXPOSE 8080

CMD ["/smarters"]
