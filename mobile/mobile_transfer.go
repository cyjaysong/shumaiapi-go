package mobile

import "github.com/cyjaysong/shumaiapi-go"

type MobileTransferQueryReq struct {
	AppId     string `json:"appid"`     // 服务商分配的唯一标识
	Timestamp string `json:"timestamp"` // 当前时间的毫秒数
	Sign      string `json:"sign"`      // 签名，签名算法说明
	Mobile    string `json:"mobile"`    // 电话号码
}

func (req *MobileTransferQueryReq) Do(cli *shumaiapi.Client) (res *shumaiapi.BaseRes[MobileTransferQueryResData], err error) {
	req.AppId, req.Timestamp, req.Sign = cli.Sign()
	res = &shumaiapi.BaseRes[MobileTransferQueryResData]{}
	if err = cli.Get("/v4/mobile-transfer/query", req, res); err != nil {
		return nil, err
	}
	return
}

type MobileTransferQueryResData struct {
	OrderNo    string `json:"order_no"`   // 订单号
	Mobile     string `json:"mobile"`     // 提交的手机号码
	Area       string `json:"area"`       // 手机号归属地(预留字段)
	IspType    string `json:"ispType"`    // 转网前运营商
	NewIspType string `json:"newIspType"` // 转网后运营商
}
