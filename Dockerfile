FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o tender ./cmd/tender/main.go

EXPOSE 8080
CMD ["./tender"]