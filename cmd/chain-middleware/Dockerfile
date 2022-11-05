FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY chain-middleware /app

ENTRYPOINT ["/app/chain-middleware"]

EXPOSE 50570 50571 50572 50573

CMD ["run"]
