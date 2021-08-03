FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./partner/*.go ./

RUN go build -o /ze-delivery

EXPOSE 8080

CMD [ "./partner" ]