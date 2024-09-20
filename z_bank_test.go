package shumaiapi_test

import (
	"fmt"
	"github.com/cyjaysong/shumaiapi-go/bank"
	"testing"
)

func TestBankInfo(t *testing.T) {
	req := &bank.BankInfoReq{Bankcard: "4563514204078808813"}
	req.AppId, req.Timestamp, req.Sign = client.Sign()
	res := &bank.BankInfoRes{}
	err := client.Get(req.Path(), req, res)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", res)
}

func TestLianHangQuery(t *testing.T) {
	req := &bank.BankBranchQueryReq{Bankcard: "4563514204078808813"}
	req.AppId, req.Timestamp, req.Sign = client.Sign()
	res := &bank.BankBranchQueryRes{}
	err := client.Get(req.Path(), req, res)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", res)
}
