# jmaxml-seis-parser-go

気象庁防災情報 XML の パーサー Go 実装

## 対応フォーマット

- VXSE43: 緊急地震速報（警報）
- VXSE51: 震度速報
- VXSE52: 地震情報（震源に関する情報）
- VXSE53: 地震情報（震源・震度に関する情報）
- VTSE41: 津波警報・注意報・予報

## 使用方法

[Releases](https://github.com/p2pquake/jmaxml-seis-parser-go/releases) から実行可能なバイナリが入手可能です。

```sh
$ ./jmaxml-seis-parser-go
気象庁防災情報 XML 地震火山情報の一部 (VXSE43, VXSE51, VXSE52, VXSE53, VTSE41) のパーサ

Usage:
  jmaxml-seis-parser-go [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  convert     XML から EPSP JSON 形式への変換
  help        Help about any command
  parse       XML のパース

Flags:
  -h, --help      help for jmaxml-seis-parser-go
  -v, --version   version for jmaxml-seis-parser-go

Use "jmaxml-seis-parser-go [command] --help" for more information about a command.

$ ./jmaxml-seis-parser-go convert data/20210216071046_0_VXSE53_270000.xml
{"expire":null,"issue":{"source":"気象庁","time":"2021/02/16 16:10:46","type":"DetailScale","correct":"None"},"earthquake":{"time":"2021/02/16 16:07:00","hypocenter":{"name":"大阪府南部","latitude":34.3,"longitude":135.2,"depth":0,"magnitude":1.5},"maxScale":10,"domesticTsunami":"None","foreignTsunami":"Unknown"},"points":[{"pref":"和歌山県","addr":"和歌山市一番丁","scale":10,"isArea":false}]}
```

## 参考

- [気象庁 | 気象庁防災情報XMLフォーマット](http://xml.kishou.go.jp/)
- [p2pquake/dmdata-jp-api-v2-websocket-client: DMDATA.JP (Project DM-D.S.S) API v2 の非公式 WebSocket クライアント](https://github.com/p2pquake/dmdata-jp-api-v2-websocket-client)

## ライセンス

- ソースコード: MIT License
- XML ファイル: CC BY 4.0 ([気象庁防災情報 XML](http://xml.kishou.go.jp/xmlpull.html) Atom フィード)
