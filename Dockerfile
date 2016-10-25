FROM golang:1.6.2

ENV TIME_ZONE=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

COPY . /go/src/github.com/asiainfoLDP/ocdp_hbase_demo

WORKDIR /go/src/github.com/asiainfoLDP/ocdp_hbase_demo

RUN go build

CMD ["sh", "-c", "./ocdp_hbase_demo"]