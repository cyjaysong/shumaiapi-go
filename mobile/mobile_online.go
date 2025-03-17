package mobile

import (
	"github.com/bytedance/sonic"
	"github.com/cyjaysong/shumaiapi-go"
)

type MobileOnlineReq struct {
	AppId     string `json:"appid"`     // 服务商分配的唯一标识
	Timestamp string `json:"timestamp"` // 当前时间的毫秒数
	Sign      string `json:"sign"`      // 签名，签名算法说明
	Mobile    string `json:"mobile"`    // 电话号码
}

func (req *MobileOnlineReq) Do(cli *shumaiapi.Client) (res *shumaiapi.BaseRes[MobileOnlineResData], err error) {
	req.AppId, req.Timestamp, req.Sign = cli.Sign()
	res = &shumaiapi.BaseRes[MobileOnlineResData]{}
	if err = cli.Get("/v2/mobile_online/check", req, res); err != nil {
		return nil, err
	}
	return
}

type MobileOnlineResData struct {
	OrderNo string `json:"order_no"` // 订单号
	Channel string `json:"channel"`  // 运营商 cmcc:移动；cucc:联通；ctcc:电信
	Time    string `json:"time"`     // 在网时间区间值：[0,3), [3,6), [6,12), [12,24), [24,-1)(单位为月)
}

func (data MobileOnlineResData) Times() (times []int) {
	_ = sonic.UnmarshalString(data.Time, &times)
	return
}
