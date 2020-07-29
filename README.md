# go-join

[中文文档](https://github.com/yuchenfw/go-join/blob/master/README_CN.md)

`go-join` is a tool to join values with the defined rules easily.It commonly used to join data with specific keys' order.

general join form：

key1`KVSep`value1`Sep`key2`KVSep`value2`Sep`...`Sep`keyN`KVSep`valueN

## get & import

### get

if you use go mod ,just add below to `go.mod` file.

```txt
require github.com/yuchenfw/go-join latest
```

or

```txt
go get -u github.com/yuchenfw/go-join
```

### import

```go
import "github.com/yuchenfw/go-join"
```

## Quick Start

`go-join` is very easy to use , it just has a func.

```go
gojoin.Join(src interface{}, options Options) (dst string, err error) {
```

Join supports `map`(contains `url.Values`),`struct`,encoded `url string` and their pointer format.

options contains the rules to join .

```go
type Options struct {
    Sep           string    // the separator between keys' values (maybe contains keys)
    KVSep         string    // the separator between key & its value
    IgnoreKey     bool      // whether ignore key , if yes, the key will be ignored, but the value will reserve
    IgnoreEmpty   bool      // whether ignore empty value , if yes, the key & its value will be ignored
    ExceptKeys    []string  // the keys & their values will be ignored
    Order         joinOrder // the join order
    DefinedOrders []string  // the keys order, using with Order == Defined
    StructTag     string    // struct tag, using when src struct type, if not set, will use struct filed name, only support export fields
    URLCoding     urlCoding // the value format, using when format value
    Unwrap        bool      // whether unwrap the internal map or struct
}
```

| composite type    | handle rules                               |
| ----------- | -------------------------------------- |
| slice/array | if len>=1，get the first value；otherwise, if the element type is the basic type,get the zero value,else will be ignore |
| map/struct  | ignore                              |

If you want unwrap `map` or `struct`,just set`Unwrap`=`true`.If the internal fileds have the same key name with others, the later will cover the pre-one.

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
