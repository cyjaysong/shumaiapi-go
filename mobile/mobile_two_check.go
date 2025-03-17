package mobile

import "github.com/cyjaysong/shumaiapi-go"

type MobileTwoCheckReq struct {
	AppId     string `json:"appid"`     // 服务商分配的唯一标识
	Timestamp string `json:"timestamp"` // 当前时间的毫秒数
	Sign      string `json:"sign"`      // 签名，签名算法说明
	Name      string `json:"name"`      // 姓名
	Mobile    string `json:"mobile"`    // 电话号码
}

func (req *MobileTwoCheckReq) Do(cli *shumaiapi.Client) (res *shumaiapi.BaseRes[MobileTwoCheckResData], err error) {
	req.AppId, req.Timestamp, req.Sign = cli.Sign()
	res = &shumaiapi.BaseRes[MobileTwoCheckResData]{}
	if err = cli.Get("/v2/mobile_two/check", req, res); err != nil {
		return nil, err
	}
	return
}

type MobileTwoCheckResData struct {
	OrderNo string `json:"order_no"` // 订单号
	Fee     int    `json:"fee"`      // 是否计费 0 不计费；1 计费
	Result  int    `json:"result"`   // 验证结果 0 一致；1 不一致
}

type MobileTwoCheckV2Req struct {
	AppId     string `json:"appid"`     // 服务商分配的唯一标识
	Timestamp string `json:"timestamp"` // 当前时间的毫秒数
	Sign      string `json:"sign"`      // 签名，签名算法说明
	IdCard    string `json:"idcard"`    // 身份证号
	Mobile    string `json:"mobile"`    // 电话号码
}

func (req *MobileTwoCheckV2Req) Do(cli *shumaiapi.Client) (res *shumaiapi.BaseRes[MobileTwoCheckV2ResData], err error) {
	req.AppId, req.Timestamp, req.Sign = cli.Sign()
	res = &shumaiapi.BaseRes[MobileTwoCheckV2ResData]{}
	if err = cli.Get("/v4/mobile2/check", req, res); err != nil {
		return nil, err
	}
	return
}

type MobileTwoCheckV2ResData struct {
	OrderNo string `json:"order_no"` // 订单号
	Desc    string `json:"desc"`     // 验证结果描述
	Result  int    `json:"result"`   // 验证结果 0 一致 计费；1 不一致 计费 ；2 无记录
}
