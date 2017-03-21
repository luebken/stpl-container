FROM scratch
ADD stplsrv /
EXPOSE 8080
CMD ["/stplsrv"]