# Habit-Formation-Offstage
iot, openapi


## build

1. 启动编译镜像
```bash
docker container run -it --rm --name golang_build --hostname golang_build \
  -v /home/wales/growing_iot/Habit-Formation-Offstage:/home/wales \
  -w /home/wales \
  golang:1.24rc1-alpine3.21 sh
```

```
docker container run -it --rm --name golang_build --hostname golang_build \
  -v /home/wales/growing_iot/Offstage:/home/wales \
  -w /home/wales \
  golang:1.24rc1-alpine3.21 sh
```


2. 配置Go模块代理
```bash
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
```

3. 加载配置
```bash
go mod init hf
go mod tidy

go get github.com/spf13/cobra
go get github.com/lib/pq
go get github.com/emicklei/go-restful-openapi/v2
go get github.com/emicklei/go-restful/v3
go get github.com/go-openapi/spec
```

4. 编译
```bash
go build -o hf main.go
```

## 启动

```bash
docker container run -d --name growing-iot --hostname growing-iot \
  -p 8090:8090 \
  -v /home/wales/growing_iot/Offstage:/growing \
  -w /growing \
  alpine:3.21 ./hf
```

docker container run -it --name growing-iot --hostname growing-iot \
  -p 8090:8090 \
  -v /home/wales/growing_iot/Offstage:/growing \
  -w /growing \
  alpine:3.21 sh

## 测试

1. 登录测试

```bash
curl 'http://172.25.1.22:8090/users/login' -X POST \
  -H 'Accept: */*' -H 'Content-Type: application/json;charset=UTF-8' \
  --data-raw '{"username": "xxxxx", "password": "xxxxxx"}'
```

2. 获取轮播

```bash
curl 'http://172.25.1.22:8090/carousels' -X GET \
  -H 'HF-User-Id: xxxxxx' \
  -H 'Accept: */*' -H 'Content-Type: application/json;charset=UTF-8'
```

3. 获取轮播项
```bash
curl 'http://172.25.1.22:8090/carousels/1/items' -X GET \
  -H 'HF-User-Id: xxxxxx' \
  -H 'Accept: */*' -H 'Content-Type: application/json;charset=UTF-8'
```

4. 习惯计数
```bash
curl 'http://172.25.1.22:8090/habit/1/record' -X GET \
  -H 'HF-User-Id: xxxxxx' \
  -H 'Accept: */*' -H 'Content-Type: application/json;charset=UTF-8' \
  --data-raw '{ "type": 3, "serial" : "browser", "serial": "",  "remark": "" }'
```