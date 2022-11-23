FROM alpine:latest

RUN mkdir /app

COPY accountingService /app

ADD config/accounting.json /config/accounting.json

CMD ["/app/accountingService"]