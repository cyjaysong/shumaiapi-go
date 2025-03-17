package shumaiapi

type BaseRes[T any] struct {
	Success bool   `json:"success"` // 接口请求成功标识，true为成功，false为失败，失败情况下，会有对应描述和状态码
	Msg     string `json:"msg"`     // code对应的说明描述
	Code    int    `json:"code"`    // 成功为200，其它为失败状态码
	Data    T      `json:"data"`    // 结果详细信息
}
