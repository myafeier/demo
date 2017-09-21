FROM alpine:3.6


# https://golang.org/issue/14851 (Go 1.8 & 1.7)
# https://golang.org/issue/17847 (Go 1.7)
#COPY *.patch /go-alpine-patches/

RUN	mkdir /go
COPY testbin/test /go
WORKDIR /go
VOLUME ["/go"]
EXPOSE 8081
ENTRYPOINT ["/go/test"]



