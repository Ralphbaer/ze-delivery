FROM golang:1.8

WORKDIR /app

ADD ./.bin/partner-service .
ADD ./gen/swagger.yaml ./gen/

RUN groupadd -r appusers && useradd --no-log-init -r -g appusers appuser

EXPOSE 3000

ENTRYPOINT ["./partner-service"]
