FROM golang:1.16.4 as server
RUN go env -w GO111MODULE="on"
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go get github.com/myml/web-file-manager

FROM debian
COPY --from=server /go/bin/web-file-manager /
CMD ["/web-file-manager","-d","/data"]