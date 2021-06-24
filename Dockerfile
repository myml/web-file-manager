FROM golang:1.16.4 as server
WORKDIR /app
COPY go.* ./
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
COPY . .
RUN go build

FROM debian
COPY --from=server /app/web-file-manager /
CMD ["/web-file-manager","-d","/data"]