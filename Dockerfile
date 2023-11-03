FROM alpine

RUN apk update --no-cache
RUN apk add --update gcc g++ libc6-compat
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
RUN apk add --no-cache tini
ENV TZ Asia/Shanghai


COPY ./easy-admin  /easy-admin
COPY ./config/settings.demo.yml /config/settings.yml
EXPOSE 8000
RUN  chmod +x /easy-admin


ENTRYPOINT ["/sbin/tini", "--"]

CMD ["/easy-admin","server","-c", "/config/settings.yml"]