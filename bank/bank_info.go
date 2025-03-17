package bank

import "github.com/cyjaysong/shumaiapi-go"

type BankInfoReq struct {
	AppId     string `json:"appid"`              // 服务商分配的唯一标识
	Timestamp string `json:"timestamp"`          // 当前时间的毫秒数
	Sign      string `json:"sign"`               // 签名，签名算法说明
	Bankcard  string `json:"bankcard,omitempty"` // 卡号
	Bank      string `json:"bank,omitempty"`     // 银行名称
	Province  string `json:"province,omitempty"` // 银行所在省
	City      string `json:"city,omitempty"`     // 银行所在市
	Key       string `json:"key,omitempty"`      // 关键字，匹配支行名称
	Page      int    `json:"page,omitempty"`     // 页数 1-5页，默认第一页
}

func (req *BankInfoReq) Do(cli *shumaiapi.Client) (res *shumaiapi.BaseRes[BankInfoResData], err error) {
	req.AppId, req.Timestamp, req.Sign = cli.Sign()
	res = &shumaiapi.BaseRes[BankInfoResData]{}
	if err = cli.Get("/v4/bank_info/query", req, res); err != nil {
		return nil, err
	}
	return
}

type BankInfoResData struct {
	OrderNo      string `json:"order_no"`
	Bank         string `json:"bank"`
	Province     string `json:"province"`
	City         string `json:"city"`
	CardName     string `json:"card_name"`
	Tel          string `json:"tel"`
	Type         string `json:"type"`
	Logo         string `json:"logo"`
	Abbreviation string `json:"abbreviation"`
	CardBin      string `json:"card_bin"`
	BinDigits    int    `json:"bin_digits"`
	CardDigits   int    `json:"card_digits"`
	IsLuhn       bool   `json:"isLuhn"`
	Weburl       string `json:"weburl"`
}
