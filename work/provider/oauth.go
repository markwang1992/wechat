// @Author markwang <wangyu@uniondrug.cn>
// @Date   2022/8/16

package provider

import (
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	// GetLoginInfoURL 获取登录用户信息
	GetLoginInfoURL = "https://qyapi.weixin.qq.com/cgi-bin/service/get_login_info?access_token=%s"
)

type (
	// GetLoginInfoRequest 获取登录用户信息请求
	GetLoginInfoRequest struct {
		AuthCode string `json:"auth_code"`
	}
	// GetLoginInfoResponse 下单购买帐号响应
	GetLoginInfoResponse struct {
		util.CommonError
		Usertype int      `json:"usertype"`
		UserInfo UserInfo `json:"user_info"`
		CorpInfo CorpInfo `json:"corp_info"`
		Agent    []*Agent `json:"agent"`
		AuthInfo AuthInfo `json:"auth_info"`
	}
	UserInfo struct {
		UserID     string `json:"userid"`
		OpenUserID string `json:"open_userid"`
		Name       string `json:"name"`
		Avatar     string `json:"avatar"`
	}
	CorpInfo struct {
		CorpID string `json:"corpid"`
	}
	Agent struct {
		AgentID  int `json:"agentid"`
		AuthType int `json:"auth_type"`
	}
	AuthInfo struct {
		Department []*Department `json:"department"`
	}
	Department struct {
		Id       int  `json:"id"`
		Writable bool `json:"writable"`
	}
)

// GetLoginInfo 获取登录用户信息
// see https://developer.work.weixin.qq.com/document/path/91125
func (r *Client) GetLoginInfo(request *GetLoginInfoRequest) (*GetLoginInfoResponse, error) {
	var (
		err      error
		response []byte
	)
	if response, err = util.PostJSON(fmt.Sprintf(GetLoginInfoURL, r.AccessToken), request); err != nil {
		return nil, err
	}
	result := &GetLoginInfoResponse{}
	if err = util.DecodeWithError(response, result, "GetLoginInfo"); err != nil {
		return nil, err
	}
	return result, nil
}
