#FROM golang:alpine 
# FROM ubuntu
FROM alpine
WORKDIR /go_users
WORKDIR /go_users/logs
EXPOSE 8020
MAINTAINER wjp <973513569@qq.com>
VOLUME /go_users/logs
#tar gzip bzip2 xz
#ADD users-go.gzip users-go.gzip

COPY users-go /go_users/users-go

#COPY config-format.xml /go_users/config-format.xml 
#COPY config-format.json /go_users/config-format.json
#COPY config-format.json /go_users/config.json
COPY config-format.xml /go_users/config.xml 
#COPY users-go users-go
#COPY config-format.json config.json
COPY config-format.xml config.xml
# 运行g
# RUN bash -c 'touch /go_users/users-go'
#ENV WORKDIR /Open-IM-Server
#RUN mkdir $WORKDIR/go_users/logs 
#RUN rm -r -f  /go_users/logs  && mkdir /go_users/logs 

RUN chmod +x /go_users/users-go
#&& cd /go_users/users-go


#ENV port=8020
#pass
#ENTRYPOINT ["java", "-jar", "eureka-server.jar", "--spring.profiles.active=test", "--server.port=4000", "> /log/app.log"]
#ENTRYPOINT ["sh","-c","/go_users/users-go -Dport=$port   -Dfetch-registry=${fetch-registry}", "> /logs/go_users.log"]
#users-go: not found
ENTRYPOINT ["sh","-c","/go_users/users-go", "> /logs/go_users.log"]

# 有缓存 手动删除 
#docker build -t go_users:v1 .