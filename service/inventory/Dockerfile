FROM alpine:latest

RUN mkdir /app

COPY inventoryService /app

ADD config/inventory.json /config/inventory.json

CMD ["/app/inventoryService"]