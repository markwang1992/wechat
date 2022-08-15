package suite

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	// GetPermanentCodeURL 获取企业永久授权码
	GetPermanentCodeURL = "https://qyapi.weixin.qq.com/cgi-bin/service/get_permanent_code?suite_access_token=%s"
	// SetSessionInfoURL 设置授权配置
	SetSessionInfoURL = "https://qyapi.weixin.qq.com/cgi-bin/service/set_session_info?suite_access_token=%s"
)

type (
	// GetPermanentCodeRequest 获取企业永久授权码请求
	GetPermanentCodeRequest struct {
		AuthCode string `json:"auth_code"`
	}
	// GetPermanentCodeResponse 获取企业永久授权码响应
	GetPermanentCodeResponse struct {
		util.CommonError

		AccessToken    string `json:"access_token"`
		ExpiresIn      int    `json:"expires_in"`
		PermanentCode  string `json:"permanent_code"`
		DealerCorpInfo struct {
			Corpid   string `json:"corpid"`
			CorpName string `json:"corp_name"`
		} `json:"dealer_corp_info"`
		AuthCorpInfo struct {
			Corpid            string `json:"corpid"`
			CorpName          string `json:"corp_name"`
			CorpType          string `json:"corp_type"`
			CorpSquareLogoUrl string `json:"corp_square_logo_url"`
			CorpUserMax       int    `json:"corp_user_max"`
			CorpFullName      string `json:"corp_full_name"`
			VerifiedEndTime   int    `json:"verified_end_time"`
			SubjectType       int    `json:"subject_type"`
			CorpWxqrcode      string `json:"corp_wxqrcode"`
			CorpScale         string `json:"corp_scale"`
			CorpIndustry      string `json:"corp_industry"`
			CorpSubIndustry   string `json:"corp_sub_industry"`
		} `json:"auth_corp_info"`
		AuthInfo struct {
			Agent []struct {
				Agentid          int    `json:"agentid"`
				Name             string `json:"name"`
				RoundLogoUrl     string `json:"round_logo_url"`
				SquareLogoUrl    string `json:"square_logo_url"`
				Appid            int    `json:"appid"`
				AuthMode         int    `json:"auth_mode"`
				IsCustomizedApp  bool   `json:"is_customized_app"`
				AuthFromThirdapp bool   `json:"auth_from_thirdapp"`
				Privilege        struct {
					Level      int      `json:"level"`
					AllowParty []int    `json:"allow_party"`
					AllowUser  []string `json:"allow_user"`
					AllowTag   []int    `json:"allow_tag"`
					ExtraParty []int    `json:"extra_party"`
					ExtraUser  []string `json:"extra_user"`
					ExtraTag   []int    `json:"extra_tag"`
				} `json:"privilege"`
				SharedFrom struct {
					Corpid    string `json:"corpid"`
					ShareType int    `json:"share_type"`
				} `json:"shared_from"`
			} `json:"agent"`
		} `json:"auth_info"`
		AuthUserInfo struct {
			Userid     string `json:"userid"`
			OpenUserid string `json:"open_userid"`
			Name       string `json:"name"`
			Avatar     string `json:"avatar"`
		} `json:"auth_user_info"`
		RegisterCodeInfo struct {
			RegisterCode string `json:"register_code"`
			TemplateId   string `json:"template_id"`
			State        string `json:"state"`
		} `json:"register_code_info"`
		State string `json:"state"`
	}
)

// GetPermanentCode 获取永久授权码
// see https://developer.work.weixin.qq.com/document/path/90603
func (r *Client) GetPermanentCode(request *GetPermanentCodeRequest) (*GetPermanentCodeResponse, error) {
	var (
		response []byte
		err      error
	)
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	response, err = util.HTTPPost(fmt.Sprintf(GetPermanentCodeURL, r.SuiteAccessToken), string(jsonData))
	if err != nil {
		return nil, err
	}
	result := &GetPermanentCodeResponse{}
	err = util.DecodeWithError(response, result, "GetPermanentCode")
	if err != nil {
		return nil, err
	}
	return result, nil
}

type (
	// SetSessionInfoRequest 设置授权配置请求
	SetSessionInfoRequest struct {
		PreAuthCode string `json:"pre_auth_code"`
		SessionInfo struct {
			AuthType int `json:"auth_type"`
		} `json:"session_info"`
	}
	// SetSessionInfoResponse 设置授权配置相应
	SetSessionInfoResponse struct {
		util.CommonError
	}
)

// SetSessionInfo 设置授权配置
// see https://developer.work.weixin.qq.com/document/10975#%E8%AE%BE%E7%BD%AE%E6%8E%88%E6%9D%83%E9%85%8D%E7%BD%AE
func (r *Client) SetSessionInfo(request *SetSessionInfoRequest) error {
	var (
		response []byte
		err      error
	)
	jsonData, err := json.Marshal(request)
	if err != nil {
		return err
	}
	response, err = util.HTTPPost(fmt.Sprintf(SetSessionInfoURL, r.SuiteAccessToken), string(jsonData))
	if err != nil {
		return err
	}
	result := &SetSessionInfoResponse{}
	err = util.DecodeWithError(response, result, "SetSessionInfo")
	if err != nil {
		return err
	}
	if result.ErrCode == 0 {
		return nil
	}
	return fmt.Errorf(result.ErrMsg)
}
