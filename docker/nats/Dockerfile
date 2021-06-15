FROM nats:latest

LABEL maintainer="Massimo Costa <massimoc@filewave.com>"

COPY .empty /certs/

VOLUME ["/certs"]

# COPY server.* certs/
CMD ["--tls", "--tlscert=/certs/server.crt", "--tlskey=/certs/server.key"]