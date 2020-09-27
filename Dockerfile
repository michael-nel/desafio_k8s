FROM golang:1.14

COPY ./go/ ./

RUN go test -v -run TestGreetings

RUN go build -o main

EXPOSE 8000

ENTRYPOINT [ "./main" ] 