package mobile

import "github.com/cyjaysong/shumaiapi-go"

type MobileStatusReq struct {
	AppId     string `json:"appid"`     // 服务商分配的唯一标识
	Timestamp string `json:"timestamp"` // 当前时间的毫秒数
	Sign      string `json:"sign"`      // 签名，签名算法说明
	Mobile    string `json:"mobile"`    // 电话号码
}

func (req *MobileStatusReq) Do(cli *shumaiapi.Client) (res *shumaiapi.BaseRes[MobileStatusResData], err error) {
	req.AppId, req.Timestamp, req.Sign = cli.Sign()
	res = &shumaiapi.BaseRes[MobileStatusResData]{}
	if err = cli.Get("/v1/mobile_status/check", req, res); err != nil {
		return nil, err
	}
	return
}

type MobileStatusResData struct {
	OrderNo string `json:"order_no"` // 订单号
	Channel string `json:"channel"`  // 运营商 移动；联通；电信；广电
	Status  int    `json:"status"`   // 在网状态，0 在网；1 不在网或携号转网/销户
	Desc    string `json:"desc"`     // 不在网原因
}
