package gojoin

import (
	"net/url"
	"testing"
)

func TestJoinValues(t *testing.T) {
	type args struct {
		src     interface{}
		options Options
	}
	tests := []struct {
		name       string
		args       args
		wantValues string
		wantErr    bool
	}{
		{
			args: args{
				src: map[string]interface{}{
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
				},
				options: Options{
					Sep:        "&",
					KVSep:      "=",
					Order:      ASCII,
					ExceptKeys: []string{"sign", "sign_type"},
				},
			},
			wantErr:    false,
			wantValues: "app_id=2015102700040153&body=大乐透2.1&buyer_id=2088102116773037&charset=utf-8&gmt_close=2016-07-19 14:10:46&gmt_create=2016-07-19 14:10:44&gmt_payment=2016-07-19 14:10:47&notify_id=4a91b7a78a503640467525113fb7d8bg8e&notify_time=2016-07-19 14:10:49&notify_type=trade_status_sync&out_trade_no=0719141034-6418&refund_fee=0.00&seller_id=2088102119685838&subject=大乐透2.1&total_amount=2.00&trade_no=2016071921001003030200089909&trade_status=TRADE_SUCCESS&version=1.0",
		},
		{
			args: args{
				src: url.Values{
					"total_amount": {"2.00"},
					"buyer_id":     {"2088102116773037"},
					"body":         {"大乐透2.1"},
					"trade_no":     {"2016071921001003030200089909"},
					"refund_fee":   {"0.00"},
					"notify_time":  {"2016-07-19 14:10:49"},
					"subject":      {"大乐透2.1"},
					"sign_type":    {"RSA2"},
					"charset":      {"utf-8"},
					"notify_type":  {"trade_status_sync"},
					"out_trade_no": {"0719141034-6418"},
					"gmt_close":    {"2016-07-19 14:10:46"},
					"gmt_payment":  {"2016-07-19 14:10:47"},
					"trade_status": {"TRADE_SUCCESS"},
					"version":      {"1.0"},
					"sign":         {"kPbQIjX+xQc8F0/A6/AocELIjhhZnGbcBN6G4MM/HmfWL4ZiHM6fWl5NQhzXJusaklZ1LFuMo+lHQUELAYeugH8LYFvxnNajOvZhuxNFbN2LhF0l/KL8ANtj8oyPM4NN7Qft2kWJTDJUpQOzCzNnV9hDxh5AaT9FPqRS6ZKxnzM="},
					"gmt_create":   {"2016-07-19 14:10:44"},
					"app_id":       {"2015102700040153"},
					"seller_id":    {"2088102119685838"},
					"notify_id":    {"4a91b7a78a503640467525113fb7d8bg8e"},
				},
				options: Options{
					Sep:        "&",
					KVSep:      "=",
					Order:      ASCII,
					ExceptKeys: []string{"sign", "sign_type"},
				},
			},
			wantErr:    false,
			wantValues: "app_id=2015102700040153&body=大乐透2.1&buyer_id=2088102116773037&charset=utf-8&gmt_close=2016-07-19 14:10:46&gmt_create=2016-07-19 14:10:44&gmt_payment=2016-07-19 14:10:47&notify_id=4a91b7a78a503640467525113fb7d8bg8e&notify_time=2016-07-19 14:10:49&notify_type=trade_status_sync&out_trade_no=0719141034-6418&refund_fee=0.00&seller_id=2088102119685838&subject=大乐透2.1&total_amount=2.00&trade_no=2016071921001003030200089909&trade_status=TRADE_SUCCESS&version=1.0",
		},
		{
			args: args{
				src: &callback{
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
				options: Options{
					Sep:        "&",
					KVSep:      "=",
					Order:      ASCII,
					StructTag:  "json",
					ExceptKeys: []string{"sign", "sign_type"},
				},
			},
			wantErr:    false,
			wantValues: "app_id=2015102700040153&body=大乐透2.1&buyer_id=2088102116773037&charset=utf-8&gmt_close=2016-07-19 14:10:46&gmt_create=2016-07-19 14:10:44&gmt_payment=2016-07-19 14:10:47&notify_id=4a91b7a78a503640467525113fb7d8bg8e&notify_time=2016-07-19 14:10:49&notify_type=trade_status_sync&out_trade_no=0719141034-6418&refund_fee=0.00&seller_id=2088102119685838&subject=大乐透2.1&total_amount=2.00&trade_no=2016071921001003030200089909&trade_status=TRADE_SUCCESS&version=1.0",
		},
		{
			args: args{
				src: "https://api.xx.com/receive_notify.htm?total_amount=2.00&buyer_id=2088102116773037&body=大乐透2.1&trade_no=2016071921001003030200089909&refund_fee=0.00&notify_time=2016-07-19 14:10:49&subject=大乐透2.1&sign_type=RSA2&charset=utf-8&notify_type=trade_status_sync&out_trade_no=0719141034-6418&gmt_close=2016-07-19 14:10:46&gmt_payment=2016-07-19 14:10:47&trade_status=TRADE_SUCCESS&version=1.0&sign=kPbQIjX%2bxQc8F0%2fA6%2fAocELIjhhZnGbcBN6G4MM%2fHmfWL4ZiHM6fWl5NQhzXJusaklZ1LFuMo%2blHQUELAYeugH8LYFvxnNajOvZhuxNFbN2LhF0l%2fKL8ANtj8oyPM4NN7Qft2kWJTDJUpQOzCzNnV9hDxh5AaT9FPqRS6ZKxnzM%3d&gmt_create=2016-07-19 14:10:44&app_id=2015102700040153&seller_id=2088102119685838&notify_id=4a91b7a78a503640467525113fb7d8bg8e",
				options: Options{
					Sep:        "&",
					KVSep:      "=",
					Order:      ASCII,
					StructTag:  "json",
					ExceptKeys: []string{"sign", "sign_type"},
				},
			},
			wantErr:    false,
			wantValues: "app_id=2015102700040153&body=大乐透2.1&buyer_id=2088102116773037&charset=utf-8&gmt_close=2016-07-19 14:10:46&gmt_create=2016-07-19 14:10:44&gmt_payment=2016-07-19 14:10:47&notify_id=4a91b7a78a503640467525113fb7d8bg8e&notify_time=2016-07-19 14:10:49&notify_type=trade_status_sync&out_trade_no=0719141034-6418&refund_fee=0.00&seller_id=2088102119685838&subject=大乐透2.1&total_amount=2.00&trade_no=2016071921001003030200089909&trade_status=TRADE_SUCCESS&version=1.0",
		},
		{
			args: args{
				src: "https://api.xx.com/receive_notify.htm?total_amount=2.00&buyer_id=2088102116773037&body=大乐透2.1&trade_no=2016071921001003030200089909&refund_fee=0.00&notify_time=2016-07-19 14:10:49&subject=大乐透2.1&sign_type=RSA2&charset=utf-8&notify_type=trade_status_sync&out_trade_no=0719141034-6418&gmt_close=2016-07-19 14:10:46&gmt_payment=2016-07-19 14:10:47&trade_status=TRADE_SUCCESS&version=1.0&sign=kPbQIjX%2bxQc8F0%2fA6%2fAocELIjhhZnGbcBN6G4MM%2fHmfWL4ZiHM6fWl5NQhzXJusaklZ1LFuMo%2blHQUELAYeugH8LYFvxnNajOvZhuxNFbN2LhF0l%2fKL8ANtj8oyPM4NN7Qft2kWJTDJUpQOzCzNnV9hDxh5AaT9FPqRS6ZKxnzM%3d&gmt_create=2016-07-19 14:10:44&app_id=2015102700040153&seller_id=2088102119685838&notify_id=4a91b7a78a503640467525113fb7d8bg8e",
				options: Options{
					Sep:           "&",
					KVSep:         "=",
					Order:         Defined,
					DefinedOrders: []string{"app_id", "body", "buyer_id", "charset", "gmt_close", "gmt_create", "gmt_payment", "notify_id", "notify_time", "notify_type", "out_trade_no", "refund_fee", "seller_id", "subject", "total_amount", "trade_no", "trade_status", "version"},
					StructTag:     "json",
				},
			},
			wantErr:    false,
			wantValues: "app_id=2015102700040153&body=大乐透2.1&buyer_id=2088102116773037&charset=utf-8&gmt_close=2016-07-19 14:10:46&gmt_create=2016-07-19 14:10:44&gmt_payment=2016-07-19 14:10:47&notify_id=4a91b7a78a503640467525113fb7d8bg8e&notify_time=2016-07-19 14:10:49&notify_type=trade_status_sync&out_trade_no=0719141034-6418&refund_fee=0.00&seller_id=2088102119685838&subject=大乐透2.1&total_amount=2.00&trade_no=2016071921001003030200089909&trade_status=TRADE_SUCCESS&version=1.0",
		},
		{
			args: args{
				src: map[string]interface{}{
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
					"more": map[string]interface{}{
						"seller_id": "2088102119685838",
						"notify_id": "4a91b7a78a503640467525113fb7d8bg8e",
					},
				},
				options: Options{
					Sep:        "&",
					KVSep:      "=",
					Order:      ASCII,
					ExceptKeys: []string{"sign", "sign_type"},
					Unwrap:     true,
				},
			},
			wantErr:    false,
			wantValues: "app_id=2015102700040153&body=大乐透2.1&buyer_id=2088102116773037&charset=utf-8&gmt_close=2016-07-19 14:10:46&gmt_create=2016-07-19 14:10:44&gmt_payment=2016-07-19 14:10:47&notify_id=4a91b7a78a503640467525113fb7d8bg8e&notify_time=2016-07-19 14:10:49&notify_type=trade_status_sync&out_trade_no=0719141034-6418&refund_fee=0.00&seller_id=2088102119685838&subject=大乐透2.1&total_amount=2.00&trade_no=2016071921001003030200089909&trade_status=TRADE_SUCCESS&version=1.0",
		},
		{
			args: args{
				src: callbackWrap{
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
				},
				options: Options{
					Sep:        "&",
					KVSep:      "=",
					Order:      ASCII,
					ExceptKeys: []string{"sign", "sign_type"},
					Unwrap:     true,
					StructTag:  "json",
				},
			},
			wantErr:    false,
			wantValues: "app_id=2015102700040153&body=大乐透2.1&buyer_id=2088102116773037&charset=utf-8&gmt_close=2016-07-19 14:10:46&gmt_create=2016-07-19 14:10:44&gmt_payment=2016-07-19 14:10:47&notify_id=4a91b7a78a503640467525113fb7d8bg8e&notify_time=2016-07-19 14:10:49&notify_type=trade_status_sync&out_trade_no=0719141034-6418&refund_fee=0.00&seller_id=2088102119685838&subject=大乐透2.1&total_amount=2.00&trade_no=2016071921001003030200089909&trade_status=TRADE_SUCCESS&version=1.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValues, err := Join(tt.args.src, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("JoinValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValues != tt.wantValues {
				t.Errorf("JoinValues() gotValues = %v, want %v", gotValues, tt.wantValues)
			}
		})
	}
}

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
