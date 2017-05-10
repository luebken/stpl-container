FROM scratch

ADD ca-certificates.crt /etc/ssl/certs/

ADD stplsrv /
EXPOSE 8080

# Environment:
#  * $REDIS_URL     - redis connection. defaults to redis://localhost:6379
LABEL environment.optional="REDIS_URL"
#  * $LIBRARIES_IO_API_KEY     - https://libraries.io/
LABEL environment.required="LIBRARIES_IO_API_KEY"

CMD ["/stplsrv"]