FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./partner/*.go ./

COPY Dockerfile ./partner/gen/swagger.yaml* ./gen/

RUN rm ./partner/./gen/Dockerfile
RUN go build -o /ze-delivery

EXPOSE 8080

CMD [ "./partner" ]