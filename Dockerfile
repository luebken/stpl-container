FROM scratch

ADD ca-certificates.crt /etc/ssl/certs/

ADD stplsrv /
EXPOSE 8080
CMD ["/stplsrv"]