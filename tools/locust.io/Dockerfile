FROM alpine:3.9

RUN apk -U add build-base \
               ca-certificates \
               python3 \
               python3-dev \
               py3-zmq && \
    python3 -m pip install locustio && \
    apk del python3-dev && \
    rm -r /var/cache/apk/*

WORKDIR /tools

ONBUILD ADD . /tools
ONBUILD RUN test -f requirements.txt && pip install -r requirements.txt; exit 0

EXPOSE 8089 5557 5558
ENTRYPOINT ["/usr/bin/locust"]