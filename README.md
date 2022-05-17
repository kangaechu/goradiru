# goradiru

## Overview

らじる・らじる 取得

## Usage

```
goradiru.go ― らじる らじるを取得

 Usage: goradiru <command> [arguments...] [options...]

 コマンドの簡単な説明:
   dl     指定したエピソードをダウンロードします
   pg     最新のプログラムを取得します


```

### 一覧を取得


```
$ goradiru pg
  - Name: DJ日本史
    Url: https://www.nhk.or.jp/hoge1
  - Name: ちきゅうラジオ
    Url: https://www.nhk.or.jp/hoge2
  - Name: 石丸謙二郎の山カフェ
    Url: https://www.nhk.or.jp/hoge3
  - Name: きこえタマゴ！
    Url: https://www.nhk.or.jp/hoge4
  - Name: NHKジャーナル
    Url: https://www.nhk.or.jp/hoge4
>
```

### 取得対象を設定ファイルに追加

`config/conf.yaml` に取得したいUrlを記述

```yaml
ProgDir: ./prog
DownloadedHistoryConfFile: config/downloaded.yaml
Programs:
  - Name: FMシアター
    Url: https://www.nhk.or.jp/hoge1
  - Name: 新日曜名作座
    Url: https://www.nhk.or.jp/hoge2
```

### 番組を取得

```shell-session
$ goradiru dl
```
