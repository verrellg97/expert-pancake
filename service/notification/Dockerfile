FROM alpine:latest

RUN mkdir /app

COPY notificationService /app

ADD config/notification.json /config/notification.json

CMD ["/app/notificationService"]