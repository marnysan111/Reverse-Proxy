# Reverse-Proxy

Goでできている**リバースプロキシ**です。

<center>
<img src="https://i.imgur.com/UUVEo1x.png" width="75%">
</center>

|用途|番号|
|---|---|
|ユーザのアクセス|80|
|CPU使用率の管理|50000|
|コンテンツ|任意|


# 実行環境
|名称|バージョン|
|---|---|
|go|1.15.2|
|MySQL|ver 14.14 Distrib 5.7.33|
|Docker|19.03.13|



## Installation

```
$ git clone https://github.com/marnysan111/Reverse-Proxy.git
```
## 初設定
- Appとして登録するサーバの情報を記す。
```
$ vi Reverse-Proxy/Host/cpu/ip.json


[
    # {"IPADD": "IPアドレス", "PORT": "リバースプロキシ用のポート番号", "FORWARD": "コンテンツ用のポート番号"}
    （例）
    {"IPADD": "192.168.1.10", "PORT": "50000", "FORWARD": "3000"},
    {"IPADD": "192.168.1.11", "PORT": "50000", "FORWARD": "3000"},
    ・
    ・
]
```
- MySQLにて、データベースの作成・情報登録
```
$ vi Reverse-Proxy/Host/.env

DBUSER=ユーザ名
DBPASS=パスワード
DBIP=tcp(127.0.0.1:3306)
DBTABLE=データベース名

$ mysql -u root -p


mysql> create database proxy;
```


## 実行

### Host

```
$ cd Reverse-Proxy/Host/proxy
$ docker-compose build
$ docker-compose up -d
$ cd ../cpu
$ go build
$ ./cpu
```

### App

```
$ cd Reverse-Proxy/App
$ go build
$ ./app
```





