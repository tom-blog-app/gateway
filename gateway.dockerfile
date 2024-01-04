FROM alpine:latest

RUN mkdir /app

COPY gateway /app

CMD [ "/app/gateway"]