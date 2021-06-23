FROM node as ui
WORKDIR /app
COPY ui/package*.json ./
RUN npm install
COPY ui/ .
RUN npm run build

FROM golang as server
WORKDIR /app
COPY go.* ./
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go mod download
COPY . .
COPY --from=ui /app/dist/ ./ui/dist/
RUN go build

FROM debian
COPY --from=server /app/web-file-manager /
CMD ["/web-file-manager","-d","/data"]