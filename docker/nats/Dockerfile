FROM nats:latest

LABEL maintainer="Massimo Costa <massimoc@filewave.com>"

COPY .empty /certs/
COPY .empty /conf/

COPY nats-server.conf /nats-server.conf

VOLUME ["/certs", "/conf"]

EXPOSE 443

CMD ["--tls", "--tlscert=/certs/server.crt", "--tlskey=/certs/server.key", "--config", "/conf/nats-server.conf"]