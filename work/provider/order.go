// @Author markwang <wangyu@uniondrug.cn>
// @Date   2022/8/11

package provider

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	// CreateNewOrderURL 下单购买帐号
	CreateNewOrderURL = "https://qyapi.weixin.qq.com/cgi-bin/license/create_new_order?provider_access_token=%s"
)

type (
	// CreateNewOrderRequest 下单购买帐号请求
	CreateNewOrderRequest struct {
		CorpID          string          `json:"corpid"`
		BuyerUserID     string          `json:"buyer_userid"`
		AccountCount    AccountCount    `json:"account_count"`
		AccountDuration AccountDuration `json:"account_duration"`
	}
	AccountCount struct {
		BaseCount            int `json:"base_count"`
		ExternalContactCount int `json:"external_contact_count"`
	}
	AccountDuration struct {
		Months int `json:"months"`
	}
	// CreateNewOrderResponse 下单购买帐号响应
	CreateNewOrderResponse struct {
		util.CommonError
		OrderId string `json:"order_id"`
	}
)

// CreateNewOrder 下单购买帐号
// see https://developer.work.weixin.qq.com/document/path/95644
func (r *Client) CreateNewOrder(request *CreateNewOrderRequest) (string, error) {
	var (
		response []byte
		err      error
	)
	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", err
	}
	if response, err = util.HTTPPost(fmt.Sprintf(CreateNewOrderURL, r.AccessToken), string(jsonData)); err != nil {
		return "", err
	}
	result := &CreateNewOrderResponse{}
	if err = util.DecodeWithError(response, result, "CreateNewOrder"); err != nil {
		return "", err
	}
	return result.OrderId, nil
}
