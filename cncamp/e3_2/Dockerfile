FROM golang:1.16-alpine AS build

COPY golang /go/src/golang
WORKDIR /go/src/golang/cncamp/e2_2
RUN go build -o /bin/cacamp/e2_2

FROM golang:1.16-alpine
COPY --from=build /bin/cacamp/e2_2 /bin/cacamp/e2_2
ENTRYPOINT [ "/bin/cacamp/e2_2/e2_2" ]