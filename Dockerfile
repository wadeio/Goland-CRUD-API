FROM golang

RUN mkdir -p /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

# app
COPY . .
RUN go build -o ./app ./main.go

EXPOSE 8000
ENTRYPOINT ["./app"]