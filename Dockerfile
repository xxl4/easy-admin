FROM alpine

RUN apk update --no-cache
RUN apk add --update gcc g++ libc6-compat
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

WORKDIR /easy-admin

COPY ./easy-admin  /easy-admin
COPY ./config/settings.demo.yml /config/settings.yml
EXPOSE 8000
RUN  chmod +x /easy-admin

ENV TINI_VERSION v0.19.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /tini
RUN chmod +x /tini
ENTRYPOINT ["/tini", "--"]

CMD ["/easy-admin","server","-c", "/config/settings.yml"]