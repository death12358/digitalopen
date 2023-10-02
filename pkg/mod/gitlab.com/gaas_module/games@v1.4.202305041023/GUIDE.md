# GUIDE

## Introduction

This guide is designed to help you get started with the development of probability programs.

First, you need to learn how to use the following languages or tools:

* [golang](https://golang.org/)
* [docker](#docker)
* [docker-compose](#docker-compose)
* [redis](#redis)

### Docker

Installation instructions can be found in the Docker official documentation: [Docker official](https://docs.docker.com/get-docker/).

Docker 基本命令，包括：

* docker run：執行一個 Docker 容器。
* docker start：啟動一個停止運行的 Docker 容器。
* docker stop：停止一個正在運行的 Docker 容器。
* docker pull：從 Docker Hub 下載一個 Docker 鏡像。
* docker build：從 Dockerfile 建立一個 Docker 鏡像。

更多 Docker 命令的用法可以參考 Docker 官方文檔：
[Docker 官方文檔](https://docs.docker.com/engine/reference/commandline/docker/)。

---

### Docker-Compose

Docker Compose 是一個用於定義和執行多容器 Docker 應用程序的工具。它能夠讓您使用單個配置文件將所有容器的配置信息放在一起，並通過一個單一的命令就能啟動和管理整個應用程序。

以下是一個使用 Docker Compose 的基本操作範例：

1. 首先，在您的系統上安裝 Docker Compose。安裝請見官方文檔：[Docker Compose](https://docs.docker.com/compose/install/)。
2. 接下來，在您的項目目錄中創建一個名為 docker-compose.yml 的文件，並在其中定義您的應用程序的容器配置。
3. 最後，在創建 docker-compose.yml 目錄中，執行 docker-compose up 命令，即可啟動您的應用程序。

```bash
docker-compose up
```

將啟動您於 docker-compose.yml 中設定的應用程序，並在前景執行。如果您希望在後台執行，可以使用 -d 參數。

```bash
docker-compose up -d
```

4. 如果您希望停止應用程序，可以使用 docker-compose down 命令。

```bash
docker-compose down
```

將停止您於 docker-compose.yml 中設定的應用程序。

---

### redis

Redis是一種開源的資料結構伺服器，它能夠用於儲存資料以便快速訪問和操作。

我們這邊使用 docker-compose 來快速啟動 redis，詳細的 docker-compose.yml 可以參考以下：

```yaml
version: '3'

services:
  redis:
    image: redis:6-alpine
    restart: always
    container_name: redis
    ports: 
      - 6379:6379
    command: ["redis-server",
    "--bind","0.0.0.0",
    "--appendonly","yes",
    "--databases","10",
    "--maxmemory-policy","volatile-lru",
    "--maxmemory","4g",
    "--hz","100"]
```

基本操作可參考官方文檔：
[Redis 官方文檔](https://redis.io/documentation)。
或參考以下：
[菜鸟教程](https://www.runoob.com/redis/redis-tutorial.html)
