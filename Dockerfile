FROM golang:alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app
COPY go.mod .
RUN go get
#COPY go.sum .
RUN go mod download
RUN go build -o main .
RUN adduser -S -D -H -h /app appuser
USER appuser
ENTRYPOINT ["./main"]