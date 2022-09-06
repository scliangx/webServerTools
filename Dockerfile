FROM alpine:latest as prod

WORKDIR /server/apps/
ENV TZ Asia/Shanghai
COPY ./apps /server/apps
RUN apk -y install vim && yum -y install net-tools
#RUN apk --no-cache add ca-certificates && update-ca-certificates
RUN apk --no-cache add ca-certificates