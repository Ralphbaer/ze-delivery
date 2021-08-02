FROM golang as builder

ARG MODULE

RUN apt-get update \
    && apt-get install -y jq

RUN curl -sSL "https://github.com/gotestyourself/gotestsum/releases/download/v0.3.1/gotestsum_0.3.1_linux_amd64.tar.gz" | tar -xz -C /usr/local/bin gotestsum 

CMD ["/bin/sh"]