package shumaiapi_test

import (
	"fmt"
	"github.com/cyjaysong/shumaiapi-go/mobile"
	"testing"
)

func TestMobileThreeCheck(t *testing.T) {
	req := &mobile.MobileThreeCheckReq{
		Name:   "数脉",
		IdCard: "500220199703176688",
		Mobile: "15066668888",
	}
	res, err := req.Do(client)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", res)
}

func TestMobileThreeDetailCheck(t *testing.T) {
	req := &mobile.MobileThreeDetailCheckReq{
		Name:   "数脉",
		IdCard: "500220199703176688",
		Mobile: "15066668888",
	}
	res, err := req.Do(client)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", res)
}

func TestMobileTwoCheck(t *testing.T) {
	req := &mobile.MobileTwoCheckReq{
		Name:   "数脉",
		Mobile: "15066668888",
	}
	res, err := req.Do(client)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", res)
}

func TestMobileTwoV2Check(t *testing.T) {
	req := &mobile.MobileTwoCheckV2Req{
		IdCard: "500220199703176688",
		Mobile: "15066668888",
	}
	res, err := req.Do(client)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", res)
}

func TestMobileStatusCheck(t *testing.T) {
	req := &mobile.MobileStatusReq{
		Mobile: "15066668888",
	}
	res, err := req.Do(client)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", res)
}
func TestMobileOnlineCheck(t *testing.T) {
	req := &mobile.MobileOnlineReq{
		Mobile: "15066668888",
	}
	res, err := req.Do(client)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", res)
}
func TestMobileEmptyCheck(t *testing.T) {
	req := &mobile.MobileEmptyReq{
		Mobile: "15066668888",
	}
	res, err := req.Do(client)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", res)
}
func TestMobileTransferQuery(t *testing.T) {
	req := &mobile.MobileTransferQueryReq{
		Mobile: "15066668888",
	}
	res, err := req.Do(client)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", res)
}
func TestMobileTwiceQuery(t *testing.T) {
	req := &mobile.MobileTwiceReq{
		Mobile: "15066668888",
		Date:   "20250317",
	}
	res, err := req.Do(client)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", res)
}
