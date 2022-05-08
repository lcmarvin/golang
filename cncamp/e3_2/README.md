Dockerfile如下：

```dockerfile
FROM golang:1.16-alpine AS build

COPY ./ /go/src/golang
WORKDIR /go/src/golang/cncamp/e2_2
RUN go build -o /bin/cncamp/e2_2/httpserver

FROM golang:1.16-alpine
COPY --from=build /bin/cncamp/e2_2/httpserver /bin/cncamp/e2_2/
EXPOSE 80
ENTRYPOINT [ "/bin/cncamp/e2_2/httpserver" ]
```

构建镜像

```
[root@VM-16-16-centos ~]# git clone https://github.com/lcmarvin/golang.git
正克隆到 'golang'...
remote: Enumerating objects: 53, done.
remote: Counting objects: 100% (53/53), done.
remote: Compressing objects: 100% (27/27), done.
展开对象中: 100% (53/53), 4.79 KiB | 817.00 KiB/s, 完成.
remote: Total 53 (delta 13), reused 53 (delta 13), pack-reused 0
[root@VM-16-16-centos ~]# cd golang
[root@VM-16-16-centos golang]# docker build -f cncamp/e3_2/Dockerfile -t mdeng2022/http-server:latest .
Sending build context to Docker daemon  140.8kB
Step 1/8 : FROM golang:1.16-alpine AS build
 ---> 7e352955f83c
Step 2/8 : COPY ./ /go/src/golang
 ---> c79281f06390
Step 3/8 : WORKDIR /go/src/golang/cncamp/e2_2
 ---> Running in 7771194f405d
Removing intermediate container 7771194f405d
 ---> 2d828ce872cb
Step 4/8 : RUN go build -o /bin/cncamp/e2_2/httpserver
 ---> Running in 69af24a613f4
Removing intermediate container 69af24a613f4
 ---> b086de4951d9
Step 5/8 : FROM golang:1.16-alpine
 ---> 7e352955f83c
Step 6/8 : COPY --from=build /bin/cncamp/e2_2/httpserver /bin/cncamp/e2_2/
 ---> 96034981699a
Step 7/8 : EXPOSE 80
 ---> Running in 4416edf08664
Removing intermediate container 4416edf08664
 ---> db28a04aff56
Step 8/8 : ENTRYPOINT [ "/bin/cncamp/e2_2/httpserver" ]
 ---> Running in 4a0714eb2c4d
Removing intermediate container 4a0714eb2c4d
 ---> 9878af419b16
Successfully built 9878af419b16
Successfully tagged mdeng2022/http-server:latest
```

推送镜像到Docker Hub

```
[root@VM-16-16-centos golang]# docker login
Authenticating with existing credentials...
WARNING! Your password will be stored unencrypted in /root/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded
[root@VM-16-16-centos ~]# docker push mdeng2022/http-server
Using default tag: latest
The push refers to repository [docker.io/mdeng2022/http-server]
ea7aef2d88d9: Layer already exists
19c4d4cefc09: Layer already exists
46e96c819e17: Layer already exists
b6f786c730a9: Layer already exists
63a6bdb95b08: Layer already exists
8d3ac3489996: Layer already exists
latest: digest: sha256:2c894778cce094dbb01e05f85ea3b8b33f49558b06b73b54c25a7481b0cbfeff size: 1576
```

启动httpserver

```
[root@VM-16-16-centos golang]# docker run -d -P mdeng2022/http-server:latest
1fb832d8589c0438f6e9b1abe2b22219621f72427a25222fa5890ccf2017b808
```

进入容器查看IP配置

```
[root@VM-16-16-centos golang]# docker inspect 1fb832d858|grep Pid
            "Pid": 2938394,
            "PidMode": "",
            "PidsLimit": null,
[root@VM-16-16-centos golang]# nsenter -t 2938394 -n ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
65: eth0@if66: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```