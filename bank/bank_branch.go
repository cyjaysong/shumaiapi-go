package bank

type BankBranchQueryReq struct {
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

func (BankBranchQueryReq) Path() string {
	return "/v2/lianhang/query"
}

type BankBranchQueryRes struct {
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Data    struct {
		OrderNo string `json:"order_no"` //订单号
		Result  struct {
			Totalpage  int              `json:"totalpage,string"`  //总页数
			Totalcount int              `json:"totalcount,string"` //总记录数
			Bank       string           `json:"bank"`              //输入的银行名称
			Province   string           `json:"province"`          //输入的省
			City       string           `json:"city"`              //输入的市
			Record     []LianHangRecord `json:"record"`
			Page       int              `json:"page"`
			Card       string           `json:"card"`
			Key        string           `json:"key"`
		} `json:"result"`
	} `json:"data"`
}

type LianHangRecord struct {
	Bank     string `json:"bank"`          //总行名称
	Lname    string `json:"lname"`         //支行名称
	Bankcode string `json:"bankcode"`      //联行号
	IsHead   int    `json:"isHead,string"` //总行标识 0-否 1-是 预留
	Tel      string `json:"tel"`           //支行电话
	Addr     string `json:"addr"`          //支行地址
	Province string `json:"province"`      //省
	City     string `json:"city"`          //市
	District string `json:"district"`      //所在区 预留
	Lng      string `json:"lng"`           //经度 预留
	Lat      string `json:"lat"`           //纬度 预留
}
