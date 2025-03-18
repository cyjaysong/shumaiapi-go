package mobile

import "github.com/cyjaysong/shumaiapi-go"

type MobileThreeCheckReq struct {
	AppId     string `json:"appid"`     // 服务商分配的唯一标识
	Timestamp string `json:"timestamp"` // 当前时间的毫秒数
	Sign      string `json:"sign"`      // 签名，签名算法说明
	Name      string `json:"name"`      // 姓名
	IdCard    string `json:"idcard"`    // 身份证号
	Mobile    string `json:"mobile"`    // 电话号码
}

func (req *MobileThreeCheckReq) Do(cli *shumaiapi.Client) (res *shumaiapi.BaseRes[MobileThreeCheckResData], err error) {
	req.AppId, req.Timestamp, req.Sign = cli.Sign()
	res = &shumaiapi.BaseRes[MobileThreeCheckResData]{}
	if err = cli.Get("/v4/mobile_three/check", req, res); err != nil {
		return nil, err
	}
	return
}

type MobileThreeCheckResData struct {
	OrderNo  string `json:"order_no"`      // 订单号
	Result   int    `json:"result,string"` // 验证结果 0: 一致，收费；1: 不一致，收费；2: 无记录，不收费
	Desc     string `json:"desc"`          // 结果描述
	Channel  string `json:"channel"`       // 运营商 cmcc:移动；cucc:联通；ctcc:电信；gdcc:广电
	Sex      string `json:"sex"`           // 性别
	Birthday string `json:"birthday"`      // 生日
	Address  string `json:"address"`       // 籍贯
}

type MobileThreeDetailCheckReq struct {
	AppId     string `json:"appid"`     // 服务商分配的唯一标识
	Timestamp string `json:"timestamp"` // 当前时间的毫秒数
	Sign      string `json:"sign"`      // 签名，签名算法说明
	Name      string `json:"name"`      // 姓名
	IdCard    string `json:"idcard"`    // 身份证号
	Mobile    string `json:"mobile"`    // 电话号码
}

func (req *MobileThreeDetailCheckReq) Do(cli *shumaiapi.Client) (res *shumaiapi.BaseRes[MobileThreeDetailResData], err error) {
	req.AppId, req.Timestamp, req.Sign = cli.Sign()
	res = &shumaiapi.BaseRes[MobileThreeDetailResData]{}
	if err = cli.Get("/v4/mobile_three/check/detail", req, res); err != nil {
		return nil, err
	}
	return
}

type MobileThreeDetailResData struct {
	OrderNo string `json:"order_no"` // 订单号
	Result  int    `json:"result"`   // 验证结果 1:一致;2:手机号已实名，但是身份证和姓名均与实名信息不一致;3:手机号已实名，手机号和证件号一致，姓名不一致;4:手机号已实名，手机号和姓名一致，身份证不一致;5:无记录
}
