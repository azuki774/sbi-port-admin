## sbiport-client

```
./sbiport-client regist <dir>
```
- `<dir>` に入っているCSVを sbiport-server 経由でDBに登録
- `<dir>` 内のファイル命名規則は `YYYYMMDD_<categoryTag>.csv`
    - YYYYMMDD は登録日時
    - categoryTag は 登録するテーブルを識別するポートフォリオ種別
    - `/regist/<category_tag>/{YYYYMMDD}` に対応
