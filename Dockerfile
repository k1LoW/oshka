FROM docker:latest

RUN apk add --no-cache bash curl git

ENTRYPOINT ["/entrypoint.sh"]

COPY scripts/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

COPY oshka_*.apk /tmp/
RUN apk add --allow-untrusted /tmp/oshka_*.apk
