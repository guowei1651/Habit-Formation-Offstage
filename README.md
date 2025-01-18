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