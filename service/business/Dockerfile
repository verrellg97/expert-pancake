FROM alpine:latest

RUN mkdir /app

COPY businessService /app

ADD config/business.json /config/business.json

CMD ["/app/businessService"]