# Reverse-Proxy

Goでできている**リバースプロキシ**です。

![](https://i.imgur.com/ETUaFqA.png)

|用途|番号|
|---|---|
|ユーザのアクセス|80|
|CPU使用率の管理|50000|

## Installation

```
$ git clone https://github.com/marnysan111/Reverse-Proxy.git
```

### Server

```
$ cd Reverse-Proxy/server/proxy
$ docker-compose build
$ docker-compose up -d
$ cd ../cpu
$ go build
$ ./cpu
```

### Client

```
$ cd Reverse-Proxy/Client/
```
