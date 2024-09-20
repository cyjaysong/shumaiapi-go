package shumaiapi

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	reqclient "github.com/imroc/req/v3"
	"strconv"
	"time"
)

type Client struct {
	appId       string
	appSecurity string
	HttpClient  *reqclient.Client
}

func New(appId, appSecurity string) *Client {
	client := &Client{appId: appId, appSecurity: appSecurity}
	client.HttpClient = reqclient.C().SetTimeout(time.Second * 10).SetCommonRetryCount(2)
	client.HttpClient.SetBaseURL("https://api.shumaidata.com").SetUserAgent("").DevMode()
	return client
}

func (c *Client) Sign() (appId, timestamp, sign string) {
	timestamp = strconv.FormatInt(time.Now().UnixMilli(), 10)
	ret := md5.Sum([]byte(fmt.Sprintf("%s&%s&%s", c.appId, timestamp, c.appSecurity)))
	return c.appId, timestamp, hex.EncodeToString(ret[:])
}

func (c *Client) Get(path string, req, res any) (err error) {
	var jsonNode ast.Node
	if jsonBytes, err := sonic.Marshal(req); err != nil {
		return err
	} else if jsonNode, err = sonic.Get(jsonBytes); err != nil {
		return err
	}
	queryParams, _ := jsonNode.Map()
	httpRes, err := c.HttpClient.R().SetQueryParamsAnyType(queryParams).Get(path)
	if err != nil {
		return err
	} else if httpRes.StatusCode != 200 {
		return fmt.Errorf("http error, status code:%d, msg:%s", httpRes.StatusCode, httpRes.Status)
	} else if err = httpRes.Unmarshal(res); err != nil {
		return err
	}
	return
}

func (c *Client) PostJson(path string, req, res any) (err error) {
	httpRes, err := c.HttpClient.R().SetBodyJsonMarshal(req).Post(path)
	if err != nil {
		return err
	} else if httpRes.StatusCode != 200 {
		return fmt.Errorf("http error, status code:%d, msg:%s", httpRes.StatusCode, httpRes.Status)
	} else if err = httpRes.Unmarshal(res); err != nil {
		return err
	}
	return
}

func (c *Client) PostForm(path string, req, res any) (err error) {
	var jsonNode ast.Node
	if jsonBytes, err := sonic.Marshal(req); err != nil {
		return err
	} else if jsonNode, err = sonic.Get(jsonBytes); err != nil {
		return err
	}
	formData, _ := jsonNode.Map()
	httpRes, err := c.HttpClient.R().SetFormDataAnyType(formData).Post(path)
	if err != nil {
		return err
	} else if httpRes.StatusCode != 200 {
		return fmt.Errorf("http error, status code:%d, msg:%s", httpRes.StatusCode, httpRes.Status)
	} else if err = httpRes.Unmarshal(res); err != nil {
		return err
	}
	return
}
