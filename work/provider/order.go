// @Author markwang <wangyu@uniondrug.cn>
// @Date   2022/8/11

package provider

import (
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	// CreateNewOrderURL 下单购买帐号
	CreateNewOrderURL = "https://qyapi.weixin.qq.com/cgi-bin/license/create_new_order?provider_access_token=%s"
	// ListOrderAccountURL 获取订单中的帐号列表
	ListOrderAccountURL = "https://qyapi.weixin.qq.com/cgi-bin/license/list_order_account?provider_access_token=%s"
	// ActiveAccountURL 激活帐号
	ActiveAccountURL = "https://qyapi.weixin.qq.com/cgi-bin/license/active_account?provider_access_token=%s"
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
		OrderID string `json:"order_id"`
	}
)

// CreateNewOrder 下单购买帐号
// see https://developer.work.weixin.qq.com/document/path/95644
func (r *Client) CreateNewOrder(request *CreateNewOrderRequest) (string, error) {
	var (
		err      error
		response []byte
	)
	if response, err = util.PostJSON(fmt.Sprintf(CreateNewOrderURL, r.AccessToken), request); err != nil {
		return "", err
	}
	result := &CreateNewOrderResponse{}
	if err = util.DecodeWithError(response, result, "CreateNewOrder"); err != nil {
		return "", err
	}
	return result.OrderID, nil
}

type (
	ListOrderAccountRequest struct {
		OrderID string `json:"order_id"`
		Limit   int    `json:"limit"`
		Cursor  string `json:"cursor"`
	}
	ListOrderAccountResponse struct {
		util.CommonError
		NextCursor  string         `json:"next_cursor"`
		HasMore     int            `json:"has_more"`
		AccountList []*AccountList `json:"account_list"`
	}
	AccountList struct {
		ActiveCode string `json:"active_code"`
		UserID     string `json:"userid"`
		Type       int    `json:"type"`
	}
)

// ListOrderAccount 获取订单中的帐号列表
// see https://developer.work.weixin.qq.com/document/path/95649
func (r *Client) ListOrderAccount(request *ListOrderAccountRequest) (*ListOrderAccountResponse, error) {
	var (
		err      error
		response []byte
	)
	if response, err = util.PostJSON(fmt.Sprintf(ListOrderAccountURL, r.AccessToken), request); err != nil {
		return nil, err
	}
	result := &ListOrderAccountResponse{}
	if err = util.DecodeWithError(response, result, "ListOrderAccount"); err != nil {
		return nil, err
	}
	return result, nil
}

type (
	ActiveAccountRequest struct {
		ActiveCode string `json:"active_code"`
		CorpID     string `json:"corpid"`
		UserID     string `json:"userid"`
	}
	ActiveAccountResponse struct {
		util.CommonError
	}
)

// ActiveAccount 激活帐号
// see https://developer.work.weixin.qq.com/document/path/95553
func (r *Client) ActiveAccount(request *ActiveAccountRequest) (*ActiveAccountResponse, error) {
	var (
		err      error
		response []byte
	)
	if response, err = util.PostJSON(fmt.Sprintf(ActiveAccountURL, r.AccessToken), request); err != nil {
		return nil, err
	}
	result := &ActiveAccountResponse{}
	if err = util.DecodeWithError(response, result, "ActiveAccount"); err != nil {
		return nil, err
	}
	return result, nil
}
