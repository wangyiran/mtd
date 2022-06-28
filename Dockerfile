FROM golang
MAINTAINER KB
WORKDIR /go/src
COPY . .
RUN go env -w GO111MODULE="on"
RUN go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
RUN chmod +x ./rm.sh
RUN go build
RUN mv myTodoList main
EXPOSE 3100
CMD ./main
