FROM alpine:latest

RUN mkdir /app

COPY accountService /app

ADD config/account.json /config/account.json

CMD ["/app/accountService"]