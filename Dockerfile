FROM ubuntu:20.04

WORKDIR /shorturl_app
COPY shorturl /shorturl_app/
EXPOSE 8080

CMD ["./shorturl"]