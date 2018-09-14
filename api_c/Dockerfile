FROM alpine
LABEL MAINTAINER="xushikuan@mozat.com"

ENV ENVIROMENT_PROFILES_ACTIVE dt
ENV TIME_ZONE=Asia/Singapore
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone
RUN apk add --no-cache tzdata

##CA证书，https请求
RUN apk add --no-cache ca-certificates

WORKDIR /go
COPY main /go
COPY . /go/api_c


ENTRYPOINT ./main -enviroment=$ENVIROMENT_PROFILES_ACTIVE