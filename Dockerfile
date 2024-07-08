FROM golang:1.18

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /smarters

EXPOSE 8080

CMD [ "/smarters" ]
