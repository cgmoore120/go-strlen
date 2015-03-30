FROM scratch

MAINTAINER Chris Moore <chris@cloudspace.com>

WORKDIR /

ADD go-strlen go-strlen

ADD ./microservice.yml microservice.yml
