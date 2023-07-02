package wechat

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type AccessTokenResp struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// AccessToken
// https://developer.work.weixin.qq.com/document/path/91039#15074
func AccessToken(ctx context.Context) (string, error) {
	resp, err := http.Get("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=ww0980d2723fb39cdb&corpsecret=BYgtnJGYH3FMODUCNbKC2gg3paTHNYBt-8gKNDQmaoc")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var accessTokenResp = AccessTokenResp{}
	err = json.Unmarshal(body, &accessTokenResp)
	if err != nil {
		return "", err
	}

	return accessTokenResp.AccessToken, nil
}
