# sbiport-server
- SBI証券のポートフォリオの情報を管理するAPI

## Usage
```
Usage:
  sbiport-server start [flags]

Flags:
      --db-host string   DB Host
      --db-name string   DB Name
      --db-pass string   DB Pass
      --db-port string   DB Port
      --db-user string   DB User
  -h, --help             help for start
```

# sbiport-client
- CSVファイルを `sbiport-server` に登録するためのクライアント
- ファイル名の形式は、`2006-01-02.csv` か `20060102.csv`

## Usage
```
Usage:
  sbiport-client regist <file or directory name> [flags]

Flags:
  -h, --help            help for regist
      --host string     server host
      --port string     server port
      --scheme string   http or https
```

# sbi-fetcher
- Web Page(https://www.sbisec.co.jp/ETGate) から元データを取り込んでくるスクリプト
- 環境変数 `sbi_user` に ID、`sbi_pass` に パスワード、`CSV_DIR` に保存先ディレクトリを入れる
