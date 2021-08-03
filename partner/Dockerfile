FROM golang:1.8

WORKDIR /app

RUN go build -o /ze-delivery/partner/.bin/partner-service

ADD ./partner/.bin/partner-service .
ADD ./partner/gen/swagger.yaml ./partner/gen/

RUN groupadd -r appusers && useradd --no-log-init -r -g appusers appuser

EXPOSE 8080

CMD [ "./partner" ]