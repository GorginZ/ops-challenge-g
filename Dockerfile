FROM golang:1.16-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
ARG SECRET
ENV SECRET=$SECRET
RUN go build -o main .
CMD ["/app/main"]