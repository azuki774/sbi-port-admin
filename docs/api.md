## GET /
- healthCheck 用

## GET /daily/<category_tag>/{YYYYMMDD}
- daily 単位で登録済のデータを取得する

### request:
- パスパラメータで日付を `YYYYMMDD` 指定する

### response:
-
```
[
    {
        RecordDate    : "20060102", 
        FundName      : "AAA",
        Amount        : 47917,
        AcquisitionPrice  : 12796,
        NowPrice          : 12796,
        ThedayBefore      : -284,
        ThedayBeforeRatio : -2.16,
        Profit            : 326.03,
        ProfitRatio       : 0.53,
        Valuation         : 61679.02
    },
    {
        RecordDate    : "20060102", 
        FundName      : "BBB",
        Amount        : 47917,
        AcquisitionPrice  : 12796,
        NowPrice          : 12796,
        ThedayBefore      : -284,
        ThedayBeforeRatio : -2.16,
        Profit            : 326.03,
        ProfitRatio       : 0.53,
        Valuation         : 61679.02
    }
]
```



## POST /regist/<category_tag>/{YYYYMMDD}
- データを追加する。
### request:
- request body 部分
- 例
    ```
    取引,ファンド名,買付日 ,数量,取得単価,現在値,前日比,前日比（％）,損益,損益（％）,評価額,編集
    積立  売却,AAA,--/--/--,"131,210","12,888","16,075",-43,-0.27,"+41,816.62",+24.73,"210,920.07",詳細  
    積立  売却,BBB,--/--/--,"49,572","34,213","46,607",-178,-0.38,"+61,439.53",+36.23,"231,040.22",詳細  
    ```
    - Webから元データを取り込んでくるスクリプトは https://github.com/azuki774/myscrapers/ の形式

### response:
-
```
{
    "created_number": 1, // 新しくDBに登録した数
    "updated_number": 1, // DB上のデータを更新した数
    "skipped_number": 1, // DB上にデータがあり、処理をスキップした数
    "failed_number": 1, // 処理に失敗した数
}
```
