FROM golang:1.15.7
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get gopkg.in/yaml.v2
RUN go build -o main .
CMD ["/app/main"]