FROM alpine:latest

RUN mkdir /app

COPY businessRelationService /app

ADD config/business_relation.json /config/business_relation.json

CMD ["/app/businessRelationService"]