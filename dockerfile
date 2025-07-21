FROM golang:1.24 
FROM postgres

WORKDIR /app

COPY . .

RUN go build -o main ./cmd

CMD [ "./main" ]