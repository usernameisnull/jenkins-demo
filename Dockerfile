FROM release.daocloud.io/demo/ubuntu:22.04
COPY simple-http-server /
CMD  [ "/simple-http-server" ]
