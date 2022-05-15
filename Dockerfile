FROM golang:1.18
Add . /dog
WORKDIR /dog
RUN GOPROXY=https://goproxy.io
RUN go mod download
RUN go build -o dog
CMD ["./dog"]
