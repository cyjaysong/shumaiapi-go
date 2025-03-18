package mobile

import "github.com/cyjaysong/shumaiapi-go"

type MobileTwiceReq struct {
	AppId     string `json:"appid"`     // 服务商分配的唯一标识
	Timestamp string `json:"timestamp"` // 当前时间的毫秒数
	Sign      string `json:"sign"`      // 签名，签名算法说明
	Mobile    string `json:"mobile"`    // 电话号码
	Date      string `json:"date"`      // 日期,yyyyMMdd格式
}

func (req *MobileTwiceReq) Do(cli *shumaiapi.Client) (res *shumaiapi.BaseRes[MobileTwiceResData], err error) {
	req.AppId, req.Timestamp, req.Sign = cli.Sign()
	res = &shumaiapi.BaseRes[MobileTwiceResData]{}
	if err = cli.Get("/v4/mobile_twice/check", req, res); err != nil {
		return nil, err
	}
	return
}

type MobileTwiceResData struct {
	OrderNo string `json:"orderNo"` // 订单号
	Result  int    `json:"result"`  // 验证结果 0 是二次卡， 1 不是二次卡， 2 数据库中无信息(预留)
	Channel string `json:"channel"` // 运营商，cmcc:移动 cucc:联通 ctcc:电信
	Desc    string `json:"desc"`    // 验证结果描述
}
