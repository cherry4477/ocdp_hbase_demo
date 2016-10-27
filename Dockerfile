FROM golang:1.6.2

ENV TIME_ZONE=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

COPY . /go/src/github.com/asiainfoLDP/ocdp_hbase_demo

WORKDIR /go/src/github.com/asiainfoLDP/ocdp_hbase_demo

RUN go build
COPY krb5.conf /etc/
RUN apt-get update   && \
    apt-get install -y --no-install-recommends  krb5-user krb5-config

COPY start.sh /hadoop_demo/

CMD ./start.sh
