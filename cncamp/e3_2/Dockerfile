FROM golang:1.16-alpine AS build

COPY golang /go/src/golang
WORKDIR /go/src/golang/cncamp/e2_2
RUN go build -o /bin/cncamp/e2_2/httpserver

FROM golang:1.16-alpine
COPY --from=build /bin/cncamp/e2_2/httpserver /bin/cncamp/e2_2/
EXPOSE 80
ENTRYPOINT [ "/bin/cncamp/e2_2/httpserver" ]