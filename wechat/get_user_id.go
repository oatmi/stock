package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetUserIDResp struct {
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	Userid     string `json:"userid"`
	UserTicket string `json:"user_ticket"`
}

// https://developer.work.weixin.qq.com/document/path/91023
func GetUserID(ctx context.Context, AT, code string) (string, error) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/auth/getuserinfo?access_token=%s&code=%s",
		AT,
		code,
	)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response = GetUserIDResp{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	return response.Userid, nil
}
