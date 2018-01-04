FROM golang:latest

ADD kubeless.go /
RUN  go build -buildmode=plugin *.go
EXPOSE 8080
CMD ["go","run","/kubeless.go"]