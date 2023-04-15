FROM alpine:latest

RUN mkdir /app

COPY salesService /app

ADD config/sales.json /config/sales.json

CMD ["/app/salesService"]