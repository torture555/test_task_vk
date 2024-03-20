FROM ubuntu:latest
LABEL authors="pavel"

FROM golang:1.21
WORKDIR /app
COPY / ./
RUN go build

ENV HOST_DB="localhost"
ENV PORT_DB=5432
ENV TYPE_DB="postgres"
ENV LOGIN_DB=""
ENV PASSWD_DB=""
ENV NAME_DB="default"

ENV TIMEOUT_CHECK=60
ENV COUNT_CHECK=10

RUN chmod +x ./start.sh
ENTRYPOINT "./start.sh"

EXPOSE 8080
EXPOSE $PORT_DB