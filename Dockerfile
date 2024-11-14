FROM golang:1.18-alpine

WORKDIR /

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

CMD ["/cmd/jogodavelha2"]

EXPOSE 8080
