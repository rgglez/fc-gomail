FROM alpine:3.17.2

ENV container docker

WORKDIR /usr/src/app

COPY src/sendmail/src/handler .

CMD ["./handler"]