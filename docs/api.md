## GET /
- healthCheck 用

## GET /daily/{YYYYMMDD}
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



## POST /regist/{YYYYMMDD}
- データを追加する。
### request:
- request body 部分
- 例
    ```
    取引,ファンド名,買付日,数量,取得単価,現在値,前日比,前日比（％）,損益,損益（％）,評価額,編集
    積立  売却,AAA,--/--/--,26231,13000,11403,-258,-2.21,-4189.09,-12.28,29911.2,詳細 

    積立  売却,BBB,--/--/--,10946,31610,29726,+235,+0.80,-2062.22,-5.96,32538.07,詳細 

    ```
    - Webから元データを取り込んでくるスクリプトは https://github.com/azuki774/sbi-fetcher の形式

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
