package mobile

import (
	"github.com/cyjaysong/shumaiapi-go"
)

type MobileEmptyReq struct {
	AppId     string `json:"appid"`     // 服务商分配的唯一标识
	Timestamp string `json:"timestamp"` // 当前时间的毫秒数
	Sign      string `json:"sign"`      // 签名，签名算法说明
	Mobile    string `json:"mobile"`    // 电话号码
}

func (req *MobileEmptyReq) Do(cli *shumaiapi.Client) (res *shumaiapi.BaseRes[MobileEmptyResData], err error) {
	req.AppId, req.Timestamp, req.Sign = cli.Sign()
	res = &shumaiapi.BaseRes[MobileEmptyResData]{}
	if err = cli.Get("/v4/mobile_empty/check", req, res); err != nil {
		return nil, err
	}
	return
}

type MobileEmptyResData struct {
	OrderNo string `json:"order_no"` // 订单号
	Area    string `json:"area"`     // 手机号归属地
	Channel string `json:"channel"`  // 运营商
	Status  int    `json:"status"`   // 手机号状态
}
