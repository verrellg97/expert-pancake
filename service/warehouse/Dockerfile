FROM alpine:latest

RUN mkdir /app

COPY warehouseService /app

ADD config/warehouse.json /config/warehouse.json

CMD ["/app/warehouseService"]