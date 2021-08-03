FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go get github.com/Ralphbaer/ze-delivery/common

COPY ./partner/*.go ./

COPY ./partner/gen/swagger.yaml* ./partner/gen/

RUN go build -o /ze-delivery

EXPOSE 8080

CMD [ "./partner" ]