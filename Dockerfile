FROM alpine:3.2

ENV GOPATH=/gopath \
  SRC=/gopath/src/github.com/caarlos0/getantibody

WORKDIR $SRC
ADD . $SRC

RUN apk add -U git go && \
  go get -v -d ./... && \
  go get -v github.com/GeertJohan/go.rice/rice && \
  /gopath/bin/rice embed-go -i ./server && \
  go install -v ./... && \
  apk del git go && \
  rm -rf /gopath/src /gopath/pkg /var/cache/apk/*

CMD /gopath/bin/server
