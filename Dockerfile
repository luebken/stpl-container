FROM alpine:3.5

RUN apk add --update -U ca-certificates \
    && rm -rf /var/cache/apk/*

ADD stplsrv /
EXPOSE 8080
CMD ["/stplsrv"]