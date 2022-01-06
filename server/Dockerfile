FROM golang:latest as BUILDER

MAINTAINER zengchen1024<chenzeng765@gmail.com>

# build binary
WORKDIR /go/src/github.com/opensourceways/repo-owners-cache
COPY . .
RUN cd server && GO111MODULE=on CGO_ENABLED=0 go build -a -o server .

# copy binary config and utils
FROM alpine:3.14
COPY  --from=BUILDER /go/src/github.com/opensourceways/repo-owners-cache/server/server /opt/app/repo-owners-cache

ENTRYPOINT ["/opt/app/repo-owners-cache"]
