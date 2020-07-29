# go-join

go-join是一个可以快速按照指定规则拼接字符串的工具，主要用于按照特定key顺序的数据拼接。

拼接的格式：

>key1`KVSep`value1`Sep`key2`KVSep`value2`Sep`...`Sep`keyN`KVSep`valueN

## get & import

### get

如果使用go mod ,将以下内容添加至`go.mod`文件中即可。

```txt
require github.com/yuchenfw/go-join latest
```

如果使用`gopath`，使用以下方式添加。

```txt
go get -u github.com/yuchenfw/go-join
```

### import

```go
import "github.com/yuchenfw/go-join"
```

## Quick Start

`go-join`仅有一个func，就是Join。

```go
gojoin.Join(src interface{}, options Options) (dst string, err error) {
```

Join支持`map`(包括 `url.Values`)、`struct`、encoded的`url string`及它们的指针形式。

options是拼接的具体参数设置，结构如下：

```go
type Options struct {
    Sep           string    // keys对应values（可能包含keys）间的连接符
    KVSep         string    // key和它对应value间的连接符
    IgnoreKey     bool      // 是否忽略key，如果true，则连带其对应的value一起忽略
    IgnoreEmpty   bool      // 是否忽略空值（主要针对空字符串）的key，如果是，当其值为空时忽略
    ExceptKeys    []string  // 被排除拼接的keys，直接忽略
    Order         joinOrder // 拼接规则
    DefinedOrders []string  // 当Order == Defined时，指定的keys顺序
    StructTag     string    // struct tag，用于指定field别名，如不指定则使用field name，仅支持struct且export的fields
    URLCoding     urlCoding // key对应value的处理方式
    Unwrap        bool      // 当内部存在复合类型如map，struct，是否解析，是则解析，默认false。
}
```

注意：

如果key对应的value是复合类型时，处理规则如下：

| 复合类型    | 处理规则                               |
| ----------- | -------------------------------------- |
| slice/array | 如果len>=1，取第一个值；否则，内部元素类型是基本类型的取零值，其他类型直接忽略 |
| map/struct  | 直接忽略                               |

如果需要解析`map`或者`struct`，需要设置`Unwrap`=`true`。

另外，多个复合类型中如果出现的同名key，后出现的会覆盖之前出现的key。

### examples

1.Join map

```go
gojoin.Join(map[string]interface{}{
               "total_amount": "2.00",
               "buyer_id":     "2088102116773037",
               "body":         "大乐透2.1",
               "trade_no":     "2016071921001003030200089909",
               "refund_fee":   "0.00",
               "notify_time":  "2016-07-19 14:10:49",
               "subject":      "大乐透2.1",
               "sign_type":    "RSA2",
               "charset":      "utf-8",
               "notify_type":  "trade_status_sync",
               "out_trade_no": "0719141034-6418",
               "gmt_close":    "2016-07-19 14:10:46",
               "gmt_payment":  "2016-07-19 14:10:47",
               "trade_status": "TRADE_SUCCESS",
               "version":      "1.0",
               "sign":         "kPbQIjX+xQc8F0/A6/AocELIjhhZnGbcBN6G4MM/HmfWL4ZiHM6fWl5NQhzXJusaklZ1LFuMo+lHQUELAYeugH8LYFvxnNajOvZhuxNFbN2LhF0l/KL8ANtj8oyPM4NN7Qft2kWJTDJUpQOzCzNnV9hDxh5AaT9FPqRS6ZKxnzM=",
               "gmt_create":   "2016-07-19 14:10:44",
               "app_id":       "2015102700040153",
               "seller_id":    "2088102119685838",
               "notify_id":    "4a91b7a78a503640467525113fb7d8bg8e",
            },Options{
               Sep:        "&",
               KVSep:      "=",
               Order:      ASCII,
               ExceptKeys: []string{"sign", "sign_type"},
            })
```

the join result is `app_id=2015102700040153&body=大乐透2.1&buyer_id=2088102116773037&charset=utf-8&gmt_close=2016-07-19 14:10:46&gmt_create=2016-07-19 14:10:44&gmt_payment=2016-07-19 14:10:47&notify_id=4a91b7a78a503640467525113fb7d8bg8e&notify_time=2016-07-19 14:10:49&notify_type=trade_status_sync&out_trade_no=0719141034-6418&refund_fee=0.00&seller_id=2088102119685838&subject=大乐透2.1&total_amount=2.00&trade_no=2016071921001003030200089909&trade_status=TRADE_SUCCESS&version=1.0`

2.Join struct

```go
gojoin.Join(callbackWrap{
               &callback{
                  "2.00",
                  "2088102116773037",
                  "大乐透2.1",
                  "2016071921001003030200089909",
                  "0.00",
                  "2016-07-19 14:10:49",
                  "大乐透2.1",
                  "RSA2",
                  "utf-8",
                  "trade_status_sync",
                  "0719141034-6418",
                  "2016-07-19 14:10:46",
                  "2016-07-19 14:10:47",
                  "TRADE_SUCCESS",
                  "1.0",
                  "kPbQIjX+xQc8F0/A6/AocELIjhhZnGbcBN6G4MM/HmfWL4ZiHM6fWl5NQhzXJusaklZ1LFuMo+lHQUELAYeugH8LYFvxnNajOvZhuxNFbN2LhF0l/KL8ANtj8oyPM4NN7Qft2kWJTDJUpQOzCzNnV9hDxh5AaT9FPqRS6ZKxnzM=",
                  "2016-07-19 14:10:44",
                  "2015102700040153",
                  "2088102119685838",
                  "4a91b7a78a503640467525113fb7d8bg8e",
               },
            },Options{
               Sep:        "&",
               KVSep:      "=",
               Order:      ASCII,
               ExceptKeys: []string{"sign", "sign_type"},
               Unwrap:     true,
               StructTag:  "json",
            })

type callback struct {
   TotalAmount string `json:"total_amount"`
   BuyerId     string `json:"buyer_id"`
   Body        string `json:"body"`
   TradeNo     string `json:"trade_no"`
   RefundFee   string `json:"refund_fee"`
   NotifyTime  string `json:"notify_time"`
   Subject     string `json:"subject"`
   SignType    string `json:"sign_type"`
   Charset     string `json:"charset"`
   NotifyType  string `json:"notify_type"`
   OutTradeNo  string `json:"out_trade_no"`
   GmtClose    string `json:"gmt_close"`
   GmtPayment  string `json:"gmt_payment"`
   TradeStatus string `json:"trade_status"`
   Version     string `json:"version"`
   Sign        string `json:"sign"`
   GmtCreate   string `json:"gmt_create"`
   AppId       string `json:"app_id"`
   SellerId    string `json:"seller_id"`
   NotifyId    string `json:"notify_id"`
}

type callbackWrap struct {
   Callback *callback `json:"callback"`
}
```

the join result is `app_id=2015102700040153&body=大乐透2.1&buyer_id=2088102116773037&charset=utf-8&gmt_close=2016-07-19 14:10:46&gmt_create=2016-07-19 14:10:44&gmt_payment=2016-07-19 14:10:47&notify_id=4a91b7a78a503640467525113fb7d8bg8e&notify_time=2016-07-19 14:10:49&notify_type=trade_status_sync&out_trade_no=0719141034-6418&refund_fee=0.00&seller_id=2088102119685838&subject=大乐透2.1&total_amount=2.00&trade_no=2016071921001003030200089909&trade_status=TRADE_SUCCESS&version=1.0`

3.Join url encoded string

```go
gojoin.Join("https://api.xx.com/receive_notify.htm?total_amount=2.00&buyer_id=2088102116773037&body=大乐透2.1&trade_no=2016071921001003030200089909&refund_fee=0.00&notify_time=2016-07-19 14:10:49&subject=大乐透2.1&sign_type=RSA2&charset=utf-8&notify_type=trade_status_sync&out_trade_no=0719141034-6418&gmt_close=2016-07-19 14:10:46&gmt_payment=2016-07-19 14:10:47&trade_status=TRADE_SUCCESS&version=1.0&sign=kPbQIjX%2bxQc8F0%2fA6%2fAocELIjhhZnGbcBN6G4MM%2fHmfWL4ZiHM6fWl5NQhzXJusaklZ1LFuMo%2blHQUELAYeugH8LYFvxnNajOvZhuxNFbN2LhF0l%2fKL8ANtj8oyPM4NN7Qft2kWJTDJUpQOzCzNnV9hDxh5AaT9FPqRS6ZKxnzM%3d&gmt_create=2016-07-19 14:10:44&app_id=2015102700040153&seller_id=2088102119685838&notify_id=4a91b7a78a503640467525113fb7d8bg8e",Options{
               Sep:        "&",
               KVSep:      "=",
               Order:      ASCII,
               StructTag:  "json",
               ExceptKeys: []string{"sign", "sign_type"},
            })
```

the join result is `app_id=2015102700040153&body=大乐透2.1&buyer_id=2088102116773037&charset=utf-8&gmt_close=2016-07-19 14:10:46&gmt_create=2016-07-19 14:10:44&gmt_payment=2016-07-19 14:10:47&notify_id=4a91b7a78a503640467525113fb7d8bg8e&notify_time=2016-07-19 14:10:49&notify_type=trade_status_sync&out_trade_no=0719141034-6418&refund_fee=0.00&seller_id=2088102119685838&subject=大乐透2.1&total_amount=2.00&trade_no=2016071921001003030200089909&trade_status=TRADE_SUCCESS&version=1.0`
