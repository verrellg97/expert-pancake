FROM alpine:latest

RUN mkdir /app

COPY purchasingService /app

ADD config/purchasing.json /config/purchasing.json

CMD ["/app/purchasingService"]