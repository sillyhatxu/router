FROM alpine
LABEL MAINTAINER="xushikuan@mozat.com"

ENV ENVIROMENT_PROFILES_ACTIVE dt
ENV CONSUL_ADDRESS localhost
ENV MODULE_NAME deja
ENV PORT 80

ENV TIME_ZONE=Asia/Singapore
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone
RUN apk add --no-cache tzdata

##CA证书，https请求
RUN apk add --no-cache ca-certificates

WORKDIR /go
COPY main /go
COPY . /go/api_a

##-consul-address=172.28.2.106:8500 -module-name=api-a -port=12001
ENTRYPOINT ./main -enviroment=$ENVIROMENT_PROFILES_ACTIVE -consul-address=$CONSUL_ADDRESS -module-name=$MODULE_NAME -port=$PORT