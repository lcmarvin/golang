FROM golang:1.16-alpine AS build

COPY ./ /go/src/golang
WORKDIR /go/src/golang/cncamp/e10_hw
RUN export GO111MODULE=on && \
  export GOPROXY=https://goproxy.cn,direct && \
  go build -o /bin/cncamp/e10_hw/httpserver

FROM golang:1.16-alpine
COPY --from=build /bin/cncamp/e10_hw/httpserver /bin/cncamp/e10_hw/
EXPOSE 80
ENTRYPOINT [ "/bin/cncamp/e10_hw/httpserver" ]