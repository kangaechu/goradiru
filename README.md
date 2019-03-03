# goradiru

## Overview

らじる・らじる 取得

## Usage

### 設定の記述

`config/conf.yaml` に取得したいURLを記述

```yaml
ProgDir: ./prog
DownloadedHistoryConfFile: config/downloaded.yaml
Programs:
  - Name: FMシアター
    Url: https://www.nhk.or.jp/radioondemand/json/0058/bangumi_0058_01.json
  - Name: 新日曜名作座
    Url: https://www.nhk.or.jp/radioondemand/json/0930/bangumi_0930_01.json
#  - Name: 青春アドベンチャー
#    Url: https://www.nhk.or.jp/radioondemand/json/0164/bangumi_0164_01.json
#  - Name: 特集オーディオドラマ
#    Url: https://www.nhk.or.jp/radioondemand/json/P000025/bangumi_P000025_01.json
```

### 実行

```shell-session
$ goradiru
```