# xrobot项目

### 一、附表

1.[错误码表](./docs/table/code.md)

### 二、开发工具

1.安装gorm代码生成工具

```shell
go install github.com/dobyte/gorm-dao-generator@latest
```

2.安装rpcx代码生成工具

```shell
go install github.com/rpcxio/protoc-gen-rpcx@latest
```

3.安装protobuf go代码生成工具

```shell
go install github.com/gogo/protobuf/protoc-gen-gofast@latest
```

4.安装swagger文档生成工具

```shell
go install github.com/swaggo/swag/cmd/swag@latest
```

5.下载protobuf编译器

- Linux, using apt or apt-get, for example:

```shell
$ apt install -y protobuf-compiler
$ protoc --version  # Ensure compiler version is 3+
```

- MacOS, using Homebrew:

```shell
$ brew install protobuf
$ protoc --version  # Ensure compiler version is 3+
```

- Windows, download from [Github](https://github.com/protocolbuffers/protobuf/releases):

### 三、项目部署
1.进入项目运行目录

```shell
$ cd /workspace/xrobot_bin
```

2.停止所有运行的项目

```shell
$ ./deploy.sh stop all
```

3.进入项目源码目录

```shell
$ cd /workspace/xrobot_server
```

4.拉取最新的代码

```shell
$ git pull origin main
```

5.执行编译

```shell
$ ./deploy.sh make local all
```

6.进入项目运行目录

```shell
$ cd /workspace/xrobot_bin
```

7.启动所有项目

```shell
$ ./deploy.sh start all
```

8.额外补充，非必要步骤（当due框架代码更新后，需进入due框架目录拉取最新的框架代码）
```shell
$ cd /workspace/xbase
$ git pull origin v2-feature-main
```

### 四、API文档

- [WebAPI](http://127.0.0.1:8081/swagger/index.html)
- [Game](./docs/game/README.md)