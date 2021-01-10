# todo

01/10 追記  
- ラズパイからPOSTされたデータをどう比較するのか→どう値を保存して、比較するのか
- CPU使用率とメモリ使用率を-からの絶対値で評価
- jsonで管理


cpuの取得のため、jsonで管理する  
http://pppurple.hatenablog.com/entry/2018/04/26/225932


できたこと

- Serverに対してClientがPOSTできる


CPUに負荷を掛けるコマンド
```
$ yes > /dev/null
```
確認するコマンド
```
$ vmstat -t 2  --unit M
```



ロードバランサーとして、なにを見て判断するか
- CPU使用率
- メモリ使用率


